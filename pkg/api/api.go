package api

import (
	utils "issue-tracker/cmd/utils"
	"issue-tracker/pkg/client/cassandra"
	"issue-tracker/pkg/handlers"
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


	router.POST("api/user/", userHandler.CreateUser)
	router.GET("api/user/:id", userHandler.GetUserById)

    router.Run(config.Server.Host + ":" + config.Server.Port)
}