package main

import (
	"fmt"

	"golang/golang-study/internal/conversation"
	"golang/golang-study/internal/database"
	transportHttp "golang/golang-study/internal/server/transport/http"
)

func Run() error {
	fmt.Println("Starting API")
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect")
		return err
	}

	if err = db.MigrateDB(); err != nil {
		fmt.Printf("ERROR MIGRATING UP THE DATABASE %v\n", err)
		return err
	}

	messageService := conversation.NewMessageService(db)
	httpHandler := transportHttp.NewHandler(messageService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Printf("ERROR running the API %v\n", err)
	}
}
