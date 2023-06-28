package main

import (
	"main"
)

func greet(roc <-chan string){

	fmt.Println("Hello " + <-roc + "!")
}

//希望在某一個 goroutine 中只能從 channel 讀取資料，但在 main goroutine 中可以對這個 channel 讀和寫資料時，
//可以透過 go 提供的語法來將 bi-directional channel 轉換成 unidirectional channel

func main() {

	fmt.Println("main() started")

	c := make(chan string)

	go greet(c)

	c <- "John"

	fmt.Println("main() Finished")
}