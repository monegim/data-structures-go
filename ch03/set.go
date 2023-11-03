package main

type Set[T comparable] struct {
	t_map map[T]bool
}

func (set *Set[T]) New() {
	set.t_map = make(map[T]bool)
}

func main() {

}
