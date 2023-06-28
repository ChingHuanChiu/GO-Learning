package goroutine

import (
	"fmt"
	"time"
)


func earnMoney(stock string){
	for year := 1 ; year < 10; year++ {

		Printf("I earn money from %s in %d year", stock, year)
		time.Sleep(time.Second)
	}
}


func main() {

	go earnMoney("T")
	go earnMoney("TSLA")
	go earnMoney("MO")
	go earnMoney("DOGE")
	go earnMoney("RUBY")

	for idx := 1; id < 5: id++{
		fmt.Printf("main routine:  %d", idx)
		time.Sleep(time.Second)
	}
	time.Sleep(3 * time.Second)
}