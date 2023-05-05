package main

import (
    "log"
    "bufio"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", "139.59.129.241:45454")
    if err != nil {
        log.Fatal("error starting listener:", err)
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)

    for {
        message, err := reader.ReadString('\n')
        log.Println(message)

        response := "ack\n"
        _, err = writer.WriteString(response)
        if err != nil {
            log.Println("error sending response:", err)
            return
        }
        writer.Flush()
    }
}