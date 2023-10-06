package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(22)
	intList.PushBack(33)
	for element := intList.Front(); element != nil ; element = element.Next() {
		fmt.Printf("Type: %T\n",element.Value)
	}
}
