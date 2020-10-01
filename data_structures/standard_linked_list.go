package main

import (
	"container/list"
	"errors"
	"fmt"
)

func main(){
	linkedList := list.New()

	type person struct {
		name string
		age int
	}
	type personWithProfession struct{
		name string
		age int
		profession string
	}
	linkedList.PushBack(person{"Berkay", 28})
	linkedList.PushBack(person{"Meltem", 28})
	linkedList.PushBack(person{"Beril", 24})
	linkedList.PushBack(personWithProfession{"Beril", 24, "Chemist"})
	current := linkedList.Front()
	for current != nil{
		val, ok := current.Value.(person)
		if !ok{
			panic(errors.New(fmt.Sprintf("Cannot convert: %+v\n", current.Value)))
		}
		fmt.Printf("Person Name: %s Age: %d\n", val.name, val.age)
		current = current.Next()
	}




}

