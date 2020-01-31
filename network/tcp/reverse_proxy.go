package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/project-1/config"
)

var connectionCount int

// BUFFLENGTH is the length for buffers used in this 1024 is one kilobyte
const BUFFLENGTH int = 1024

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}

	newConn := make(chan string)

	for {
		go makeConnection(lis, newConn)
		<-newConn
	}
}

func makeConnection(lis net.Listener, newConn chan string) {

	clientConn, err := lis.Accept()
	connectionCount++
	fmt.Println(connectionCount)
	newConn <- "new connection made"

	addr := "localhost:" + strconv.FormatInt(config.PORT, 10)
	backEndConn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	input := make(chan string)
	output := make(chan string)

	go listen(clientConn, input, "client")
	go write(backEndConn, input, "back end")
	go listen(backEndConn, output, "back end")
	write(clientConn, output, "client")
	fmt.Println("Close connections")
	backEndConn.Close()
	clientConn.Close()
}

func listen(conn net.Conn, ch chan string, connName string) {
	for {
		buff := make([]byte, BUFFLENGTH)
		conn.SetReadDeadline(time.Now().Add(600 * time.Second))
		_, err := conn.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("listening to " + connName)
		fmt.Println(string(buff))
		ch <- string(buff)
	}
}

func write(conn net.Conn, ch chan string, connName string) {
	for {
		buff := <-ch
		conn.Write([]byte(buff))
		fmt.Println("writing to " + connName)
		fmt.Println(buff)
	}
}
