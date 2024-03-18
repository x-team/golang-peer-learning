package tests

import (
	"context"
	"golang/golang-study/internal/conversation"
	"golang/golang-study/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMessage(t *testing.T) {
	t.Run("When message does not exists", func(t *testing.T) {
		mockStore := mocks.NewMessageStore(t)

		service := conversation.NewMessageService(mockStore)
		mockStore.On("GetMessage", mock.Anything, mock.AnythingOfType("string")).
			Return(func(ctx context.Context, s string) (conversation.Message, error) {
				return conversation.Message{}, nil
			})

		message, err := service.GetMessage(context.Background(), "1")

		assert.NoError(t, err)
		assert.Equal(t, conversation.Message{}, message)
		mockStore.AssertCalled(t, "GetMessage", context.Background(), "1")
	})

	t.Run("When message exists", func(t *testing.T) {
		mockStore := mocks.NewMessageStore(t)

		service := conversation.NewMessageService(mockStore)
		expected := conversation.Message{
			ID:            "1",
			Content:       "THE_CONTENT",
			ToPhoneNumber: "123",
		}
		mockStore.On("GetMessage", mock.Anything, mock.AnythingOfType("string")).
			Return(func(ctx context.Context, s string) (conversation.Message, error) {
				return expected, nil
			})

		message, err := service.GetMessage(context.Background(), "1")

		assert.NoError(t, err)
		assert.Equal(t, expected, message)
		mockStore.AssertCalled(t, "GetMessage", context.Background(), "1")
	})
}
