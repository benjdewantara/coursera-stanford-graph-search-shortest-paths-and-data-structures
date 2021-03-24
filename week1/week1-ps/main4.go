package main

import (
	"fmt"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(b []byte) (int, error) {
	counter := 0
	for ; counter < len(b); counter++ {
		b[counter] = 'A'
	}
	return counter, nil
}

func main() {
	var r = MyReader{}
	var b = make([]byte, 20)
	fmt.Println("b before", b)
	c, _ := r.Read(b)
	fmt.Println("counted = ", c)
	fmt.Println("b after", b)
	fmt.Printf("b = %s", b)
	//reader.Validate(MyReader{})
}
