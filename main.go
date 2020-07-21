package main

import (
	"time"
)

type (
	Customer struct {
		ID        string    `json:"id"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Phone     string    `json:"phone"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"createdAt"`
		UpdateAt  time.Time `json:"updatedAt"`
	}
)

func main() {
}
