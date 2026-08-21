package examplesconditions

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func BasicIf() {
	if 1 < 2 {
		fmt.Print("opa")
	}
}

func DifferentIf(){
	if x:= math.Sqrt(4); x < 10{
		fmt.Println("10 e maior")
	}else if x > 0{
		fmt.Println("x e maior que zero")
	}else{
		fmt.Println("Caiu no else")
	}
}

func ReturnError () error{
	return errors.New("erro")
}

func SwitchStatement(x int){
	switch x{
	case 1:
		fmt.Println(1)
		//fallthrough caso queira que va para o de baixo
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(x)
	}
}

func IsWeekDay(x time.Time) bool{
	switch{
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default:
		return true
	}
}