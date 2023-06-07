// Adapter pattern
package main

import (
	"fmt"
)

func main() {
	mac := mac{}
	window := window{}
	client := client{}
	client.Adapter(mac)
	client.Adapter(window)

}

type client struct{}
func (client client) Adapter(machine interface{}) {
	switch v := machine.(type) {
	case mac:
		v.insertIntoLighting()
	case window:
		v.insertIntoUSB()
	}
}


type mac struct{}

func (mac mac) insertIntoLighting() {
	fmt.Println("Mac is ready to use")
}

type window struct{}

func (win window) insertIntoUSB() {
	fmt.Println("Window is ready to use")
}
