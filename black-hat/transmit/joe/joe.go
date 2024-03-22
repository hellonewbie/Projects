package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Connect to the server on TCP port 20080.
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		log.Fatalln("Unable to connect to the server")
	}
	defer conn.Close()

	// Create a scanner to read input from the user.
	scanner := bufio.NewScanner(os.Stdin)

	// Create a goroutine to read and display server responses.
	go func() {
		for {
			// Read server response.
			response := make([]byte, 512)
			size, err := conn.Read(response)
			if err != nil {
				log.Fatalln("Error reading from server:", err)
				return
			}
			fmt.Printf("imformation: %s\n", string(response[:size]))
		}
	}()

	// Read user input and send it to the server.
	for {
		fmt.Print("Enter text to send to the server: \n")
		scanner.Scan()
		text := scanner.Text()

		// Send user input to the server.
		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Fatalln("Error writing to server:", err)
			return
		}
	}
}
