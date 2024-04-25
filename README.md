# FTP Protocol Checker

This project provides a command-line tool written in Go that checks an FTP server's support for plaintext FTP, FTPS (FTP Secure over TLS), and optionally SFTP (SSH File Transfer Protocol). It first checks for plaintext FTP and FTPS, and upon user request, it can also check for SFTP support.

## Prerequisites

Before you run the tool, make sure you have the following installed:
- Go (version 1.13 or newer)
- Dependencies:
  - `golang.org/x/crypto/ssh`
  - `github.com/pkg/sftp`

You can install the necessary dependencies using Go's package manager with the following commands:

```bash
go get -u golang.org/x/crypto/ssh
go get -u github.com/pkg/sftp
```


## Compilation

To compile the project, navigate to the directory containing the source code and run:

```
go build -o ftpchecker main.go
```

This command will compile the main.go file into an executable named ftpchecker.

## Usage

After compiling the tool, you can run it directly from the command line. Here is how to use the executable:

```
./ftpchecker <IP>
```

## Example

```
./ftpchecker 192.168.1.100
```

## Output

The program outputs the results directly to the terminal. Here's what you might see:

```
FTP Server response: 220 Welcome to FTP service.
Server allows plaintext FTP
Server supports FTPS
Do you want to check if the server supports SFTP? (y/n): y
Server supports SFTP
```

## Contributing

Contributions to this project are welcome! Please fork the repository, make your changes, and submit a pull request.
