package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	GetAccountById(int) (*Account, error)
	UpdateAccount(int) error
	DeleteAccount(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=bankAdmin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial primary key,
		first_name varchar(30),
		last_name varchar(30),
		balance integer(10),
		ac_number integer(10),
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

// implement functions to make interface valid
func (s *PostgresStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) UpdateAccount(id int) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(int) error {
	return nil
}
