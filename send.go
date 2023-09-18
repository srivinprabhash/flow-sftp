package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func getConnection(flow *Flow) (*ssh.Client, error) {

	// Read Config
	// var config, err = config.ReadConfig()
	// if err != nil {
	// 	return nil, err
	// }

	// Read private key file
	pk, err := ioutil.ReadFile(flow.PrivateKey)
	if err != nil {
		return nil, err
	}

	// Parse the private key
	signer, err := ssh.ParsePrivateKey(pk)
	if err != nil {
		return nil, err
	}

	// SSH Configurtions
	sshConfig := &ssh.ClientConfig{
		User:            flow.RemoteUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}

	// Dial the SSH Connection
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", flow.RemoteHost, flow.RemotePort), sshConfig)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Send(fp string, flow *Flow) error {

	// Get a SSH connection
	conn, err := getConnection(flow)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create a SFTP Client
	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}
	defer client.Close()

	// Upload file
	rfn := filepath.Base(fp)
	rfp := filepath.Join(flow.RemoteDirectory, rfn)
	rf, err := client.Create(rfp)
	if err != nil {
		return err
	}
	defer rf.Close()

	// Get file content
	ct, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}

	// Write file content to remote file
	_, err = rf.Write(ct)
	if err != nil {
		return err
	}

	return nil
}
