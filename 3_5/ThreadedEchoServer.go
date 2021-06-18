package main

import (
    "net"
    "os"
    "fmt"
)

func main() {

    service := ":1201"
    tcpAddr, err := net.ResolveTCPAddr("ip4", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        go handleClient(conn)   // run as a goroutine. this is the beauty of go!
    }

}

func handleClient(conn net.Conn) {
    defer conn.Close()  // close the connection on exit TODO: lookup "defer" keyword

    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }
        _, err2 := conn.Write(buf[0:n])
        if err2 != nil {
            return
        }
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
