package main

import (
	"fmt"
	_ "fmt"
	"math/rand"
	_ "math/rand"
	"time"
	_ "time"
)

func calculateValue(c chan int)  {
	value:= rand.Intn(10)
	fmt.Println("Calculated Random Value: {} :", value)
	c <- value
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

func compute (value int)  {
	for i:=0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main()  {
	fmt.Println("Goroutine Tutorial")
	//go keyword creates a coroutine
	//go compute(10)
	//go compute(10)
	//var input string
	//fmt.Scanln(&input)

	//creates new channel
	channel := make(chan int)
	//waits to close the channel until the end of the main function call
	defer close(channel)

	go calculateValue(channel)
	go calculateValue(channel)
	//we calculate a single random value between 1-10, print
	//this out and then send
	//this value to our values channel by calling values <- value.
	value := <- channel
	fmt.Println(value)
}