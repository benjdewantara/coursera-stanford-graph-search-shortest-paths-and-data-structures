/*package main

import "fmt"

func main() {
	var asdf = "asdf"
	var intArr int[3]
	fmt.Println("asdf = " + asdf)
	fmt.Println("Hell on earth")
	//file, _ := ioutil.ReadFile("./_410e934e6553ac56409b2cb7096a44aa_SCC.txt")
	//fmt.Println(file)
}
*/

package main

import (
	"fmt"
	"strings"
	"time"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	a := 3 / 2
	fmt.Println("a=", a)
}

func main1() {
	fmt.Println(Sqrt(2))
	time.Now().Hour()
}

func main2() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	fmt.Println(b[0:2])

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func mainArrMake() {
	//func main() {
	var arr = make([][]int, 3)
	fmt.Println(arr)
}

func mainForRange() {
	//func main1() {

	//for idx, e := range []string{"asdf", "qwer"} {
	for idx, e := range strings.Fields("asdf qwer") {
		fmt.Println("idx =", idx, "e =", e)
	}
	if !false {
		fmt.Println("zxcvzxcv")
	}
}
