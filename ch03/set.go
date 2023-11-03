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

func (set *Set[T]) InterSect(anotherSet *Set[T]) *Set[T] {
	intersectSet := &Set[T]{}
	intersectSet.New()
	for value, _ := range set.t_map {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func (set *Set[T]) Union(anotherSet *Set[T]) *Set[T] {
	unionSet := set
	for value, _ := range anotherSet.t_map {
		unionSet.AddElement(value)
	}
	return unionSet
}

func main() {
	set := &Set[int]{}
	set.New()
	element := 2
	fmt.Println("does the set have the element: ", set.ContainsElement(element))
	set.AddElement(element)
	set.AddElement(3)
	fmt.Println("initial set: ", set)
	anotherSet := &Set[int]{}
	anotherSet.New()
	anotherSet.AddElement(2)
	fmt.Println(set.InterSect(anotherSet))
}
