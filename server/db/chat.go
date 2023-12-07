package db

import "database/sql"

type Message struct {
	Id      int64
	Content string
	Room    int64
	Author  string
	Length  int64
}

// sql.DB is pool of connections. sql.DB comes from the standard libary. (dependig on driver)
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
