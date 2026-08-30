package exampleParametersGenerics

import (
	"fmt"
	"slices"
)

func FirstExample() {
	generic(1)
	generic(" ola")
	generic([]int{1})
}

func generic[T any](arg T) {
	fmt.Println(arg)
}

//so variaveis que sao comparaveis

func comparableMethod[T comparable](arg T){
	fmt.Println(arg)
}
// exemplo meu tipo onde vai implementar a interface
func CallMyTypeExample(){
	var mt myType = ""
	comparableItem(mt)
}

type myType string

func (myType) comparableItem() {}

type MyConstrant interface {
	comparableItem()
}

func comparableItem[T MyConstrant](arg T){
	fmt.Println(arg)
}
// ---

// generic method
func CallMySecondConstraint(){
	reciver(1)
	reciver("")
	reciver([]int{2})
}

type MySecondConstraint interface{
	int | string | []int
}

func reciver[T MySecondConstraint](arg T){
	fmt.Println(arg)
}
//---

// meu tipo 3

func CallMyThirdConstraint(){
	var mt3 myType3 = ""
	reciver2(1)
	reciver2(mt3)
	reciver2([]int{2})
	// item do proprio go
	//constraints.Float
}

type myType3 string

type MyThirdConstraint interface{
	int | ~string | []int
}

func reciver2[T MyThirdConstraint](arg T){
	fmt.Println(arg)
}
//---

// myStruct limitation
	/*
		func CallMyFourthConstraint (){
			var ms MyStruct[string] = MyStruct[string]{}
		}

		type MyStruct[T any] struct{
			reciver2 T
		}
		func reciver3[T MyThirdConstraint](arg T){
			fmt.Println(arg)
		}

		func (MyStruct[T]) reciver3() {}
	*/
//-------

// praticle example

// contains pacote go igual ao custom
func ContainsGo(){
	slices.Contains([]int{2}, 2)
}

// contains custom 
func ContainsCustom[T comparable](s [] T, cmp T) bool{
	for _, str := range s{
		if str == cmp{
			return true
		}
	}
	return false
}


//---