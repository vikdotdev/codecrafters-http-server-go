package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

const protocolVersion = "HTTP/1.1"
const statusOk = "200 OK"
const statusNotFound = "404 Not Found"

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	defer conn.Close()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

    buf := make([]byte, 0, 1024)
    tmp := make([]byte, 256)

	bytes, err := conn.Read(tmp)
	if err != nil {
		fmt.Println("Error reading the request: ", err.Error())
		os.Exit(1)
	}

	buf = append(buf, tmp[:bytes]...)

	fmt.Print("Request line: \n", string(buf))
	headers, _, found := strings.Cut(string(buf), "\r\n\r\n")
	if !found {
		fmt.Println("Error parsing the request: ", err.Error())
		os.Exit(1)
	}
	requestLine, _, found := strings.Cut(string(headers), "\r\n")
	if !found {
		fmt.Println("Error parsing the request line: ", err.Error())
		os.Exit(1)
	}

	method, rest, found := strings.Cut(requestLine, " ")
	if !found {
		fmt.Println("Error parsing the method: ", err.Error())
		os.Exit(1)
	}
	path, rest, found := strings.Cut(rest, " ")

	fmt.Println("METHOD:", method)
	fmt.Println("PATH:", path)

	if path == "/" {
		conn.Write([]byte(fmt.Sprintf("%s %s\r\n\r\n", protocolVersion, statusOk)))
	} else {
		conn.Write([]byte(fmt.Sprintf("%s %s\r\n\r\n", protocolVersion, statusNotFound)))
	}

}
