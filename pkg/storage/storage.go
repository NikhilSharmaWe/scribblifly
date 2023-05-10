package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	Create(interface{}) error
	Delete(username string) error
	Update(oldRecord, newRecord *Account) error
	GetAll() ([]*Account, error)
	Get(username string) (*Account, error)
}

func NewPostgresStore() (*sql.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("DB_NAME")
	pass := os.Getenv("POSTGRES_PASSWORD")

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbName, pass))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Init(db *sql.DB) error {
	err := createAccountsTable(db)
	if err != nil {
		return err
	}

	err = createScriptTable(db)
	return err
}

func createAccountsTable(db *sql.DB) error {
	query := `create table if not exists accounts (
		username varchar(100),
		first_name varchar(100),
		last_name varchar(100),
		email varchar(100),
		about varchar(1000),
		encrypted_password varchar(500),
		created_at timestamp
	)`

	_, err := db.Exec(query)
	return err
}

func createScriptTable(db *sql.DB) error {
	query := `create table if not exists scripts (
		title varchar(100) primary key,
		username varchar(100),
		type varchar(100),
		content text,
		created_at timestamp
	)`

	_, err := db.Exec(query)
	return err
}
