package main

import "fmt"

var counter int = 0

func main() {
	c := make(chan int)
	done := make(chan bool)
	counter = 3

	go chunkWriter(c, done)
	go chunkWriter(c, done)
	go chunkWriter(c, done)

	go finishWriting(c, done)

	for n := range c {
		fmt.Println(n)
	}
}

func chunkWriter(c chan<- int, done chan<- bool) {
	for i:= 0; i < 10; i ++ {
		c <- i
	}

	done <- true
}

func finishWriting(c chan int, done chan bool) {

	for i := 0; i < counter; i++ {
		<-done
	}
	
	close(c)
}
