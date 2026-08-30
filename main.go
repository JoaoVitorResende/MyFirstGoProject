package main

import (
	examplearraysandslices "FirstProject/exampleArraysAndSlices"
	exampleerrors "FirstProject/exampleErrors"
	examplegoruntime "FirstProject/exampleGoRunTime"
	examplemethodandstructs "FirstProject/exampleMethodAndStructs"
	"FirstProject/exampleParametersGenerics"
	"FirstProject/exampleReaderAndWriters"
	exampleinterfaces "FirstProject/exampleinterfaces"
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
	//DeferExamples()
	//SliceExamples()
	//structExample()
	//interfacesExamples()
	//errorsExample()
	//exampleReadAndWrite()
	//exampleParamsGeneric()
	exampleGoRunTimeFunctions()
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

func SliceExamples(){
	examplearraysandslices.GetResultFromApi();
}

func structExample(){
	examplemethodandstructs.MyFirstStruct()
	//examplemethodandstructs.StructWithTag()
}

func interfacesExamples(){
	exampleinterfaces.CallValidationAnimal()
}

func errorsExample(){
	/*a := 10
	b := 0
	res, error := exampleerrors.Dividir(a,b)
	fmt.Println(res, error)
	//--
	user,error := exampleerrors.NewUser(true)
	if (error != nil){
		fmt.Println("aconteceu algum erro")
	}else{
		user.SayUserName()
	}
	//--
	x:= -10.0
	res, erro := exampleerrors.RaizQuadrada(x)
	if erro != nil{
		fmt.Println(erro)
	}
	fmt.Println(res)*/
	exampleerrors.FindTypeOfErrorIs()
	exampleerrors.FinTypeOfErrorAs()
	exampleerrors.ErrorWrapping()
}

func exampleReadAndWrite(){
	//exampleReaderAndWriters.FirstExample()
	//exampleReaderAndWriters.ReadThStringTillEnd()
	//exampleReaderAndWriters.ReadAndWrite()
	exampleReaderAndWriters.ReadFull()
}

func exampleParamsGeneric(){
	exampleParametersGenerics.FirstExample()
}

func exampleGoRunTimeFunctions(){
	//examplegoruntime.FirtAsyncExample()
	examplegoruntime.WorkingAsync()
}