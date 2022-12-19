package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

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

	sess, err := client.NewSession()
	if err != nil {
		log.Println(err)
		return
	}

	defer sess.Close()

	// sess.Stdin = os.Stdin
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	stdin, err := sess.StdinPipe()
	if err != nil {
		log.Println(err)
		return
	}

	err = sess.Start("/bin/bash")
	if err != nil {
		log.Println(err)
		return
	}

	commands := []string{
		"mkdir reyhan",
		"cd reyhan",
		"echo '#from Reyhan' >> readme.md",
		"exit",
	}

	var stdout, stderr bytes.Buffer
	sess.Stdout = &stdout
	sess.Stderr = &stderr

	for _, cmd := range commands {
		_, err := fmt.Fprintln(stdin, cmd)
		if err != nil {
			fmt.Println("error when execute command", cmd, "with error", err.Error())
			return
		}
	}

	err = sess.Wait()
	if err != nil {
		log.Println(err)
		return
	}

	outputErr := stderr.String()
	fmt.Println("============== ERROR")
	fmt.Println(strings.TrimSpace(outputErr))

	outputString := stdout.String()
	fmt.Println("============== OUTPUT")
	fmt.Println(strings.TrimSpace(outputString))
}
