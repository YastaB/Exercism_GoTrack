package main

import "fmt"

type Record struct {
	ID int
	Parent int
}

func main(){
	record := Record{ID:0}
	fmt.Println(record)

}