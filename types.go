package main

import "math/rand"

type Account struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Balance   int64  `json:"balance"`
	AcNumber  int64  `json:"acNumber"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        int64(rand.Intn(10000)),
		FirstName: firstName,
		LastName:  lastName,
		AcNumber:  int64(rand.Intn(10000)),
	}
}
