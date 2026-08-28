package exampleerrors

import (
	"errors"
	"fmt"
)



func Dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("nao pode dividir por zero")
	}
	return a/b, nil
}

type User struct{
	name string
}

func (u User) SayUserName(){
	fmt.Println(u.name)
}

func NewUser(wantErr bool) (*User, error){
	if (wantErr){
		return nil, errors.New("um erro")
	}
	return &User{}, nil
}