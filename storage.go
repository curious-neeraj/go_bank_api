package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	GetAccounts() ([]*Account, error)
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
		balance serial,
		ac_number serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

// implement functions to make interface valid
func (s *PostgresStore) CreateAccount(ac *Account) error {

	query := `INSERT INTO account 
	(first_name, last_name, balance, ac_number, created_at)
	VALUES
	($1, $2, $3, $4, $5)`

	resp, err := s.db.Query(query,
		ac.FirstName,
		ac.LastName,
		ac.Balance,
		ac.AcNumber,
		ac.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v \n", resp)
	return nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		if err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Balance,
			&account.AcNumber,
			&account.CreatedAt,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
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
