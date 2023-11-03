package main

import "fmt"

type Set[T comparable] struct {
	t_map map[T]bool
}

func (set *Set[T]) New() {
	set.t_map = make(map[T]bool)
}

func (set *Set[T]) ContainsElement(element T) bool {
	_, exists := set.t_map[element]
	return exists
}

func (set *Set[T]) DeleteElement(element T) {
	delete(set.t_map, element)
}

func (set *Set[T]) AddElement(element T) {
	if !set.ContainsElement(element) {
		set.t_map[element] = true
	}
}

func main() {
	set := Set[int]{}
	set.New()
	element := 2
	fmt.Println("does the set have the element: ", set.ContainsElement(element))
	set.AddElement(element)
	fmt.Println("initial set: ", &set)
}
