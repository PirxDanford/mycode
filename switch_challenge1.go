package main

import (
    "fmt"
    "runtime"
		"strings"
)

func main() {

    gover := runtime.Version()
    switch {
    case strings.HasPrefix(gover,"go1.18"):
        fmt.Println("You are using the latest version of GoLang")
    case strings.HasPrefix(gover,"go1.17"):
        fmt.Println("This version of Go is fine")
    case strings.HasPrefix(gover,"go1.16") || strings.HasPrefix(gover,"go1.15") :
        fmt.Println("You are using an older, but acceptable version of GoLang")
    default:                       // in all other cases run the code below
        fmt.Println("I cannot make a recommendation.")
    }
}

