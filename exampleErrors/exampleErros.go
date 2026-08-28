package exampleerrors

import (
	"errors"
	"fmt"
	"math"
)

func Dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("nao pode dividir por zero")
	}
	return a / b, nil
}

type User struct {
	name string
}

func (u User) SayUserName() {
	fmt.Println(u.name)
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("um erro")
	}
	return &User{}, nil
}

type SqrError struct {
	msg string
}

func (s SqrError) Error() string {
	return s.msg
}

func RaizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrError{"Nao foi possivel fazer a raiz quadrada desse item"}
	}
	return math.Sqrt(x), nil
}

var errorNotFound = errors.New("not found")

func returnNil() error{
	return nil
}

func returnSqrError() error{
	return SqrError{msg: "erro de sqr"}
}

func FindTypeOfErrorIs(){
	err := returnNil()
	if err != nil && errors.Is(err, errorNotFound){
		fmt.Println(errorNotFound)
		return
	}
	fmt.Println("erro desconhecido")
}

func FinTypeOfErrorAs(){

	err := returnSqrError()
	var sqrErrorV SqrError

	if err != nil && errors.As(err, &sqrErrorV){
		fmt.Println(sqrErrorV.msg)
	}
	fmt.Println("erro desconhecido")
}

func ErrorWrapping(){
	err:= findTypeError()
	if err != nil && errors.Is(err, anyError){
		 fmt.Println("deu erro:", err)
	}
}

func findTypeError() error{
	err:= bar()
	if err != nil{
		return fmt.Errorf("findTypeError %w ", err)
	}
	return nil
}

var anyError = errors.New("error")

func bar() error{
	return anyError
}