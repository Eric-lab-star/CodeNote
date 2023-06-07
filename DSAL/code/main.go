// Bridge Pattern
package main

import "fmt"

func main(){
	pc := mac{file:"file from mac"}
	p := &cannon{}
	pc.setPrinter(*p)
	pc.print()
	mobile := iphone{file:"file from mobile"}
	mobile.setPrinter(p)
	mobile.print()
}

// abstraction
type computer interface{
	print() // calls Printer.printFile
	setPrinter(Printer)
}
// refinced abstraction
type mac struct{
	file string
	printer Printer
}

func (mac mac) print(){
	mac.printer.printFile(mac.file)
}

func (mac *mac) setPrinter(printer Printer){
	mac.printer = printer
}

type iphone struct{
	file string
	printer Printer
} 

func (i iphone) print(){
	i.printer.printFile(i.file)
}

func (i *iphone) setPrinter(p Printer){
	i.printer = p
}
// Implementaion
type Printer interface{
	printFile(string)
}

// concrete implementation
type cannon struct{}

func (c cannon) printFile(file string){
	fmt.Println(file)	
}

