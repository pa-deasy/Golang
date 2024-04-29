package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readFiles(".")
	readFiles("/wrong")
}

func readFiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
