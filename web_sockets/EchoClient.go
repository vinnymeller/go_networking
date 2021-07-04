package main
// ws://localhost:12345 to connect
import (
    "golang.org/x/net/websocket"
    "fmt"
    "io"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "ws://host:port")
        os.Exit(1)
    }
    service := os.Args[1]

    conn, err := websocket.Dial(service, "", "http://localhost")
    checkError(err)

    var msg string
    for {
        err := websocket.Message.Receive(conn, &msg)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Couldn't receive msg " + err.Error())
            break
        }
        fmt.Println("Received from server: " + msg)
        // return the msg
        err = websocket.Message.Send(conn, msg)
        if err != nil {
            fmt.Println("Couldn't return msg")
            break
        }
    }
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}
