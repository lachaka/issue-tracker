package main

import (
	"issue-tracker/cmd/utils"
	"issue-tracker/pkg/api"

	"github.com/spf13/viper"
)

func main() {
	config := readConfig(".")
	
	api.Init(config)
}

func readConfig(path string) utils.Config {
    viper.AddConfigPath(path)
    viper.SetConfigName(".env")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
		utils.Logger.ErrorLog.Fatal(err)
    }

	var config utils.Config
	unmarshalConfig(&config.Db)
	unmarshalConfig(&config.Server)

	utils.Logger.InfoLog.Println("File .env is read correctly")
    
	return config
}

func unmarshalConfig(config interface{}) {
	err := viper.Unmarshal(&config)
	if err != nil {
		utils.Logger.ErrorLog.Fatal(err)
	}
}