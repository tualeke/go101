package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	var next = intSeq()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	var new = intSeq()
	fmt.Println(new())

}



