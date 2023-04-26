package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", "file:image_db.sqlite?cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}

	err = createSchema(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createSchema(db *sqlx.DB) error {
	schema := `
CREATE TABLE images (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	file_name TEXT NOT NULL,
	data BLOB NOT NULL,
	thumbnail BLOB NOT NULL,
	width INTEGER NOT NULL,
	height INTEGER NOT NULL,
	size INTEGER NOT NULL,
	content_type INTEGER NOT NULL
);
	

CREATE TABLE global_tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    tag_type INTEGER NOT NULL
);

CREATE TABLE image_tags (
    image_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    PRIMARY KEY (image_id, tag_id),
    FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES global_tags (id) ON DELETE CASCADE
);
`
	_, err := db.Exec(schema)
	return err
}
