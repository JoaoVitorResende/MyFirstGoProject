package examplemethodandstructs

import "fmt"

func MyFirstStruct() {
	//pointer interaction
	user := User{"joao", 1}
	fmt.Println(user)
	fmt.Println(user.ID)
	fmt.Println(user.NAME)
	user.myFirstMethod()
	updateNameWithPointer(&user,"Luan")
}

func StructWithTag(){
	item := Item{"celular", 1}
	fmt.Println(item)
}

type User struct {
	NAME string
	ID   uint64
}
// u seria igual a this
func (u User) myFirstMethod(){
	fmt.Println(u.NAME)
}
//adicionando o ponteiro para o lugar verdadeiro
func (u *User) UpdateNameMethod(newName string){
	u.NAME = newName
	fmt.Println(u.NAME)
}
// precisa de um metodo apontando com ponteiro para alterar aqui se nao seria erro
func updateNameWithPointer(u *User, newName string){
	u.NAME = newName
}
//structs tags
// ele vai seguir as tags json nesse caso vai estar minusculo
type Item struct{
	NAME string `json:"name"`
	ID uint64 `json:"id"`
}