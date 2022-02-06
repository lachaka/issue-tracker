package cql

import (
	"issue-tracker/pkg/models"

	"github.com/gocql/gocql"
)

type UserModel interface {
	Save(user models.User) (*models.User, error)
	GetById(id string) (*models.User, error)
}

type userModel struct {
	session *gocql.Session
}

func NewUserModel(session *gocql.Session) UserModel {
	return &userModel{ session: session }
}

func (u *userModel) Save(user models.User) (*models.User, error) {
	var query string = "INSERT INTO user(id,email) VALUES(?,?)"

	if err := u.session.Query(query, user.Id, user.Email).Exec(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userModel) GetById(id string) (*models.User, error) {
	var user models.User

	var query string = `SELECT id, email FROM user where id=?`

	if err := u.session.Query(query, id).Scan(&user.Id, &user.Email); err != nil {

		if err == gocql.ErrNotFound {
			return nil, err
		}
		
		return nil, err
	}

	return &user, nil
}