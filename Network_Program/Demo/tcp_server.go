package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	address := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8000,
	}

	listener, err := net.ListenTCP("tcp4", &address)

	if err != nil {
		log.Fatal(err) //Println + os.Exit
	}

	for true {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("远程地址", conn.RemoteAddr())
		go echo(conn)
	}
}

func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second) //5s请求一次
	for now := range tick {
		n, err := conn.Write([]byte(now.String()))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}

		fmt.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}
