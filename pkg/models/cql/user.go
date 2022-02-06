package cql

import (
	"issue-tracker/pkg/models"

	"github.com/gocql/gocql"
)

type UserModel struct {
	Session *gocql.Session
}

func (m *UserModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}

func (m *UserModel) Latest() ([]*models.User, error) {
	return nil, nil
}