/*
 write a program that
	 - launches 10 goroutines
		 - each goroutine adds 10 numbers to a channel
	 - pull the numbers off the channel and print them
*/

package main 
import "fmt"

func main()  {
	 c := make(chan int)

	 for j:= 0; j < 10; j++ {
		 go func ()  {
			 for i := 0; i < 10; i++ {
				 c <- i
			 }
		 }()
	 }


	 for z:= 0; z < 100; z ++ {
		 fmt.Println(z, <-c)
	 }
}