package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func scan(ip string, port int) {
	address := net.JoinHostPort(ip, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err == nil {
		fmt.Printf("PORT %d is OPEN\n", port)
		if err := conn.Close(); err != nil {
			log.Fatal("Error occurred when closing connection")
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 1; i <= 1024; i++ {
		wg.Go(func() {
			scan("127.0.0.1", i)
		})
	}

	wg.Wait()
}
