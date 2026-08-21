package main

import (
	examplesArrays "FirstProject/examplesArrays"
	examplesdefer "FirstProject/examplesDefer"
	examplesfunctions "FirstProject/examplesFunctions"
	examplesVariables "FirstProject/examplesVariables"
	"fmt"
)

func main() {
	//functionsExamples()
	//variablesExamples()
	//ArraysAndLoops()
	DeferExamples()
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

func variablesExamples() {
	examplesVariables.SayNameSurnameAndAge()
}

func ArraysAndLoops() {
	examplesArrays.LoopsArray()
	examplesArrays.MemoryLoop()
}

func DeferExamples() {
	examplesdefer.DoDefer()
	examplesdefer.DoDeferAnonimus()
}
