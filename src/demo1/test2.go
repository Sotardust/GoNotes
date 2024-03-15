package demo1

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func Main1() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, s string, duration time.Duration) {

	fmt.Fprintf(c, "\t", strings.ToUpper(s))
	fmt.Println("strings", strings.ToUpper(s))
	time.Sleep(duration)

	fmt.Fprintf(c, "\t", strings.ToLower(s))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	for input.Scan() {

		fmt.Println("", input.Text())
		echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
}
