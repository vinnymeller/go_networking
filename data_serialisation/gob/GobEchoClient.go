package main

import (
    "fmt"
    "net"
    "os"
    "encoding/gob"
)

type Person struct {
    Name Name
    Email []Email
}

type Name struct {
    Family string
    Personal string
}

type Email struct {
    Kind string
    Address string
}

func (p Person) String() string {
    s := p.Name.Personal + " " + p.Name.Family
    for _, v := range p.Email {
        s += "\n" + v.Kind + ": " + v.Address
    }
    return s
}

func main() {
    person := Person{
        Name: Name{Family: "Meller", Personal: "Vinny"},
        Email: []Email{Email{Kind: "home", Address: "vinnymeller@gmail.com"},
                       Email{Kind: "school", Address: "vmeller@umich.edu"}}}

    if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "host:port")
        os.Exit(1)
    }

    service := os.Args[1]
    conn, err := net.Dial("tcp", service)
    checkError(err)

    encoder := gob.NewEncoder(conn)
    decoder := gob.NewDecoder(conn)

    for n := 0; n < 10; n++ {
        encoder.Encode(person)
        var newPerson Person
        decoder.Decode(&newPerson)
        fmt.Println(newPerson.String())
    }
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
