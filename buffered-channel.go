//unbuffered channel 不論是「從 channel 讀值」（需等到值被其他 goroutine 寫入），
//或「把值寫入 channel」（需等到值被其他 goroutine 讀出）都會阻塞當下的 goroutine
package main

import (
	"fmt"
)

/*當 buffer size 不是 0 的話，就屬於 buffered channel
    - 「從 channel 讀值」時，只有在 buffered 是空的時才會 blocking
    -「把值寫入 channel」時，該 goroutine 並不會被阻塞，除非該 buffer 已經填滿（full）且溢出（overflow）。當 buffer 已經填滿（full）時，再把新的一筆資料傳入 channel 時會造成溢出（overflow），此時 goroutine 才會被阻塞。
    - 讀值的動作一旦開始，就會一直到 buffer 變成 empty 為止才會結束。也就是說，讀取 channel 的那個 goroutine 需到等到 buffer 完全清空後才會阻塞。
*/

func main() {

	c := make(chan bool, 1)

	go func() {
		fmt.Println("IN ANOTHER GOROUTINE")

		fmt.Printf("Receive value from channel %v\n", <-c)
	}()

	c <- true
	
	fmt.Println("After Receive")
}