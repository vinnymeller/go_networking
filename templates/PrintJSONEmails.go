package main

import (
    "html/template"
    "os"
    "fmt"
)

type Person struct {
    Name string
    Emails []string
}

const templ = `{"Name": "{{.Name}}",
"Emails": [
{{range $index, $elmt := .Emails}}
    {{if $index}}
        , "{{$elmt}}"
    {{else}}
        "{{$elmt}}"
    {{end}}
{{end}}
]
}
`

func main() {
    person := Person{
        Name: "vinny",
        Emails: []string{"vinnymeller@gmail.com", "vmeller@umich.edu"},
    }

    t := template.New("Person template")
    t, err := t.Parse(templ)
    checkError(err)

    err = t.Execute(os.Stdout, person)
    checkError(err)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
