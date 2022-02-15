package api

import (
	utils "issue-tracker/cmd/utils"
	"issue-tracker/pkg/client/cassandra"
	handler "issue-tracker/pkg/handlers"
	"issue-tracker/pkg/middleware"
	"issue-tracker/pkg/models/cql"

	"github.com/gin-gonic/gin"
)

func Init(config utils.Config) {
	dbSession, err := cassandra.ConnectCassandra(config.Db)

	if err != nil {
		utils.Logger.ErrorLog.Fatal(err)
	}

	utils.Logger.InfoLog.Printf("Database connected on port %s", config.Db.Port)

	defer dbSession.Close()

	router := gin.Default()

	userRep := cql.NewUserModel(dbSession)
	userHandler := handler.NewUserHandler(&userRep)

	projRep := cql.NewProjectModel(dbSession)
	projectHandler := handler.NewProjectHandler(&projRep)

	router.POST("api/user/register", middleware.ValidateUserData(), userHandler.Register)
	router.POST("api/user/login", userHandler.Login)
	router.POST("api/user/logout", middleware.Authorization(dbSession), userHandler.Logout)
	router.GET("api/users", middleware.Authorization(dbSession),
		middleware.ValidateUserRole([]string{"admin"}),
		userHandler.GetUsers)

	router.POST("api/project", middleware.Authorization(dbSession),
		middleware.ValidateUserRole([]string{"admin"}),
		projectHandler.CreateProject)

	router.PUT("api/project/:id", middleware.Authorization(dbSession),
		middleware.ValidateUserRole([]string{"admin", "pm"}),
		projectHandler.UpdateProject)

	router.DELETE("api/project/:id", middleware.Authorization(dbSession),
		middleware.ValidateUserRole([]string{"admin"}),
		projectHandler.DeleteProject)

	router.GET("api/project/:id", middleware.Authorization(dbSession),
		projectHandler.GetProject)

	router.GET("api/projects", middleware.Authorization(dbSession),
		projectHandler.GetAllProjects)

	router.Run(config.Server.Host + ":" + config.Server.Port)
}
