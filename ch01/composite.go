package main

import "fmt"

type IComposite interface {
	perform()
}
type Leaflet struct {
	name string
}

func (l *Leaflet) perform() {
	fmt.Println("Leaflet " + l.name)
}

type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)
	for _, l := range branch.leafs {
		l.perform()
	}
	for _, b := range branch.branches {
		b.perform()
	}
}

func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}
func (branch *Branch) getLeafs() []Leaflet {
	return branch.leafs
}
