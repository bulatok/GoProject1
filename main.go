package main

import (
	"fmt"
	//"log"
	//"sort"
	"os"
)

func IsDoteGit(s string) bool {
	return (s == ".git" || s == "pkg")
}

const way = "../../"

func recurs(entry []os.DirEntry, sway string, tabs int) {
	for _, v := range entry{
		if IsDoteGit(v.Name()){
			continue
		} // checking that file is not .git

		for i := 0; i < tabs; i++{
			for j:=0; j < 4; j++{
				fmt.Printf(" ")
			}
		}
		fmt.Printf("%s\n", v.Name())

		if v.IsDir() == true{
			//fmt.Printf("/")
			nxt, _ := os.ReadDir(sway + "/" + v.Name())
			tabs += 1
			recurs(nxt, sway + "/" + v.Name(), tabs)
			tabs -= 1
		}
	}
}


func main() {
	//g, _ := os.ReadDir("../../pkg")
	//for _, v := range(g){
	//	fmt.Println(v.Name())
	//}

	start, _ := os.ReadDir(way)
	recurs(start, way, 0)


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
