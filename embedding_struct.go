package main

import "fmt"


type TeamMember struct {

	MemberName string
	Age int

}


func (tm *TeamMember) earn(amount int) {
	fmt.Printf("Earn the %v dollars \n", amount)

}

type Team struct {

	*TeamMember // name of column is no need
	TeamName string
}

func (t *Team) work(kind string){
	fmt.Printf("Earn money from %s \n", kind)
}


func main(){

	m1 := &TeamMember{"Steven", 28}


	team1 := &Team{m1, "Richest"}

	team1.work("Stock")
	team1.earn(100000000000) // same as team1.TeamMember.earn(100000000000)







	




}