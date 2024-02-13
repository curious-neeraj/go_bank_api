package main

// list of packages imported
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// storage interface
type Storage interface {
	CreateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
	UpdateAccount(int) error
	DeleteAccount(int) error
}

// database store
type PostgresStore struct {
	db *sql.DB
}

// postgres store connection
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

// initialize db to contain account table
func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

// execute create table query to initialize db
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

// get list of all accounts and corresponding data
func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

// get account details based on `id`
func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	row, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		return scanAccount(row)
	}
	return nil, fmt.Errorf("account with `id = %d` not found", id)
}

func (s *PostgresStore) UpdateAccount(id int) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	if err != nil {
		return err
	}

	return fmt.Errorf("account with `id = %d` not found", id)
}

func scanAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Balance,
		&account.AcNumber,
		&account.CreatedAt,
	)

	return account, err
}
