package main

import "fmt"

type Set struct{
	integerMap map[int]bool
}

func (set *Set) New(){
	set.integerMap = make(map[int]bool)
}

func (set *Set) ContainElement(ele int) bool{
	_, exists := set.integerMap[ele]
	return exists
}

func (set *Set) AddElement(ele int){
	if !set.ContainElement(ele){
		set.integerMap[ele] = true
	}
}

func (set *Set) deleteElement(ele int){
	delete(set.integerMap, ele)

}

func (set *Set) InterSect(anotherSet *Set) *Set{
	intersectSet := new(Set)
	intersectSet.New()
	for v := range set.integerMap{
		if anotherSet.ContainElement(v) {
			intersectSet.AddElement(v)
		}	
	}
	return intersectSet

}

func (set *Set) Union (anotherSet *Set) *Set{
	

}
func main(){
	set1:= new(Set)
	set1.New()
	set1.AddElement(2)
	set1.AddElement(4)
	set1.AddElement(6)
	set1.AddElement(8)

	set2 := new(Set)
	set2.New()
	set2.AddElement(3)
	set2.AddElement(6)
	set2.AddElement(9)
	set2.AddElement(12)

	intersect:=	set1.InterSect(set2)
	fmt.Println(intersect)
	
	
}
