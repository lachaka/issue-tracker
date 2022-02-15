package handler

import (
	"issue-tracker/cmd/utils"
	"issue-tracker/pkg/models"
	"issue-tracker/pkg/models/cql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type ProjectHandler interface {
	CreateProject(*gin.Context)
	UpdateProject(*gin.Context)
	DeleteProject(*gin.Context)
	GetProject(*gin.Context)
	GetAllProjects(*gin.Context)
}

type projectHandler struct {
	projectModel cql.ProjectModel
}

func NewProjectHandler(projectModel *cql.ProjectModel) ProjectHandler {
	return &projectHandler{projectModel: *projectModel}
}

func (p *projectHandler) CreateProject(c *gin.Context) {
	var user models.User
	getUser := c.MustGet("user")
	user = *getUser.(*models.User)

	var project models.Project
	c.BindJSON(&project)
	project.Id = gocql.TimeUUID()
	project.Users = append(project.Users, user.Id)

	data, err := p.projectModel.CreateProject(project)

	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func (p *projectHandler) UpdateProject(c *gin.Context) {
	var project models.Project
	c.BindJSON(&project)
	id, err := gocql.ParseUUID(c.Param("id"))

	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	project.Id = id

	err = p.projectModel.UpdateProject(project)

	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "Project updated"})
	}
}

func (p *projectHandler) DeleteProject(c *gin.Context) {
	id, err := gocql.ParseUUID(c.Param("id"))

	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	err = p.projectModel.DeleteProject(id)

	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "Project deleted"})
	}
}

func (p *projectHandler) GetProject(c *gin.Context) {
	id, err := gocql.ParseUUID(c.Param("id"))

	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	project, err := p.projectModel.GetProject(id)

	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, project)
	}
}

func (p *projectHandler) GetAllProjects(c *gin.Context) {
	data, err := p.projectModel.GetAllProjects()

	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, data)
}
