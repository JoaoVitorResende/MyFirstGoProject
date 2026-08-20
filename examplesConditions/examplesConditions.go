package examplesconditions

import (
	"errors"
	"fmt"
	"math"
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