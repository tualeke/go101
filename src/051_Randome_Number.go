package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(100), ",",rand.Intn(100))

	fmt.Println(rand.Float64())

	fmt.Println(rand.Float64()* 5 + 5, ",", rand.Float64() * 5 + 5)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100),",",r1.Intn(100))

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	fmt.Println(r2.Intn(100), ",", r2.Intn(100))

}
