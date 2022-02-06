package utils

type Config struct {
	Db     CassandraConfig
	Server ServerSettings
}

type CassandraConfig struct {
	Host        string `mapstructure:"DB_HOST"`
	Port        string `mapstructure:"DB_PORT"`
	Keyspace    string `mapstructure:"DB_KEYSPACE"`
	Consistancy string `mapstructure:"DB_CONSISTANCY"`
}

type ServerSettings struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}
