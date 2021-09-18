package main

import (
	"fmt"
	"os"
	"bytes"
	"sort"
	"strings"
)

var tabs = 1
const ParentPath = "C:/Users/kutlu/OneDrive/Рабочий стол/hw1_tree/"

func IgnoreDirs(s string) bool {
	baned := []string{
		".md",
		"dockerfile",
	}
	for _, v := range baned {
		if strings.HasSuffix(s, v) {
			return true
		}
	}
	return false
}

func Govno(tabs int, last bool, out *bytes.Buffer) {
	for i := 0; i < tabs-1; i++ {
		for j := 0; j < 4; j++ {
			out.Write([]byte(" "))
		}
	}
	if last == false {
		out.Write([]byte("├"))
	} else {
		out.Write([]byte("└"))
	}
	for j := 0; j < 3; j++ {
		out.Write([]byte("─"))
	}
}
func dirTree(out *bytes.Buffer, way string, printFiles bool) error {
	entry, _ := os.ReadDir(way)
	// чтобы вывести всё в отсортированным виде
	sort.Slice(entry, func(i, j int) bool {
		return entry[i].Name() < entry[j].Name()
	})


	for i, v := range entry {
		// игнорить определенные расширения файлов
		if IgnoreDirs(v.Name()) {
			continue
		}


		// вывод имени файла + его размер
		infa, _ := v.Info()
		if v.IsDir() == true {
			Govno(tabs, (i == len(entry)-1) || (i == 0 && len(entry) == 1), out)
			fmt.Fprintf(out, "%s \n", v.Name())
		} else if printFiles == true {
			Govno(tabs, (i == len(entry)-1) || (i == 0 && len(entry) == 1), out)
			if infa.Size() == 0 {
				fmt.Fprintf(out, "%s (empty)\n", v.Name())
			} else {
				fmt.Fprintf(out, "%s (%db)\n", v.Name(), infa.Size())
			}
		}


		if v.IsDir() == true {
			tabs += 1
			dirTree(out, way+"/"+v.Name(), printFiles)
			tabs -= 1
		}
	}
	return nil
}
func main() {
	var out bytes.Buffer
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := ParentPath + os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(&out, path, printFiles)
	fmt.Println(out.String())
	if err != nil{
		panic(err.Error())
	}
}