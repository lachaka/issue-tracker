package cql

import (
	"issue-tracker/pkg/models"

	"github.com/gocql/gocql"
)

type ProjectModel interface {
	Create(proejct models.Project) (*models.Project, error)
}

type projectModel struct {
	session *gocql.Session
}

func NewProjectModel(session *gocql.Session) ProjectModel {
	return &projectModel{session: session}
}

func (p *projectModel) Create(project models.Project) (*models.Project, error) {
	return nil, nil
}
