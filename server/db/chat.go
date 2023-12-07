package db

import (
	"database/sql"
	"errors"
)

var (
	NOTFOUND    = errors.New("Record not Found")
	UpdateFail  = errors.New("Record could not be updated")
	ErrorCreate = errors.New("Error during Creation of Record")
	GenError    = errors.New("Error on request")
)

type Message struct {
	Id      int64
	Content string
	Room    int64
	Author  string
	Length  int64
}

type SQLiteRepo struct {
	db *sql.DB
}

func newSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{
		db: db,
	}
}

func (c *SQLiteRepo) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS messages(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL UNIQUE,
        room TEXT NOT NULL,
        author INTEGER NOT NULL,
		length INTEGER NOT NULL,
	);
	`
	_, err := c.db.Exec(query)
	return err
}

func (c *SQLiteRepo) Create(msg Message) (*Message, error) {

	return nil, nil
}

func (c *SQLiteRepo) All() ([]Message, error) {

	return nil, nil
}

func (c *SQLiteRepo) GetByAuthor(author string) (*Message, error) {
	record, err := c.db.Query("SELECT * FROM Message WHERE author =?", author)
	if err != nil {
		return nil, GenError
	}
	var msg Message
	err = record.Scan(&msg.Id, &msg.Author, &msg.Content, &msg.Length, &msg.Room)
	if err != nil {
		return nil, NOTFOUND
	}

	return &msg, nil
}

func (c *SQLiteRepo) Update(id int64, update Message) (*Message, error) {

	return nil, nil
}

func (c *SQLiteRepo) Delete(id int64) error {
	_, err := c.db.Query("DELETE FROM Message WHERE id = ?", id)
	return err
}
