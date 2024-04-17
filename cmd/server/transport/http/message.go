package http

import (
	"context"
	"encoding/json"
	"fmt"
	"golang/golang-study/internal/conversation"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type MessageService interface {
	GetMessage(ctx context.Context, id string) (conversation.Message, error)
	UpdateMessage(ctx context.Context, id string, updated conversation.Message) (conversation.Message, error)
	DeleteMessage(ctx context.Context, id string) error
	CreateMessage(ctx context.Context, message conversation.Message) (conversation.Message, error)
}

func getId(r *http.Request) string {
	args := mux.Vars(r)
	fmt.Println(args)
	return args["id"]
}

func decodeMessage(r *http.Request) (conversation.Message, error) {
	var msg conversation.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		return conversation.Message{}, err
	}

	return msg, nil
}

type PostCommentRequest struct {
	Content       string `json:"content" validate:"required"`
	ToPhoneNumber string `json:"toPhoneNumber" validate:"required"`
}

func convertPostMessageRequestToMessage(request PostCommentRequest) conversation.Message {
	return conversation.Message{
		Content:       request.Content,
		ToPhoneNumber: request.ToPhoneNumber,
	}
}

func (h *Handler) PostMessage(w http.ResponseWriter, r *http.Request) {
	var msgRequest PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&msgRequest); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(msgRequest)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	postedMessage, err := h.MessageService.CreateMessage(r.Context(), convertPostMessageRequestToMessage(msgRequest))

	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(postedMessage); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	msgToUpdate, err := decodeMessage(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := getId(r)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg, err := h.MessageService.UpdateMessage(r.Context(), id, msgToUpdate)

	if err := json.NewEncoder(w).Encode(msg); err != nil {
		log.Print(err)
		return
	}
}

func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.MessageService.DeleteMessage(r.Context(), id); err != nil {
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg, err := h.MessageService.GetMessage(r.Context(), id)

	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(msg); err != nil {
		log.Print(err)
		return
	}
}

func (h *Handler) messageRoutes() {
	h.Router.HandleFunc("/api/v1/message", JWtAuth(h.PostMessage)).Methods("POST")
	h.Router.HandleFunc("/api/v1/message/{id}", h.GetMessage).Methods("GET")
	h.Router.HandleFunc("/api/v1/message/{id}", JWtAuth(h.DeleteMessage)).Methods("DELETE")
	h.Router.HandleFunc("/api/v1/message/{id}", JWtAuth(h.UpdateMessage)).Methods("PUT")
}
