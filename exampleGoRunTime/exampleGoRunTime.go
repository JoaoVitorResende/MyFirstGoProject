package examplegoruntime

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func FirtAsyncExample() {
	start := time.Now()
	for range 10{
		go func(){ //parte que deixa assincrona
			resp, err := http.Get("https://www.google.com/")
			if err != nil{
				panic(err)
			}
			defer resp.Body.Close()
			fmt.Println("ok")
		}()
	}
	fmt.Println(time.Since(start))
}

func WorkingAsync(){
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for range n{
		go func(){ //parte que deixa assincrona
			defer wg.Done() // parte importante await e defer para devolver tudo quando terminar
			resp, err := http.Get("https://www.google.com/")
			if err != nil{
				panic(err)
			}
			defer resp.Body.Close()
			fmt.Println("ok")
		}()
	}
	/*
	faz com que a funcao trave ate receber 10 sinais
	e entao faz o restante await do c#
	mas ainda fazendo paralelismo
	*/
	wg.Wait()
	fmt.Println(time.Since(start))
}