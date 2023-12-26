package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"
)

func dial(host string) bool {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Println("error connecting:", err)

		return true
	}

	defer conn.Close()

	log.Println("Connected to host!")

	for {
		_, err = conn.Write([]byte("Hello, server!"))
		if err != nil {
			log.Println("error writing to connection:", err)

			break
		}

		ticker := time.NewTicker(1 * time.Second) // Change the interval as needed
		defer ticker.Stop()

		for range ticker.C {
			// send ping message
			_, er := conn.Write([]byte("Ping\n"))
			if er != nil {
				log.Println("Error writing to connection:", er)

				return true
			}

			log.Println("Ping sent to server")
		}
	}

	return false
}

func main() {
	var (
		hostFlag = flag.String("host", "localhost", "target host address")
	)

	flag.Parse()

	for i := 0; i < 1000; i++ {
		go func() {
			if dial(*hostFlag) {
				os.Exit(1)
			}
		}()
	}

	os.Exit(0)
}
