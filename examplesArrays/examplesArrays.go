package examplesarrays

import "fmt"

func SimpleArray() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
}

func IndexArrayCustom(){
	arr := [5]int{4: 300, 2: 200}
	fmt.Println(arr)
}

func LoopsArray(){
	arr := [3]int{1, 2, 3}

	for i:= 0; i < len(arr); i++{
		fmt.Print(arr[i])
	}

	var j int

	for j < 10 {
		fmt.Print(j)
		j++
	}
}

func RangeExample(){
	arr:=[10] int{}
	for range arr{
		fmt.Print("isso seria um loop")
	}

	for i, elem := range arr{
		fmt.Print(i, elem)
	}

	for _, elem := range arr{
		fmt.Print(elem)
	}
}

func LoopNewWay(){
	for range 10{
		fmt.Print("isso seria um loop")
	}
}

func MemoryLoop(){
	arr := [3]int{1,2,3}
	for i, elem := range arr{
		fmt.Print(&i, &elem)
	}
}
