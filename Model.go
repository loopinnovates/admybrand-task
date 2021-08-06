package main

import (
	"time"
)

// Struct for User
type User struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        string    `json:name`
	Dob         string    `json:dob`
	Address     string    `json:address`
	Description string    `json:description`
	CreatedAt   time.Time `bson:"omitempty" json:createdAt`
}
