package variablesexamples

import "fmt"

var age int
var name, surname string = "Pedro", "Antunes"

func SayNameSurnameAndAge() {
	fmt.Println(name, surname, age)
	name2 := "joao"
	surname2 := "resende"
	fmt.Println(name2, surname2)
}

func ReturnBoolean() bool{
	var isThisAboolean = true;
	return isThisAboolean
}

