package models

type User struct {
	Id        gocql.UUID   `cql:"id"`
	Firstname string       `cql:"firstname"`
	Lastname  string       `cql:"lastname"`
	Email     string       `cql:"email"`
	Password  string       `cql:"password"`
	Role      string       `cql:"role"`
	Projects  []gocql.UUID `cql:"projects"`
}