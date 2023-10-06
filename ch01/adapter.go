package main

import "fmt"

type IProcess interface {
	process()
}
type Adapter struct {
	adaptee Adaptee
}
type Adaptee struct {
	adapterType int
}
func main() {
	
	fmt.Println("")
}