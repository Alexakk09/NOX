package repository

import (
	"context"

	"veerai/internal/database"
)

func SaveMemory(userID int64, key string, value string) error {
	_, err := database.DB.Exec(
		context.Background(),
		`INSERT INTO memories(user_id, key, value)
		 VALUES($1, $2, $3)`,
		userID,
		key,
		value,
	)

	return err
}

func GetMemories(userID int64) (map[string]string, error) {
	rows, err := database.DB.Query(
		context.Background(),
		`SELECT key, value
		 FROM memories
		 WHERE user_id = $1
		 ORDER BY id ASC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	memories := make(map[string]string)

	for rows.Next() {
		var key string
		var value string

		err := rows.Scan(&key, &value)
		if err != nil {
			return nil, err
		}

		memories[key] = value
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return memories, nil
}