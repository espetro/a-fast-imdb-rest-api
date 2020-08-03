package config

import (
	"log"
	"strconv"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	APIPort int
	DBHost, DBPort string
)

func init() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	
	viper.SetDefault("api_port", "8080")
	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", "27017")

	// Search config in home directory and current directory with name ".cliApp.yaml"
	viper.SetConfigName(".apiApp")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading configuration file")
	}

	APIPort = viper.Get("api_port").(int)
	DBHost = viper.Get("db_host").(string)
	DBPort = strconv.Itoa(viper.Get("db_port").(int))
}