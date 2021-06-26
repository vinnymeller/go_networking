package main

import (
    "fmt"
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

func main() {
    person := Person{
        Name: Name{Family: "Meller", Personal: "Vinny"},
        Email: []Email{Email{Kind: "home", Address: "vinnymeller@gmail.com"},
                 Email{Kind: "school", Address: "vmeller@umich.edu"}}}

    saveGob("person.gob", person)
}

func saveGob(fileName string, key interface{}) {
    outFile, err := os.Create(fileName)
    checkError(err)

    encoder := gob.NewEncoder(outFile)
    err = encoder.Encode(key)
    checkError(err)

    outFile.Close()
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
