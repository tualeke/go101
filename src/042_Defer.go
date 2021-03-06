package main

import (
	"os"
	"fmt"
)

func main() {
	f := createFile("/tmp/defer.text")  // 1
	readFile(f)
	defer closeFile(f) // 3
	writeFile(f) // 2
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func readFile(f *os.File) {
	fmt.Println("reading")
	fmt.Fprintln(f,"read")
}
