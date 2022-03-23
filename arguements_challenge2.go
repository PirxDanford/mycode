package main

import (
	"fmt"
	"os"
)

func main(){
	for iArg := 1; iArg < len(os.Args); iArg++ {
		fmt.Println("Argument",iArg,"is",os.Args[iArg])
	}
}
