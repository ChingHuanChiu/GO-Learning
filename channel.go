package channel

import "fmt"


func square (arr []int, ch chan int){

	for _, data := range arr {

		ch <- data * data
	}
	close(ch)
}


func main() {

	arr := []int{1, 3, 5, 7, 9}

	chn := make(chan, int)

	go square(arr, chan)

	for square_val := range arr {
		fmt.Println(square_val)
	}
}