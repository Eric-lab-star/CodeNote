// composite pattern
// folders and files
package main

import "fmt"

type searcher interface{
	search()
}      

type file struct{
	name string
}

func (f file) search(){
	fmt.Println(f.name)	
}

type folder struct{
	searchers []searcher
	name string
}

func (f folder) search(){
	fmt.Println(f.name)
	for _, searcher := range f.searchers{
		searcher.search()
	}
}

func (f *folder) add( searcher ...searcher){
	f.searchers = append(f.searchers, searcher...)
}

func main(){
	root := folder{name: "root"}	
	file1 := file{name: "file1"}
	file2 := file{name: "file2"}
	file3 := file{name: "file3"}
	file4 := file{name: "file4"}
	dir1 := folder{name: "dir1"}
	dir2 := folder{name: "dir2"}
	dir1.add(file3, file4)
	root.add(file1, file2, dir1, dir2)	
	root.search()
}
