import main,

import "fmt"

//Referenfce : https://haren.medium.com/notes-go-golang-%E5%9F%BA%E6%9C%AC%E8%AA%9E%E6%B3%95-part3-332c00b048aa

type Counter struct() {

	add func()
	minus func()
	print func()
}



func main() {

	counter := func counter {
		idx := 0
		func() {idx++}
		func() {idx--}
		func ()  { fmt.Println("idx=", idx)}
	}()

	counter.add()
	counter.print()
	counter.minus()
	counter.print()
}