package main

import (
	"fmt"
	transportHttp "golang/golang-study/cmd/server/transport/http"
	"golang/golang-study/internal/conversation"
	"golang/golang-study/internal/database"
)

func Run() error {
	fmt.Println("Starting API")
	db, err := database.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Errorf("ERROR RUNNING UP %w", err)
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
		fmt.Println("ERROR running the API")
		fmt.Println(err)
	}
}
