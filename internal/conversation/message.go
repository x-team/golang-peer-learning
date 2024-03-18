package conversation

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingMessage = errors.New("failed to fetch message")
)

type Message struct {
	ID            string
	Content       string
	ToPhoneNumber string
}

type MessageStore interface {
	GetMessage(context.Context, string) (Message, error)
	PostMessage(context.Context, Message) (Message, error)
	DeleteMessage(ctx context.Context, uuid string) (bool, error)
	UpdateMessage(ctx context.Context, id string, message Message) (Message, error)
}

type MessageService struct {
	Repository MessageStore
	Repo2      MessageStore
	UniqueID   string
}

func NewMessageService(repository MessageStore) *MessageService {
	return &MessageService{
		Repository: repository,
		Repo2:      repository,
	}
}

func MutateMessage(messasge *Message) {
	messasge.ToPhoneNumber = "3184931849"
}

func (s *MessageService) GetMessage(ctx context.Context, id string) (Message, error) {
	message, err := s.Repository.GetMessage(ctx, id)
	if err != nil {
		fmt.Println(err)

		return Message{}, ErrFetchingMessage
	}

	MutateMessage(&message)

	return message, nil
}

func (s *MessageService) UpdateMessage(ctx context.Context, id string, updated Message) (Message, error) {
	return s.Repository.UpdateMessage(ctx, id, updated)
}

func (s *MessageService) DeleteMessage(ctx context.Context, id string) error {
	_, err := s.Repository.DeleteMessage(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *MessageService) CreateMessage(ctx context.Context, message Message) (Message, error) {
	message, err := s.Repository.PostMessage(ctx, message)

	if err != nil {
		return Message{}, err
	}

	return message, nil
}
