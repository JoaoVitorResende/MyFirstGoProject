package examplearraysandslices

import "fmt"

func FirstExampleSlice() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	arr[2] = 15
	fmt.Println(slice) //de 2 a 5 -> 2,3,4
	slice[0] = 123
	fmt.Println(arr) // -> 1, 123, 15, 4, 5 como o 0 no slice seria o 1 no array ele troca
}

func Slice() {
	slice := []int{1, 2, 3}
	fmt.Println(slice) // -> 1,2,3
}

func SliceRecivingType() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:5]
	fmt.Println(slice) // -> 1,2,3,4,5
	slice2 := arr[0:]
	fmt.Println(slice2) // -> 1,2,3,4,5
	slice3 := arr[:]
	fmt.Println(slice3) // -> 1,2,3,4,5
}

func SliceCapacityLenght() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	fmt.Println(slice, len(slice), cap(slice)) // -> 1,2,3,4,5, 5, 5
}

func SliceNull() {
	slice := []int{}
	fmt.Println(slice == nil)
}

var filmsDb = []string{"HarryPotter 1", "HarryPotter 2", "HarryPotter 3", "HarryPotter 4", "HarryPotter 5", "HarryPotter 6", "HarryPotter 6",
	"HarryPotter 7", "Percy Jackson 1", "Percy Jackson 2"}

func GetResultFromApi() {

	resultApi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	films := make([]string, 0, 10)

	for _, id := range resultApi {
		film := filmsDb[id]
		films = append(films, film)
	}

	fmt.Println(films)
}

func GetResultFromApiWithMatrix() {
	resultApi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//matrix2D := [][]int{}
	//matrix3D := [][][]int{}
	films := make([]string, 0, 10)

	for _, id := range resultApi {
		film := filmsDb[id]
		films = append(films, film)
	}

	fmt.Println(films)
}
