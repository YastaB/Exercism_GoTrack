package main

import (
	"fmt"
)

func test(){

	ch := make(chan string)
	go func() {
		ch <- "a post"
		ch <- "a post"
		ch <- "a post"
		ch <- "a post"
		fmt.Println("Sent")
	}()

	fmt.Printf("Received: %s\n", <-ch)
	fmt.Printf("Here %s\n", <- ch )
	fmt.Printf("Here %s\n", <- ch )
	fmt.Printf("Here %s\n", <- ch )
	fmt.Println("Done")
}