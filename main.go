package main

import (
	"flag"
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
	ip := flag.String("ip", "127.0.0.1", "IP address to scan")
	startPort := flag.Int("startPort", 1, "Start port")
	endPort := flag.Int("endPort", 1024, "End port")
	flag.Parse()

	wg := &sync.WaitGroup{}

	for i := *startPort; i <= *endPort; i++ {
		wg.Go(func() {
			scan(*ip, i)
		})
	}

	wg.Wait()
}
