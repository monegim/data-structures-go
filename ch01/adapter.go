package main

import "fmt"

type IProcess interface {
	process()
}
type Adapter struct {
	adaptee Adaptee
}

func (a Adapter) process() {
	fmt.Println("Adapter process")
	a.adaptee.convert()
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}
func main() {

	var processor IProcess = Adapter{}
	processor.process()
}
