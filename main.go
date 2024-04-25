package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func checkPlainFTP(ip string) {
	// Connect to the FTP server
	conn, err := net.DialTimeout("tcp", ip+":21", 5*time.Second)
	if err != nil {
		fmt.Printf("Error connecting to FTP server: %v\n", err)
		return
	}
	defer conn.Close()

	// Read the server's greeting message
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Println("FTP Server response:", scanner.Text())
	}

	// Check if the server's response indicates that it's ready for a new user
	if strings.Contains(scanner.Text(), "220") {
		fmt.Println("Server allows plaintext FTP")
	} else {
		fmt.Println("Server did not respond with a welcome message, it may not allow plaintext FTP")
	}
}

func checkFTPS(ip string) {
	// Try to connect to the FTP port
	conn, err := net.DialTimeout("tcp", ip+":21", 5*time.Second)
	if err != nil {
		fmt.Printf("Error connecting to FTP server: %v\n", err)
		return
	}
	defer conn.Close()

	// Read the server's greeting message
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading from FTP server: %v\n", err)
		return
	}

	// Check if the server mentions support for TLS (FTPS)
	if strings.Contains(string(buffer), "TLS") {
		fmt.Println("Server supports FTPS")
	} else {
		fmt.Println("Server does not support FTPS")
	}
}

func checkSFTP(ip string) {
	// SSH client config with a non-existent user and no authentication.
	sshConfig := &ssh.ClientConfig{
		User:            "nonexistentuser",
		Auth:            []ssh.AuthMethod{},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // For simplicity, ignore host key verification
		Timeout:         5 * time.Second,
	}

	// Try to establish an SSH connection
	sshConn, err := ssh.Dial("tcp", ip+":22", sshConfig)
	if err != nil {
		fmt.Printf("Error connecting to SSH server: %v\n", err)
		return
	}
	defer sshConn.Close()

	// Create an SFTP client
	sftpClient, err := sftp.NewClient(sshConn)
	if err != nil {
		fmt.Println("Server does not support SFTP")
	} else {
		fmt.Println("Server supports SFTP")
		sftpClient.Close()
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <IP>")
		os.Exit(1)
	}

	ip := os.Args[1]

	// Check what the server supports
	checkFTPS(ip)
	checkPlainFTP(ip)

	// Ask the user if they want to check for SFTP
	fmt.Print("Do you want to check if the server supports SFTP? (y/n): \n this feature still requires testing")
	var response string
	fmt.Scanln(&response)

	if strings.ToLower(response) == "y" {
		checkSFTP(ip)
	}
}
