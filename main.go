package main

import (
	"fmt"
	//"log"
	"sort"
	"os"
)

func IgnoreDirs(s string) bool {
	return (s == ".git" || s == "pkg")
}

const way = "../../"
func Govno(tabs int, last bool){
	for i := 0; i < tabs-1; i++{
		for j:=0; j < 4; j++{
			fmt.Printf(" ")
		}
	}
	if last == false {
		fmt.Printf("├")
	} else{
		fmt.Printf("└")
	}
		for j:=0; j < 3; j++ {
		fmt.Printf("─")
	}
}
func recurs(entry []os.DirEntry, sway string, tabs int) {
	// чтобы вывести всё было отсортированным dude
	sort.Slice(entry, func(i, j int) bool{
		return entry[i].Name() < entry[j].Name()
	})
	for i, v := range entry{
		if IgnoreDirs(v.Name()){
			continue
		}

		//fmt.Printf("%d %d", i, len(entry))
		Govno(tabs, (i == len(entry) - 1))

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
