package main

import (
	"log"
)

// func init() {

// 	// config, err := config.ReadConfig()
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// fmt.Println(config)

// 	// os.Exit(0)

// 	// // Read Configurations
// 	// config, err := config.ReadConfig()
// 	// if err != nil {
// 	// 	log.Fatalln("ERROR (Reading configurations from init):: ", err)
// 	// }
// 	// cfg = config

// 	// // Check if the path exists
// 	// source, err := os.Open(cfg.Flow.Source)
// 	// if err != nil {
// 	// 	log.Fatalln("ERROR :: ", err)
// 	// }

// 	// // Get source information
// 	// sourceInfo, err := source.Stat()
// 	// if err != nil {
// 	// 	log.Fatalln("ERROR :: ", err)
// 	// }

// 	// // Check if source is a dir
// 	// if !sourceInfo.IsDir() {
// 	// 	log.Fatalln("ERROR :: Source is not a directory")
// 	// }
// }

func main() {

	// Read configuration
	var cfg, err = ReadConfig()
	if err != nil {
		log.Fatalln("[ERROR] :: ", err)
	}

	// Make Run error channel
	error_chan := make(chan error)

	for _, flow := range cfg.Flows {

		thisFlow := flow

		// Run each flow
		go thisFlow.Run(error_chan)

	}

	error_msg := <-error_chan
	log.Println(error_msg)

}
