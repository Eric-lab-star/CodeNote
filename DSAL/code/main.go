package main

import (
	"fmt"
)

// IProcess interface
type IProcess interface {
	process()
}

// Adaper struct
type Adaper struct {
	adaptee Adaptee
}

// Adapter class method process

func (adapter Adaper) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

// Adaptee Struct
type Adaptee struct {
	adapterType int
}

// Adaptee class method convert
func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func main() {
	var processor IProcess = Adaper{}
	processor.process()
}
