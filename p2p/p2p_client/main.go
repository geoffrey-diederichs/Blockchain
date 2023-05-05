package main

import (
	"net"
	"log"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp", "139.59.129.241:45454")
	if err != nil {
		log.Println("error connecting to peer:", err)
		return
	}

	writer := bufio.NewWriter(conn)

	message := "Est ce que c'est bon pour vous ?"
	_, err = writer.WriteString(message)
	if err != nil {
		log.Println("error writing message to connection:", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		log.Println("error flushing writer:", err)
		return
	}
}