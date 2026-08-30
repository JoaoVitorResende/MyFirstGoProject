package examplegoruntime

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
	"errors"
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

// contextos
func WorkingAsyncContext(){
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	/*
	O pacote context em Go serve para controlar
	o tempo de vida de tarefas, cancelar operações em
	andamento, impor limites de tempo (timeouts) e passar
	dados de uso único entre funções e goroutines. 
	*/
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	// criando localhost server para testar a requisicao
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(10 * time.Second)//aguarda 10 segundos
		fmt.Println(w, "hello word")//imprime hello word apos tempo
	}),)

	for range n{
		go func(ctx context.Context){ //parte que deixa assincrona
			defer wg.Done() // parte importante await e defer para devolver tudo quando terminar
			req, err := http.NewRequestWithContext(ctx, "GET", server.URL, nil,)
			resp, err := http.DefaultClient.Do(req)
			if err != nil{
				if errors.Is(err, context.DeadlineExceeded){
					fmt.Println("timeout")
					return
				}
				panic(err)
			}
			defer resp.Body.Close()
		}(ctx)
	}
	/*
	faz com que a funcao trave ate receber 10 sinais
	e entao faz o restante await do c#
	mas ainda fazendo paralelismo
	*/
	wg.Wait()
	fmt.Println(time.Since(start))
}