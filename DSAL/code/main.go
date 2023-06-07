// Adapter pattern
package main

import (
	"fmt"
)

func main(){
	c := &Client{}
	window := &Windows{}
	adapter := WindowsAdapter{window}
	c.insertLighteningIntoComputer(&adapter)

}
type Client struct{
}

func (c *Client) insertLighteningIntoComputer(com computer){
	fmt.Println("client inserts connecter into computer")
}

type computer interface{
	InsertIntoLightningPort()
}

type Windows struct{}
 func (w *Windows) insertIntoUSBPort(){
	 fmt.Println("USB connector is plugged into windows machine")
 }

 type WindowsAdapter struct {
    windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
    fmt.Println("Adapter converts Lightning signal to USB.")
    w.windowMachine.insertIntoUSBPort()
}

