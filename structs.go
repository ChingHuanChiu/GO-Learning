package main

import (
		"fmt"
)


/*
 There are no class in GO, using the struct instead.
  - able to  use the pointer to refer the struct address
  - new(struct_name) is same as &struct_name, but can not initial the value in struct
  - Method:
	the method do not define in the struct, it assigns the methods to a struct through 'Receiver' instead.
*/


type BankAccount struct {

	name string
	deposit float64
}


func (ba *BankAccount) Lend(baBorrower *BankAccount, amount float64) {

	if ba.deposit -= amount; ba.deposit < 0{
		fmt.Println("You do not have enough money")
	}else{
		ba.deposit -= amount
		baBorrower.deposit += amount

	}
	
	
}


func main() {

	steven := &BankAccount{"Steven", 100000000.0}

	fmt.Println("Steven Original Account", *steven)

	boss := &BankAccount{"Jobs", 1000.0}

	fmt.Println("Boss Original Account", *boss)

	steven.Lend(boss, 150)

	fmt.Println("Steven remain Account", *steven)
	fmt.Println("Boss remain Account", *boss)


}

