package examplepointers

import "fmt"

func FirstExample() {
	x := 10
	p := &x
	fmt.Println(p, *p)
	//*p referencia de p extra o valor da memoria
	// numero da memora, 10
}

func Pointer(){
	x := 10
	Take(x)//copiando a variavel e passando por parametro
	fmt.Println(x)
	takeRealVariable(&x)
	fmt.Println(x)
}

func Take(x int){
	x = 100
}

func takeRealVariable(x *int){
	*x = 100
}

func ReturnEndMemoryPlaceValue() *int{
	x:= 10
	return &x
	// memoria vai ser deletada mais ainda assim retorna um valor sem quebrar
}