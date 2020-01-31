package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/project-1/config"
)

// const BUFFLENGTH int = 1024

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}

	clientConn, err := lis.Accept()
	defer clientConn.Close()

	addr := "localhost:" + strconv.FormatInt(config.PORT, 10)
	backEndConn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer backEndConn.Close()

	input := make(chan string)
	output := make(chan string)

	// for {
	// 	buff := make([]byte, 1024)
	// 	clientConn.Read(buff)
	// 	fmt.Println(string(buff))
	// 	backEndConn.Write(buff)
	// 	respBuff := make([]byte, 1024)
	// 	backEndConn.Read(respBuff)
	// 	fmt.Println(string(respBuff))
	// 	clientConn.Write(respBuff)
	// }
	// output := make([]byte, 0)

	go listen(clientConn, input)
	go write(backEndConn, input)
	go listen(backEndConn, output)
	write(clientConn, output)
}

func listen(conn net.Conn, ch chan string) {
	for {
		buff := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		_, err := conn.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		ch <- string(buff)
	}
}

func write(conn net.Conn, ch chan string) {
	for {
		buff := <-ch
		conn.Write([]byte(buff))
		fmt.Println(buff)
	}
}
