package exampleReaderAndWriters

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func FirstExample() {
	str := "hello world"
	reader := strings.NewReader(str)

	buffer:= make([]byte, 10000)
	n, err := reader.Read(buffer)

	if err != nil{
		panic(err)
	}

	fmt.Println(n)
	//traduzindo para humanos
	// buffer[:n] :n limitacao de n caracteres
	fmt.Println(string (buffer[:n]))
}

func ReadThStringTillEnd() {
	//essa classe faz com que 2 bytes sejam lidos a cada loop
	// h,e break, l,l break
	str := "hello world"
	reader := strings.NewReader(str)
	buffer:= make([]byte, 2)

	for {
		n, err := reader.Read((buffer))
		
		if err != nil {
			if errors.Is(err, io.EOF){
				break
			}
			panic(err) 
		}
		fmt.Println(n , string (buffer[:n]))
	}
}

type myWriter struct{}

func (myWriter) writeExample(b []byte) (int,error) {
	fmt.Print(string(b))
	return len(b), nil
}

func ReadAndWrite() {
	str := "hello world"
	reader := strings.NewReader(str)
	write := myWriter{}
	buffer:= make([]byte, 2)

	for {
		n, err := reader.Read((buffer))
		
		if err != nil {
			if errors.Is(err, io.EOF){
				break
			}
			panic(err) 
		}
		_,_= write.writeExample(buffer[:n])
	}
}

/*
io.ReadFull → “Read exactly N bytes into this buffer.”
io.ReadAll → “Read everything until the reader reaches EOF.”
*/
func ReadFull(){
	str := "hello world"
	reader := strings.NewReader(str)
	buf := make([]byte, len(str))
	n, err := io.ReadFull(reader, buf)
	fmt.Println(string(buf))
	fmt.Println(n, err)   
}