package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router         *mux.Router
	MessageService MessageService
	Server         *http.Server
}

func NewHandler(messageService MessageService) *Handler {
	h := &Handler{
		MessageService: messageService,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LogRequest)
	h.Router.Use(TimeoutMiddleware)
	h.Server = &http.Server{
		Addr:    "0.0.0.0:4000",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	}).Methods("GET")
	h.messageRoutes()
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	chann := make(chan os.Signal, 1)
	signal.Notify(chann, os.Interrupt)
	<-chann

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("Server shutdown")
	return nil
}
