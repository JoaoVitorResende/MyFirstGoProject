package examplesdefer

import "fmt"

func DoDefer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func DoDeferAnonimus(){
	x:= 10
	defer func(y int){
		fmt.Println(y)
	}(x)
	x = 50
	fmt.Println(x)
}