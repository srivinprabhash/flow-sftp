package config

import (
	"log"

	"github.com/spf13/viper"
)

// type SFTPConnection struct {
// 	Host       string `mapstrucutre:"host"`
// 	Port       int    `mapstrucutre:"port"`
// 	User       string `mapstrucutre:"user"`
// 	PrivateKey string `mapstrucutre:"privateKey"`
// 	RemotePath string `mapstrucutre:"remotePath"`
// }

// type FlowConfiguration struct {
// 	Source        string `mapstrucutre:"source"`
// 	EnableBackups bool   `mapstrucutre:"enableBackups"`
// 	ClearBacklog  bool   `mapstructure:"clearBacklog"`
// 	Backups       string `mapstructure:"backups"`
// }

type Flow struct {
	Name            string `mapstructure:"name"`
	SourceDirectory string `mapstructure:"source_dir"`
	BackupDirectory string `mapstructure:"backup_dir"`
	EnableBackup    bool   `mapstructure:"enable_backup"`
	ClearBacklog    bool   `mapstructure:"clear_backlog"`
	RemoteHost      string `mapstructure:"remote_host"`
	RemotePort      int    `mapstructure:"remote_port"`
	RemoteUser      string `mapstructure:"remote_user"`
	PrivateKey      string `mapstructure:"private_key"`
	RemoteDirectory string `mapstructure:"remote_dir"`
}

type Configuration struct {
	Flows []Flow `mapstcuture:"flows"`
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
