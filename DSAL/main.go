package main

import "fmt"

type pet interface {
}

type cat struct{
	name string
}

type dog struct{
	name string
}

func filterCatFromPet(petList []pet) []cat{
	catList := []cat{}
	for _, p := range petList{
		switch v := p.(type){
		case cat:
			catList = append(catList, v)

		}
	}
	return catList
}

func main(){
	d1 := dog{"dog 1"}
	d2 := dog{"dog 2"}
	c1 := cat{"cat 1"}
	c2:= cat{"cat 2"}
	petList := []pet{d1, d2, c1, c2}
	catList := filterCatFromPet(petList)
	fmt.Printf("pet list %v\n", petList)
	fmt.Printf("cat list %v\n", catList)
	
}
