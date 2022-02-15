package cql

import (
	"encoding/json"
	"errors"
	"issue-tracker/pkg/models"

	"github.com/gocql/gocql"
)

type UserModel interface {
	Save(user models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetUsers() ([]map[string]interface{}, error)
}

type userModel struct {
	session *gocql.Session
}

func NewUserModel(session *gocql.Session) UserModel {
	return &userModel{session: session}
}

func (u *userModel) registeredEmail(email string) (bool, error) {
	var query string = `SELECT COUNT(*) FROM user WHERE email=? ALLOW FILTERING`
	var emailCount int

	if err := u.session.Query(query, email).Scan(&emailCount); err != nil {
		if err == gocql.ErrNotFound {
			return false, err
		}

		return false, err
	}

	if emailCount != 0 {
		return true, errors.New("Email already registered")
	}

	return false, nil
}

func (u *userModel) Save(user models.User) (*models.User, error) {
	_, err := u.registeredEmail(user.Email)

	if err != nil {
		return nil, err
	}

	var query string = `INSERT INTO user JSON ?`

	jsonString, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	if err := u.session.Query(query, jsonString).Exec(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userModel) GetByEmail(email string) (*models.User, error) {
	return GetByEmail(email, u.session)
}

func GetByEmail(email string, session *gocql.Session) (*models.User, error) {
	var user models.User

	var query string = `SELECT * FROM user where email=? ALLOW FILTERING`

	res := make(map[string]interface{})
	err := session.Query(query, email).MapScan(res)
	if err != nil {
		return nil, errors.New("User not found")
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonStr, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userModel) GetUsers() ([]map[string]interface{}, error) {
	var query string = `SELECT id, email, firstname, lastname, projects, role FROM user`

	iter := u.session.Query(query).Iter()
	defer iter.Close()

	ret := []map[string]interface{}{}
	m := &map[string]interface{}{}
	for iter.MapScan(*m) {
		ret = append(ret, *m)
		m = &map[string]interface{}{}
	}

	return ret, nil
}
