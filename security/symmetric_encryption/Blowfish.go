package main

import (
    "bytes"
    "golang.org/x/crypto/blowfish"
    "fmt"
    "os"
)

func main() {
    key := []byte("my key")
    cipher, err := blowfish.NewCipher(key)
    checkError(err)

    src := []byte("hello\n\n\n")
    var enc [512]byte

    cipher.Encrypt(enc[0:], src)
    var decrypt [8]byte
    cipher.Decrypt(decrypt[0:], enc[0:])

    result := bytes.NewBuffer(nil)
    result.Write(decrypt[0:8])
    fmt.Println(string(result.Bytes()))
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
