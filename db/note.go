package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/pgtype"
)

type Note struct {
	Id         pgtype.UUID
	Taker      string
	About      string
	Content    string
	CreateDate time.Time
}

func NewNote(taker, about, content string) *Note {
	return &Note{
		pgtype.UUID{},
		taker,
		about,
		content,
		time.Now(),
	}
}

func (note *Note) Save() error {
	_, err := db.Exec(context.Background(), "INSERT INTO note(id, taker, about, content, create_date) VALUES(uuid_generate_v4(), $1, $2, $3, $4)", note.Taker, note.About, note.Content, note.CreateDate)
	return err
}

func GetNotes(about string) ([]Note, error) {
	var res []Note
	rows, err := db.Query(context.Background(), "SELECT id, taker, about, content, create_date FROM note WHERE about=$1 ORDER BY create_date DESC LIMIT 25", &about)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note Note
		err = rows.Scan(&note.Id, &note.Taker, &note.About, &note.Content, &note.CreateDate)
		if err != nil {
			return nil, err
		}
		res = append(res, note)
	}

	return res, nil
}
