package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/srivinprabhash/flow-sftp/config"
	send "github.com/srivinprabhash/flow-sftp/sftp"
)

var cfg config.Configuration

func init() {

	// Read Configurations
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatalln("ERROR (Reading configurations from init):: ", err)
	}
	cfg = config

	// Check if the path exists
	source, err := os.Open(cfg.Flow.Source)
	if err != nil {
		log.Fatalln("ERROR :: ", err)
	}

	// Get source information
	sourceInfo, err := source.Stat()
	if err != nil {
		log.Fatalln("ERROR :: ", err)
	}

	// Check if source is a dir
	if !sourceInfo.IsDir() {
		log.Fatalln("ERROR :: Source is not a directory")
	}
}

/*
Process existing files
*/
func processExistingFiles(cfg *config.Configuration) {

	// List files in the dir
	files, err := ioutil.ReadDir(cfg.Flow.Source)
	if err != nil {
		log.Fatalln("ERROR :: Could not list files in the directory :: ", err)
	}

	// Loop through file list
	for _, file := range files {

		// Concatnate file path
		fp := cfg.Flow.Source + file.Name()

		// Send File to SFTP Server
		err = send.Send(fp, cfg)
		if err != nil {
			log.Fatalln("ERROR :: Could not send :: ", err)
		}
		log.Println("Successfully moved file :: ", fp)

		/*
			TODO :: Implement Backup mechanism
		*/

	}

}

func main() {

	// Read configuration
	var config, err = config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Start processing existing files in background
	if cfg.Flow.ClearBacklog {
		go processExistingFiles(&config)
	}

	// Setup the directory watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
	}

	// Start watching
	err = watcher.Add(cfg.Flow.Source)
	if err != nil {
		log.Println(err)
	}
	log.Println("Watching directory :: ", cfg.Flow.Source)

	for {
		select {
		case event := <-watcher.Events:

			if event.Op&fsnotify.Create == fsnotify.Create {

				// Send File to SFTP Server
				err = send.Send(event.Name, &config)
				if err != nil {
					log.Fatalln("ERROR :: Could not send :: ", err)
				}
				log.Println("Successfully moved file :: ", event.Name)

				/*
					TODO :: Implement Backup Mechanism
				*/

			}

		case err := <-watcher.Errors:
			log.Println(err)
		}
	}

}
