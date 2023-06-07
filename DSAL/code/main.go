// Adapter pattern
package main

import (
	"fmt"

)

func main(){
	mac := mac{}		
	window := window{}
	windowsAdapter := windowsAdapter{
		window,
	}
	client := client{}
	client.lighteningPortOnlyComp(mac)
	client.lighteningPortOnlyComp(windowsAdapter)
}

type client struct{}

func (c client) lighteningPortOnlyComp(comp comp){
	fmt.Println("conntecting lightening port...\n. \n. \n.")
	comp.insertIntoLighting()
}


type comp interface {
	insertIntoLighting()
}

type mac struct{}

func (mac mac) insertIntoLighting(){
	fmt.Println("Mac is ready to use")
}

type window struct{}

func (win window) insertIntoUSB(){
	fmt.Println("Window is ready to use")
}


type windowsAdapter struct{
	windwo window
}

func (windowAdapter windowsAdapter) insertIntoLighting(){
	fmt.Println("processing window Adapter")
	windowAdapter.windwo.insertIntoUSB()
}
