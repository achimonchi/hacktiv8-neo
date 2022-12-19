package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	SSH_Address  = "103.23.198.176:22"
	SSH_Username = "hactiv8"
	SSH_Password = "_5D9aBZ6K2h3Td4"
)

func main() {
	sshConfig := &ssh.ClientConfig{
		User:            SSH_Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_Password),
		},
	}

	client, err := ssh.Dial("tcp", SSH_Address, sshConfig)
	if err != nil {
		log.Println(err)
		return
	}

	defer client.Close()

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Println(err)
		return
	}

	defer sftpClient.Close()

	destination, err := sftpClient.Create("/home/hactiv8/reyhan/index.html")
	if err != nil {
		log.Println(err)
		return
	}

	source, err := os.Open("./index.html")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = io.Copy(destination, source)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("File copied")
}
