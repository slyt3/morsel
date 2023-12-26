package models

import (
	"database/sql"
	"time"
)

type Morsel struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type MorselModel struct {
	DB *sql.DB
}

func (m *MorselModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (m *MorselModel) Get(id int) (*Morsel, error) {
	return nil, nil
}

func (m *MorselModel) Latest() ([]*Morsel, error) {
	return nil, nil
}
