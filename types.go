package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Number    int64      `json:"number"`
	Balance   int64      `json:"balance"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		Number: int64(rand.Intn(1000000)),
		Balance: 0,
	}
}