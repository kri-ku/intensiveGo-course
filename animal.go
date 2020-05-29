//http://terokarvinen.com/2020/go-programming-course-2020-w22/
// https://gobyexample.com/command-line-flags

package main

import (

	"fmt"
	"flag"

)



func main() {

        var name string
        var age int
        var animal string



        flag.StringVar(&name,"name", "","give your name")
        flag.IntVar(&age, "age", 0, "give your age")
        flag.StringVar(&animal, "animal", "", "do you have animals y/n?")
        flag.Parse()


        fmt.Println("Here is Hello for all of your years!")

        for i :=0; i < age; i++ {
        	fmt.Println((i+1),"Hello", name, "you are awesome")
	}

        fmt.Println("Total:", age, "times hello to you gorgeus!")

        if animal == "y" {
        	fmt.Println("you are so lucky to have an animal companion!")
        } else {
        	fmt.Println("oh no no pets surullinen")

	}

}
