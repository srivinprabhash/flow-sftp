package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

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

// Runs the flow.
//
// Clear backlog of files to be moved if any exists.
// Then starts watching for a changes in the
// FlowConfiguration's SourceDirectory. Once a new file
// is found, it sends to the SFTP host and backup the file
// if it is enabled.
func (f *Flow) Run(messages chan error) {

	log.Println("[INFO] :: Start running flow :: ", f.Name)

	// Clear Backlog
	if f.ClearBacklog {
		if err := f.clearBacklog(); err != nil {
			messages <- err
		}
	}

	// Setup the directory watcher
	w, err := fsnotify.NewWatcher()
	if err != nil {
		messages <- err
	}
	defer w.Close()

	// Start watching
	err = w.Add(f.SourceDirectory)
	if err != nil {
		messages <- err
	}

	log.Println("[INFO] :: Watching directory ::", f.SourceDirectory)

	go func() {
		for {
			select {
			case event := <-w.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {

					// New file identified
					log.Println("[INFO] :: New file detected.", event.Name)
					// Send file to the SFTP Server
					err = Send(event.Name, f)
					if err != nil {
						messages <- err
					}
					log.Println("[INFO] :: Successfully moved file :: ", event.Name)

					// Backing Up File
					if f.EnableBackup {
						if err := f.backupFile(event.Name); err != nil {
							messages <- err
						}
					}

				}
			}
		}
	}()

	// Block main goroutine
	<-make(chan struct{})
}

// Clear the back log of files in the source directory.
func (f *Flow) clearBacklog() error {

	log.Println("[INFO] :: Cleaning process started for :: ", f.SourceDirectory)

	// List files in the source directory
	files, err := ioutil.ReadDir(f.SourceDirectory)
	if err != nil {
		return err
	}

	// Loop through files in the source directory
	for _, file := range files {

		// Concat file path
		fp := f.SourceDirectory + file.Name()

		// // Send file to SFTP Server
		if err := Send(fp, f); err != nil {
			return err
		}

		// Backing Up File
		if f.EnableBackup {
			err = f.backupFile(fp)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// Backs up the file
//
// Moves the file to a backup location
// Takes the source file's path as a parameter
func (f *Flow) backupFile(sfp string) error {

	fileName := filepath.Base(sfp)
	destinationFilePath := filepath.Join(f.BackupDirectory, fileName)

	// Open Source file
	sContent, err := os.Open(sfp)
	if err != nil {
		return err
	}
	defer sContent.Close()

	// Create backup
	df, err := os.Create(destinationFilePath)
	if err != nil {
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sContent)
	if err != nil {
		return err
	}

	// delete source
	err = os.Remove(sfp)
	if err != nil {
		return err
	}

	log.Println("[INFO] :: File backed up successfully :: ", destinationFilePath)

	return nil

}
