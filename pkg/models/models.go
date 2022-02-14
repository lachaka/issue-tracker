package models

import (
	"github.com/gocql/gocql"
)

type User struct {
	Id        gocql.UUID   `json:"id"`
	Firstname string       `json:"firstname"`
	Lastname  string       `json:"lastname"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Role      string       `json:"role"`
	Projects  []gocql.UUID `json:"projects"`
}

type Project struct {
	Id          gocql.UUID   `json:"id"`
	description string       `json:"description"`
	title       string       `json:"title"`
	tickets     []gocql.UUID `json:"tickets"`
	users       []gocql.UUID `json:"users"`
}
