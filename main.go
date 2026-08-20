package main

import (
	examplesfunctions "FirstProject/examplesFunctions"
	examplesVariables "FirstProject/examplesVariables"
	examplesArrays "FirstProject/examplesArrays"
	"fmt"
)

func main() {
	//functionsExamples()
	//variablesExamples()
	ArraysAndLoops()
}

func functionsExamples() {
	fmt.Println("hello world")
	examplesfunctions.SayHi()
	fmt.Println(examplesfunctions.Add(1, 1))
	a, b := examplesfunctions.Swap(10, 20)
	fmt.Println(a, b)
	res, rem := examplesfunctions.Divide(4, 2)
	fmt.Println(res, rem)
	funct := examplesfunctions.AddHighOrder(1)
	result := funct(2)
	fmt.Println(result)
	fmt.Println(examplesfunctions.AddVariant(10, 10, 10))
}

func variablesExamples(){
	examplesVariables.SayNameSurnameAndAge()
}

func ArraysAndLoops(){
	//examplesArrays.LoopsArray();
	//examplesArrays.MemoryLoop();
}