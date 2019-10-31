package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Action func(int) int
}

func Get(){
	fmt.Println("get")
}

func main() {
	p := Person{Name:"james", Age: 13, Action: func(i int) int{
		return 3
	}}
	action := p.Action(3)
	fmt.Println("action = ", action)
}


