package main

import (
	"fmt"
	"github.com/okieraised/go2com"
	"log"
	"os"
)

func main() {
	go2com.InitTagDict()
	file, err := os.Open("./test_data/01.dcm")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := info.Size()

	parser, err := go2com.NewParser(file, fileSize, true, false)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = parser.Parse()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, elem := range parser.GetDataset().Elements {
		fmt.Println(elem)
	}
}
