package main

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a0 := 0
	a1 := 1
	return func() int {
		a0Before := a0
		a0 = a1
		a1 = a0Before + a1
		return a0
	}
}

/*func mainFibo() {}
//func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}*/
