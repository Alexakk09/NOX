package repository

import (
	"context"
	"time"

	"veerai/internal/database"
	"veerai/internal/models"
)

type Conversation struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateConversation(userID int64) (int64, error) {
	var id int64

	err := database.DB.QueryRow(
		context.Background(),
		`INSERT INTO conversations(user_id)
		 VALUES($1)
		 RETURNING id`,
		userID,
	).Scan(&id)

	return id, err
}

func AddMessage(conversationID int64, message models.Message) error {
	_, err := database.DB.Exec(
		context.Background(),
		`INSERT INTO messages(conversation_id, role, content)
		 VALUES($1, $2, $3)`,
		conversationID,
		message.Role,
		message.Content,
	)

	return err
}

func GetMessages(conversationID int64) ([]models.Message, error) {
	rows, err := database.DB.Query(
		context.Background(),
		`SELECT role, content
		 FROM messages
		 WHERE conversation_id = $1
		 ORDER BY id DESC
		 LIMIT 20`,
		conversationID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		var message models.Message

		err := rows.Scan(
			&message.Role,
			&message.Content,
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

func GetConversationUserID(conversationID int64) (int64, error) {
	var userID int64

	err := database.DB.QueryRow(
		context.Background(),
		`SELECT user_id
		 FROM conversations
		 WHERE id = $1`,
		conversationID,
	).Scan(&userID)

	return userID, err
}

func GetConversations(userID int64) ([]Conversation, error) {
	rows, err := database.DB.Query(
		context.Background(),
		`SELECT id, title, created_at
		 FROM conversations
		 WHERE user_id = $1
		 ORDER BY created_at DESC`,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var conversations []Conversation

	for rows.Next() {
		var conversation Conversation

		err := rows.Scan(
			&conversation.ID,
			&conversation.Title,
			&conversation.CreatedAt,

		)

		if err != nil {
			return nil, err
		}

		conversations = append(conversations, conversation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return conversations, nil
}

func DeleteConversation(conversationID int64, userID int64) error {
	_, err := database.DB.Exec(
		context.Background(),
		`DELETE FROM conversations
		 WHERE id = $1 AND user_id = $2`,
		conversationID,
		userID,
	)

	return err
}

func UpdateConversationTitle(conversationID int64, title string) error {
	_, err := database.DB.Exec(
		context.Background(),
		`UPDATE conversations
		 SET title = $1
		 WHERE id = $2`,
		title,
		conversationID,
	)

	return err
}
