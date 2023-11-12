package main

import (
	"fmt"
	"time"
)

func Ping1S(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Ping : ", i)
		time.Sleep(1 * time.Second)
	}
	c <- 10
}

func SendNoti5S(c chan string)  {
	fmt.Println("Send noti")
	time.Sleep(5 * time.Second)
	fmt.Println("Send noti")
	c <- "Success"
}

func main6() {
	fmt.Println("Go routine and channel")

	chInt := make(chan int)
	chStr := make(chan string)

	go Ping1S(chInt)
	go SendNoti5S(chStr)

	// time.Sleep(10 * time.Second)

	pingVal, notiVal := <- chInt, <- chStr
	fmt.Println(pingVal, notiVal)
	fmt.Println("Finished")
}