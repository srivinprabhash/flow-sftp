package main

import (
	"fmt"
	"log"
	"os"

	"github.com/srivinprabhash/flow-sftp/config"
)

var cfg config.Configuration

func init() {

	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(config)

	os.Exit(0)

	// // Read Configurations
	// config, err := config.ReadConfig()
	// if err != nil {
	// 	log.Fatalln("ERROR (Reading configurations from init):: ", err)
	// }
	// cfg = config

	// // Check if the path exists
	// source, err := os.Open(cfg.Flow.Source)
	// if err != nil {
	// 	log.Fatalln("ERROR :: ", err)
	// }

	// // Get source information
	// sourceInfo, err := source.Stat()
	// if err != nil {
	// 	log.Fatalln("ERROR :: ", err)
	// }

	// // Check if source is a dir
	// if !sourceInfo.IsDir() {
	// 	log.Fatalln("ERROR :: Source is not a directory")
	// }
}

// func backupFile(fp string, cfg *config.Configuration) error {

// 	// Get abs paths
// 	sourcePath := fp
// 	destination := cfg.Flow.Backups
// 	fileName := filepath.Base(sourcePath)
// 	destinationFilePath := filepath.Join(destination, fileName)

// 	// Open the source file
// 	sourceFile, err := os.Open(sourcePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer sourceFile.Close()

// 	// Create destination file
// 	destinationFile, err := os.Create(destinationFilePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer destinationFile.Close()

// 	// Copy content
// 	_, err = io.Copy(destinationFile, sourceFile)
// 	if err != nil {
// 		return err
// 	}

// 	// Remove source
// 	err = os.Remove(sourcePath)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

/*
Process existing files
*/
// func processExistingFiles(cfg *config.Configuration) {

// 	// List files in the dir
// 	files, err := ioutil.ReadDir(cfg.Flow.Source)
// 	if err != nil {
// 		log.Fatalln("ERROR :: Could not list files in the directory :: ", err)
// 	}

// 	// Loop through file list
// 	for _, file := range files {

// 		// Concatnate file path
// 		fp := cfg.Flow.Source + file.Name()

// 		// // Send File to SFTP Server
// 		err = send.Send(fp, cfg)
// 		if err != nil {
// 			log.Fatalln("ERROR :: Could not send :: ", err)
// 		}

// 		// Check if backup is enabled
// 		if cfg.Flow.EnableBackups {
// 			err := backupFile(fp, cfg)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 			log.Println("File backed up successfully :: ", fp)
// 		}

// 	}

// }

func main() {

	// Read configuration
	// var config, err = config.ReadConfig()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // Start processing existing files in background
	// if cfg.Flow.ClearBacklog {
	// 	go processExistingFiles(&config)
	// }

	// // Setup the directory watcher
	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	// 	log.Println(err)
	// }

	// // Start watching
	// err = watcher.Add(cfg.Flow.Source)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("Watching directory :: ", cfg.Flow.Source)

	// for {
	// 	select {
	// 	case event := <-watcher.Events:

	// 		if event.Op&fsnotify.Create == fsnotify.Create {

	// 			// Send File to SFTP Server
	// 			err = send.Send(event.Name, &config)
	// 			if err != nil {
	// 				log.Fatalln("ERROR :: Could not send :: ", err)
	// 			}
	// 			log.Println("Successfully moved file :: ", event.Name)

	// 			// Backing up file
	// 			if cfg.Flow.EnableBackups {
	// 				err := backupFile(event.Name, &config)
	// 				if err != nil {
	// 					log.Fatalln(err)
	// 				}
	// 				log.Println("File backed up successfully :: ", event.Name)
	// 			}

	// 		}

	// 	case err := <-watcher.Errors:
	// 		log.Println(err)
	// 	}
	// }

}
