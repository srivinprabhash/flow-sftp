package main

import (
	"log"

	"github.com/spf13/viper"
)

type FlowConfiguration struct {
	Flows []Flow `mapstcuture:"flows"`
}

func ReadConfig() (FlowConfiguration, error) {

	/*
		Read configuration
	*/
	viper.AddConfigPath(".")
	viper.SetConfigName("flow.yaml")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("ERROR READING CONFIG :: ", err)
		return FlowConfiguration{}, err
	}

	var configuration FlowConfiguration
	if err = viper.Unmarshal(&configuration); err != nil {
		log.Fatalln("ERROR UNMARHSALLING ::", err)
		return FlowConfiguration{}, err
	}

	return configuration, nil

}
