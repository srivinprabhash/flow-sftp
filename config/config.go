package config

import (
	"log"

	"github.com/spf13/viper"
)

type SFTPConnection struct {
	Host       string `mapstrucutre:"host"`
	Port       int    `mapstrucutre:"port"`
	User       string `mapstrucutre:"user"`
	PrivateKey string `mapstrucutre:"privateKey"`
	RemotePath string `mapstrucutre:"remotePath"`
}

type FlowConfiguration struct {
	Source          string `mapstrucutre:"source"`
	enableBackups   bool   `mapstrucutre:"enableBackups"`
	BackupDirectory string `mapstrucutre:"backupDir"`
}

type Configuration struct {
	SFTPConnection SFTPConnection    `mapstructure:"sftp"`
	Flow           FlowConfiguration `mapstcuture:"flow"`
}

func ReadConfig() (Configuration, error) {

	/*
		Read configuration
	*/
	viper.AddConfigPath(".")
	viper.SetConfigName("flow.yaml")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("ERROR READING CONFIG :: ", err)
		return Configuration{}, err
	}

	var configuration Configuration
	if err = viper.Unmarshal(&configuration); err != nil {
		log.Fatalln("ERROR UNMARHSALLING ::", err)
		return Configuration{}, err
	}

	return configuration, nil

}
