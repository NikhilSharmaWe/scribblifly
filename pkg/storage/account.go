package storage

import (
	"database/sql"
	"fmt"
	"time"
)

type AccountStorage interface {
	Create(Account) error
	Delete(username string) error
	Update(oldRecord, newRecord Account) error
	GetAll() ([]Account, error)
	Get(username string) (Account, error)
}

type AccountModel struct {
	DB *sql.DB
}

type Account struct {
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	About             string    `json:"about"`
	EmailAddr         string    `json:"email"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
}

func (s *AccountModel) Create(acc Account) error {

	query := `insert into accounts
	(username, first_name, last_name, email, about, encrypted_password, created_at)
	values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.DB.Exec(query, acc.Username, acc.FirstName, acc.LastName, acc.EmailAddr, acc.About, acc.EncryptedPassword, acc.CreatedAt)
	return err
}

func (s *AccountModel) Delete(username string) error {
	query := `delete from accounts where username = $1`

	_, err := s.DB.Exec(query, username)
	return err
}

func (s *AccountModel) Update(old, new Account) error {
	query := `update accounts 
	set first_name = $2, last_name = $3, email = $4, about = $5
	WHERE username = $1`

	_, err := s.DB.Exec(query, old.Username, new.FirstName, new.LastName, new.EmailAddr, new.About)
	return err
}

func (s *AccountModel) GetAll() ([]Account, error) {
	query := `select * from accounts`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	accounts := []Account{}

	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (s *AccountModel) Get(username string) (Account, error) {
	acc := Account{}
	query := `select username, first_name, last_name, email, from accounts where username = $1`
	err := s.DB.QueryRow(query, username).Scan(&acc.Username, &acc.FirstName, &acc.LastName, &acc.EmailAddr)
	if err != nil {
		return acc, err
	}

	return acc, nil
}

func scanIntoAccount(rows *sql.Rows) (Account, error) {
	account := Account{}
	if err := rows.Scan(
		&account.Username,
		&account.FirstName,
		&account.LastName,
		&account.EmailAddr,
		&account.About); err != nil {
		return account, err
	}
	fmt.Println(account)
	return account, nil
}
