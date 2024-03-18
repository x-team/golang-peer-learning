package database

import (
	"context"
	"database/sql"
	"fmt"
	"golang/golang-study/internal/conversation"

	uuid "github.com/satori/go.uuid"
)

type MessageEntity struct {
	ID            string
	ToPhoneNumber sql.NullString
	Content       sql.NullString
}

func toConversationMessage(message MessageEntity) conversation.Message {
	return conversation.Message{
		ID:            message.ID,
		Content:       message.Content.String,
		ToPhoneNumber: message.ToPhoneNumber.String,
	}
}

func (d *Database) GetMessage(ctx context.Context, uuid string) (conversation.Message, error) {
	var message MessageEntity
	row := d.Client.QueryRowContext(ctx, `SELECT * FROM messages WHERE id = $1`, uuid)

	err := row.Scan(&message.ID, &message.Content, &message.ToPhoneNumber)
	if err != nil {
		return conversation.Message{}, fmt.Errorf("Error fetching message %w", err)
	}

	return toConversationMessage(message), nil
}

func (d *Database) PostMessage(ctx context.Context, message conversation.Message) (conversation.Message, error) {
	message.ID = uuid.NewV4().String()
	messageRow := MessageEntity{
		ID:            message.ID,
		Content:       sql.NullString{String: message.Content, Valid: true},
		ToPhoneNumber: sql.NullString{String: message.ToPhoneNumber, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO messages
			(id, content, tophonenumber) 
			values 
			(:id, :content, :tophonenumber)`,
		messageRow,
	)

	if err != nil {
		return conversation.Message{}, fmt.Errorf(`Faild to insert message %w`, err)
	}

	if err := rows.Close(); err != nil {
		return conversation.Message{}, fmt.Errorf(`Faild to close messages %w`, err)
	}

	return message, nil
}

func (d *Database) DeleteMessage(ctx context.Context, uuid string) (bool, error) {
	_, err := d.Client.ExecContext(ctx, `DELETE FROM messages WHERE id = $1`, uuid)

	if err != nil {
		return false, fmt.Errorf(`Faild to delete message %w`, err)
	}

	return true, nil
}

func (d *Database) UpdateMessage(ctx context.Context, id string, message conversation.Message) (conversation.Message, error) {
	messageRow := MessageEntity{
		ID:            id,
		ToPhoneNumber: sql.NullString{String: message.ToPhoneNumber, Valid: true},
		Content:       sql.NullString{String: message.Content, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`
			UPDATE messages SET
				tophonenumber = :tophonenumber,
				content = :content
			WHERE id = :id
		`,
		messageRow,
	)

	if err != nil {
		return conversation.Message{}, fmt.Errorf(`Error updating message %w`, err)
	}

	if err := rows.Close(); err != nil {
		return conversation.Message{}, fmt.Errorf(`Error closing rows %w`, err)
	}

	return d.GetMessage(ctx, id)
}
