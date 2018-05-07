package main

import "fmt"

//Struct
type showFunc struct {
	Mr   string
	name string
}
type nomoney struct {
	showFunc
	Nomoney bool
}
type ExFriend interface {
	proof()
}

func say(s ExFriend) {
	s.proof()
}

//Function
func (sh showFunc) hello() {
	fmt.Println(sh.Mr, sh.name, "say Hellos!!")
}
func (nm nomoney) proof() {
	fmt.Println(nm.Mr, nm.name, "He's doesn't has money that", nm.Nomoney)
}

func main() {
	show := showFunc{
		"Mr.",
		"Sirawich",
	}
	no_mon := nomoney{
		showFunc{
			"Mr",
			"Sirawich",
		},
		true,
	}
	//normal
	show.hello()
	no_mon.proof()
	fmt.Println("=================")
	//polymoprism
	say(no_mon)
}
