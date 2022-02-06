package cassandra

import (	
	"issue-tracker/cmd/utils"

	"github.com/gocql/gocql"
)

func ConnectCassandra(config utils.CassandraConfig) (*gocql.Session, error)  {
	consistancy, err := gocql.MustParseConsistency(config.Consistancy)
	
	if err != nil {
		return nil, err
	}
	
	cluster := gocql.NewCluster(config.Host + ":" + config.Port)
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = consistancy

	session, err := cluster.CreateSession()

	return session, err
}
