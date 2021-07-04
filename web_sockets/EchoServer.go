package main

import (
    "fmt"
    "net/http"
    "os"
    "golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
    fmt.Println("Echoing")
    for n := 0; n < 10; n++ {
        msg := fmt.Sprintf("Hello %d", n)
        err := websocket.Message.Send(ws, msg)
        if err != nil {
            fmt.Println("Can't send")
            break
        }

        var reply string
        err = websocket.Message.Receive(ws, &reply)
        if err != nil {
            fmt.Println("Can't receive")
            break
        }
        fmt.Println("Received back from client: " + reply)
    }
}

func main() {
    http.Handle("/", websocket.Handler(Echo))
    err := http.ListenAndServe(":12345", nil)
    checkError(err)
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}
