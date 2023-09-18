package main

import (
	"log"
	"sync"
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

	// Create waitGroup
	var wg sync.WaitGroup

	for _, flow := range cfg.Flows {

		// Run each flow
		go flow.Run(&cfg, &wg, error_chan)
		wg.Add(1)

	}

	error_msg := <-error_chan
	log.Println(error_msg)

	wg.Wait()

}
