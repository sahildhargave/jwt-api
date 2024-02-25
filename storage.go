package main

import (
	"database/sql"
   _ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}
//
func NewPostgresStore() (*PostgresStore, error){
connStr := "user=postgres password=postgres123 dbname=postgres sslmode=disable"
db, err := sql.Open("postgres", connStr)
if err != nil {
	return nil, err
}
if err := db.Ping();
err != nil{
	return nil,err
}
return &PostgresStore{
	db: db,
},
nil
}

func (s *PostgresStore) CreateAccount(*Account) error{
	return nil
}

func (s *PostgresStore) DeleteAccount(int) error{
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error{
	return nil
}

func(s *PostgresStore) CreateAccount(*Account) errror{
	return nil
}

