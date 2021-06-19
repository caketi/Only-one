package main

import (
	 
	"time"
	"fmt")


func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(1 * time.Millisecond)

		}
	}()

	return c 
}

func main(){
	c := boring("boring!!!!")
	for i:=0; i < 5; i++{
		fmt.Printf("U say: %q\n", <-c)
	}
	fmt.Println("YOU!!!!")
}