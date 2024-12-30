package sqlite

import (
	"database/sql"

	"github.com/ChathuraD1392/gotest/internal/models"
)

type PostModels struct {
	DB *sql.DB
}

func (m *PostModels) Insert(title, content string) error {
	statement := "INSERT INTO posts(title,content,createdAt) VALUES(?,?,datetime('now'))"
	_, err := m.DB.Exec(statement, title, content)
	return err
}

func (m *PostModels) All() ([]models.Post, error) {
	statement := "SELECT id,title,content,createdAt FROM posts ORDER BY id DESC"

	rows, err := m.DB.Query(statement)

	if err != nil {
		return nil, err
	}

	posts := []models.Post{}
	for rows.Next() {
		p := models.Post{}
		err := rows.Scan(&p.Id, &p.Title, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
