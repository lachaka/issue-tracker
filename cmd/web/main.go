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

    err = viper.Unmarshal(&config.Db)
	if err != nil {
		utils.Logger.ErrorLog.Fatal(err)
	}
	
    err = viper.Unmarshal(&config.Host)
	if err != nil {
		utils.Logger.ErrorLog.Fatal(err)
	}

	utils.Logger.InfoLog.Println("File .env is read correctly")
    
	return config
}