package main

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "host:port")
        os.Exit(1)
    }
    url := os.Args[1]

    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
    client := &http.Client{}   
    resp, err := client.Do(request)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(2)
    }

    if resp.Status != "200 OK" {
        fmt.Println(resp.Status)
        os.Exit(2)
    }

    b, _ := httputil.DumpResponse(resp, false)
    fmt.Print(string(b))

    contentTypes := resp.Header["Content-Type"]
    if !acceptableCharset(contentTypes) {
        fmt.Println("Cannot handle", contentTypes)
        //os.Exit(4)
    }

    var buf [512]byte
    reader := resp.Body
    for {
        n, err := reader.Read(buf[0:])
        if err != nil {
            os.Exit(0)
        }
        fmt.Print(string(buf[0:n]))
    }
    os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
    // each type is like [text/html; charset=UTF-8]
    // we want the UTF-8 only
    for _, cType := range contentTypes {
        if strings.Index(strings.ToUpper(cType), "UTF-8") != -1 {
            return true
        }
    }
    return false
}
