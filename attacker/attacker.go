package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const maxBufSize = 1024

func main() {

	// Wait for connection
	attackerAddress, err := net.ResolveUDPAddr("udp", "127.0.0.1:4444")
	if err != nil {
		log.Panic(err)
	}

	conn, err := net.ListenUDP("udp", attackerAddress)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	// Create output file
	file, err := os.OpenFile("capture.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to create logs.txt: ", err)
		return
	}
	defer file.Close()

	for {
		buf := make([]byte, maxBufSize)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Panic(err)
		}

		if n > 0 {
			fmt.Printf("Captured: %s\n", string(buf))

			_, err = file.WriteString(string(buf))
			if err != nil {
				fmt.Println("Failed to write to file:", err)
			}
		}
	}

}
