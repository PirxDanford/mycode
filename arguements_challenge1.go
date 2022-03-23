package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Arguments:",(len(os.Args)-1))
}
