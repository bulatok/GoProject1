package main

import (
	"fmt"
	//"log"
	//"sort"
	"os"
)

func IgnoreDirs(s string) bool {
	return (s == ".git" || s == "pkg")
}

const way = "../../"
func Govno(tabs int){
	for i := 0; i < tabs; i++{
		for j:=0; j < 4; j++{
			fmt.Printf(" ")
		}
	}
}
func recurs(entry []os.DirEntry, sway string, tabs int) {
	for _, v := range entry{
		if IgnoreDirs(v.Name()){
			continue
		}

		Govno(tabs)

		fmt.Printf("%s\n", v.Name())

		if v.IsDir() == true{
			nxt, _ := os.ReadDir(sway + "/" + v.Name())
			tabs += 1
			recurs(nxt, sway + "/" + v.Name(), tabs)
			tabs -= 1
		}
	}
}


func main() {

	start, _ := os.ReadDir(way)
	recurs(start, way, 1)


	//d, err := os.ReadDir(way)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//for _, f := range d{
	//	nxt, _ := os.ReadDir(way + f.Name())
	//	if f.Name() == ".git"{
	//		continue
	//	}
	//	fmt.Printf("%s\n", f.Name())
	//	for _, n := range nxt{
	//		fmt.Printf("\t%s\n", n.Name())
	//
	//	}
	//}
}
