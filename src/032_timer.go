package main

import (
	"time"
	"fmt"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<- timer1.C
	fmt.Println("timer expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expeired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

}