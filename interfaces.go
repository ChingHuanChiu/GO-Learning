package interfaces

import "fmt"

//Reference : https://haren.medium.com/notes-go-golang-%E5%9F%BA%E6%9C%AC%E8%AA%9E%E6%B3%95-part2-f92a891c078f

type IAccount  interface {

	Deposit(amount int)
}

type BankAAccount struct {

	Fee int
}

func (banka *BankAAccount) Deposit(amount int) {
	
	fee := banka.Fee
	fmt.Printf("%v is add to your bank_A, and fee is %v \n", amount - fee, fee)
}



type BankBccount struct {

	Fee int
}

func (bankb *BankBccount) Deposit(amount int){

	fee := bankb.Fee

	fmt.Printf("%v is add to your bank_B, and fee is %v \n", amount - fee, fee)


}


func main() {

	var account IAccount

	account = &BankAAccount{Fee: 100}
	account.Deposit(1000)

	account = &BankBccount{Fee: 10}
	account.Deposit(1000)

}



