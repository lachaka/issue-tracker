package cql

import (
	"encoding/json"
	"errors"
	"issue-tracker/pkg/models"

	"github.com/gocql/gocql"
)

type ProjectModel interface {
	CreateProject(project models.Project) (*models.Project, error)
	UpdateProject(project models.Project) error
	DeleteProject(id gocql.UUID) error
	GetProject(id gocql.UUID) (*models.Project, error)
	GetAllProjects() ([]map[string]interface{}, error)
}

type projectModel struct {
	session *gocql.Session
}

func NewProjectModel(session *gocql.Session) ProjectModel {
	return &projectModel{session: session}
}

func (p *projectModel) CreateProject(project models.Project) (*models.Project, error) {
	var query string = `INSERT INTO project JSON ?`

	jsonString, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}

	if err := p.session.Query(query, jsonString).Exec(); err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *projectModel) UpdateProject(project models.Project) error {
	var query string = `UPDATE project SET title=?, description=? WHERE id=?`

	err := p.session.Query(query, project.Title, project.Description, project.Id).Exec()

	return err
}

func (p *projectModel) DeleteProject(id gocql.UUID) error {
	var query string = `DELETE from project WHERE id=?`

	err := p.session.Query(query, id).Exec()

	return err
}

func (p *projectModel) GetProject(id gocql.UUID) (*models.Project, error) {
	var query string = `SELECT * FROM project where id=?`

	res := make(map[string]interface{})

	err := p.session.Query(query, id).MapScan(res)
	if err != nil {
		return nil, errors.New("Project not found")
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	var project models.Project
	if err := json.Unmarshal(jsonStr, &project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *projectModel) GetAllProjects() ([]map[string]interface{}, error) {
	var query string = `SELECT * FROM project`

	iter := p.session.Query(query).Iter()
	defer iter.Close()

	ret := []map[string]interface{}{}
	m := &map[string]interface{}{}

	for iter.MapScan(*m) {
		ret = append(ret, *m)
		m = &map[string]interface{}{}
	}

	return ret, nil
}
