package exampleinterfaces

import "fmt"

func ReciveInterface(example any) {

}

func CallStructUsingInterface(){
	dog := &Dog{}
	WhatDoesThisAnimalSay(dog)
}

type Animal interface {
	Sound() string
}

type Dog struct{}
type Cat struct{}

func (d *Dog) Sound() string {
	return "au au !"
}

func (c *Cat) Sound() string {
	return "meow!"
}

func WhatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

func InterfaceDog(){
	//in go land dosen't need to extends interface is implicity 
	var a Animal
	var dog *Dog
	a = dog
	fmt.Println(a.Sound())
}

func ValidatingAnyType(a any){
	str, ok := a.(string)
	fmt.Println(str, ok)
}

func CallValidationAnimal(){
	dog := &Dog{}//cria e aponta diretiva para a memoria
	ValidatingAnimal(dog)//passa o valor
}

func ValidatingAnimal(a Animal){
	//checa e analisa o resultado e faz o som que deveria sair
	switch t := a.(type){
		case *Dog:
			fmt.Println(t.Sound())
		case *Cat:
			fmt.Println(t.Sound())
	}
}