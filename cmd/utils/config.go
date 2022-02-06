package utils

type Config struct {
	Db   CassandraConfig
	Host HostSettings
}

type CassandraConfig struct {
	Host        string `mapstructure:"DB_HOST"`
	Port        string `mapstructure:"DB_PORT"`
	Keyspace    string `mapstructure:"DB_KEYSPACE"`
	Consistancy string `mapstructure:"DB_CONSISTANCY"`
}

type HostSettings struct {
	Port string `mapstructure:"PORT"`
}
