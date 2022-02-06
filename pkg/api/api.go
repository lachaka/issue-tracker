package api

import (
	"net/http"
	"issue-tracker/cmd/utils"
	
	"github.com/gocql/gocql"
)

type application struct {
	appLogger utils.AppLogger
}

func Init(config utils.Config) {
	app := &application {
		appLogger: utils.Logger,
	}
	
	session := app.ConnectCassandra(config.Db)
	defer session.Close()
	
	srv := &http.Server{
		Addr:     ":" + config.Host.Port,
		ErrorLog: app.appLogger.ErrorLog,
		// Handler:  app.routes(),
	}

	app.appLogger.InfoLog.Printf("Starting server on %s", config.Host.Port)
	err := srv.ListenAndServe()
	app.appLogger.ErrorLog.Fatal(err)
}

func (app *application) ConnectCassandra(config utils.CassandraConfig) (*gocql.Session)  {
	consistancy := func(c string) gocql.Consistency {
		gc, err := gocql.MustParseConsistency(c)
		if err != nil {
			app.appLogger.ErrorLog.Fatal(err)
		}

		return gc
	}
	
	cluster := gocql.NewCluster(config.Host + ":" + config.Port)
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = consistancy(config.Consistancy)

	session, err := cluster.CreateSession()

	if err != nil {
		app.appLogger.ErrorLog.Fatal(err)
	}

	app.appLogger.InfoLog.Printf("Database connected on port %d", cluster.Port)

	return session
}
