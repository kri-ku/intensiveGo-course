//https://golang.org/pkg/sort/#Sort
//https://gobyexample.com/sorting
//http://terokarvinen.com/2016/publish-your-project-with-github

package main
import (
	"fmt"
	"sort"
)


func main() {

fmt.Println("Give me five numbers to sort!")
var number1 int
var number2 int
var number3 int
var number4 int
var number5 int

fmt.Scanln(&number1, &number2, &number3, &number4, &number5)

listofnumbers := [] int {number1, number2, number3, number4, number5}
sort.Ints(listofnumbers)

fmt.Println("Sorted it:", listofnumbers,"!")


}
