package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"golang/golang-study/internal/conversation"
	"golang/golang-study/internal/database"
)

func TestMessageDatabase(t *testing.T) {
	t.Run("Test create message", func(t *testing.T) {
		db, err := database.NewDatabase()
		assert.NoError(t, err)

		msg, err := db.PostMessage(context.Background(), conversation.Message{
			Content:       "TEST_COMMENT",
			ToPhoneNumber: "+5541996740459",
		})

		assert.NoError(t, err)

		newMsg, err := db.GetMessage(context.Background(), msg.ID)
		assert.NoError(t, err)

		assert.Equal(t, msg.Content, newMsg.Content)
	})

	t.Run("Test delete message", func(t *testing.T) {

		db, err := database.NewDatabase()
		assert.NoError(t, err)

		msg, err := db.PostMessage(context.Background(), conversation.Message{
			Content:       "TEST_COMMENT",
			ToPhoneNumber: "+5541996740459",
		})
		assert.NoError(t, err)

		deleted, err := db.DeleteMessage(context.Background(), msg.ID)
		assert.NoError(t, err)

		afterDelete, err := db.GetMessage(context.Background(), msg.ID)
		assert.Error(t, err, "Error fetching message sql: no rows in result set")

		assert.True(t, deleted)
		assert.Equal(t, conversation.Message{}, afterDelete)
	})
}
