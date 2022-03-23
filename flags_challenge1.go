package main

import (
  "fmt"
	"os"
	"flag"
	"bufio"
)

func main(){
	f := flag.String("f", "lorem.txt", "a file")
  l := flag.Int("l", 1, "amount of lines to display")
	flag.Parse()
	fh, err := os.Open(*f)
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
  }
  defer fh.Close()
	fb := bufio.NewScanner(fh)
	for line := 0; line < *l; line++ {
		if !fb.Scan() {
			break
		}
		fmt.Println(fb.Text())
	}
}
