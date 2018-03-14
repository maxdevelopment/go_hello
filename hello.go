package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxMapSize = 10

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func logger(channel <-chan int) {
	data := make(map[int]int)
	for {
		msg := <-channel

		if len(data) > maxMapSize {
			for i := range data {
				delete(data, i)
			}
		}
		data[msg] = msg
		fmt.Println(data)
	}
}

func randData(channel chan<- int) {
	for {
		channel <- random(0, 10)
		time.Sleep(time.Second * 2)
	}
}

func randDataExt(channel chan<- int) {
	for {
		channel <- random(100, 200)
		time.Sleep(time.Second * 3)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	channel := make(chan int)
	go logger(channel)
	go randData(channel)
	go randDataExt(channel)

	var input string
	fmt.Scanln(&input)
}
