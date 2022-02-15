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

	repository := cql.NewUserModel(dbSession)
	userHandler := handler.NewUserHandler(&repository)

	router.POST("api/user/register", middleware.ValidateUserData(), userHandler.Register)
	router.POST("api/user/login", userHandler.Login)
	router.POST("api/user/logout", middleware.Authorization(dbSession), userHandler.Logout)
	router.GET("api/users", middleware.Authorization(dbSession), middleware.ValidateUserRole("admin"), userHandler.GetUsers)

	router.Run(config.Server.Host + ":" + config.Server.Port)
}
