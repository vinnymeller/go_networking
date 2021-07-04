package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    // deliver files from the directory /var/www
    //fileServer := http.FileServer(http.Dir("/home/httpd/html/"))
    fileServer := http.FileServer(http.Dir("/home/httpd/html"))

    // register the handler and deliver requests to it
    err := http.ListenAndServe(":8000", fileServer)
    checkError(err)
    // That's it!
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
