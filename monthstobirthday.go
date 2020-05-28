// http://terokarvinen.com/2020/go-programming-course-2020-w22/
// https://gobyexample.com/
// https://stackoverflow.com/questions/4278430/convert-string-to-integer-type-in-go


package main
import (
        "fmt"
        "flag"
        s "strings"
        "strconv"
        "time"

)

func main() {

        var name string
        var birthday string
        flag.StringVar(&name, "name","none","give your name")
        flag.StringVar(&birthday, "birthday", "none.none ","give your birthday day.month.year")
        flag.Parse()


        p := s.Split(birthday, ".")
        bdmonthstr := p[1]

        bdmonth, err:= strconv.Atoi(bdmonthstr)
        if err != nil {
                fmt.Println("error:", err)
                return
        }

        month := time.Now().Month()
        monthnow := int(month)


        var monthsTillBirthday int

        if bdmonth > monthnow {
                monthsTillBirthday =  bdmonth - monthnow
        } else if monthnow > bdmonth {
                monthsTillBirthday =  (12-monthnow) +  bdmonth
        }

        fmt.Println("Hello", name, "you have about", monthsTillBirthday, "months till your birthday")



}

