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


