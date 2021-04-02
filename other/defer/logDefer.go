// Example of using defer keyword with writeing in log file 
package main

import (
	"fmt"
	"log"
	"os"
)

var LOG_FILE = "/tmp/myGo.log"

func logBlock(aLog *log.Logger, function string)  {
	aLog.Println("---- FUNCTION ", function, " ---- ")
	defer aLog.Println("---- FUNCTION ", function, " ---- ")

	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

func main()  {
	f, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "logDefer ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")

	logBlock(iLog, "first")
	logBlock(iLog, "second")
}