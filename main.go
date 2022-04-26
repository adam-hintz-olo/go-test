package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("hello world")
		time.Sleep(5 * time.Second)
	}
}
