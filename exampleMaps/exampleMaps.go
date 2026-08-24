package examplemaps

import "fmt"

func FirstExampleMap() {
	var m map[string]string
	fmt.Println(m == nil)

	m2 := make(map[string]string) //creating
	fmt.Println(m2 == nil)

	m3 := make(map[string]string, 100)// giving 100 spaces on memory
	fmt.Println(m3 == nil)
}

func MapWithValues(){
	m := map[string]string{
		"Pedro" : "pessoa",
		"joao": "pessoa",
	}
	fmt.Println(m)
}

func MapWithSlices(){
	m:= map[string][]int{
		"Pedro" :{1,2,3},
	}
	fmt.Println(m);
}

func GetValueFromMap(){
	m := make(map[string]string)
	m["pedro"] = "pessoa"
	valor := m["pedro"]
	fmt.Println(valor)//vai imprimir pessoa

	valor2, ok:= m["pedro"]
	fmt.Println(valor2, ok)// pessoa e true, pois o ok valida que a chave existe no map
}

func DeletingAnItemOnMap(){
	m := make(map[string]string)
	m["pedro"] = "pessoa"
	delete(m, "pedro")
	valor, ok:= m["pedro"]
	fmt.Println(valor, ok)
}

func ClearingMap(){
	m := map[string]string{
		"pedro":"pessoa",
		"joao": "Resendde",
	}
	clear(m)
	fmt.Println(m)
}

func RunningMaps(){
	m := map[string]string{
		"Pedro":"pessoa",
		"joao":"resende",
	}

	for key, value := range m{ //mapas em go sao desordenados entao n precisa vir na ordem que eesta ali
		fmt.Println(key,value)
	}

	for key := range m{ //Deletando a chave pedro
		if key == "pedro"{
			delete(m, key)
		}
	}

}
 