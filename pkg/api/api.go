package api

import (
	"net/http"
	"issue-tracker/cmd/utils"
	"issue-tracker/pkg/models/cql"
	"issue-tracker/pkg/client/cassandra"
	
)

type application struct {
	appLogger utils.AppLogger
	users     *cql.UserModel
}

func Init(config utils.Config) {
	app := &application {
		appLogger: utils.Logger,
		// users:      cql.UserModel{DB: session},
	}
	
	session, err := cassandra.ConnectCassandra(config.Db)

	if err != nil {
		app.appLogger.ErrorLog.Fatal(err)
	}
	
	app.appLogger.InfoLog.Printf("Database connected on port %s", config.Db.Port)
	
	defer session.Close()
	
	srv := &http.Server{
		Addr:     ":" + config.Host.Port,
		ErrorLog: app.appLogger.ErrorLog,
		// Handler:  app.routes(),
	}

	app.appLogger.InfoLog.Printf("Starting server on %s", config.Host.Port)
	err = srv.ListenAndServe()
	app.appLogger.ErrorLog.Fatal(err)
}