package examplestypesandconstants

import (
	"fmt"
	"strconv"
)

func PositiveValue() byte {
	var value uint8 = 10
	//uint and byte is the same
	return value
}

func ChangeValueType() {
	value := 50
	changedValue := float64(value)
	fmt.Println(changedValue)
}

func ChangeValueToStringDictionary(){
	value := 65
	s := string(value)
	fmt.Println(s)
}

func ConvertIntToString(){
	x:= 1022
	s:= strconv.FormatInt(int64(x), 10)
	fmt.Println(s)
}