package main

import "fmt"
//Structure
type one int

type score struct{
        high string
        med string
        low string
}
func main(){
        var print one
        print=1
        xx := []int{2,4,5,6,21,} //slice
        fmt.Println("Slice xx is  = ",xx)

        //map
        m:= map[string]int{
              "Ex" : 22,
              "Si" : 25,
              "Another" : 90,
        }
        fmt.Println("Type one value is : ",print)
        fmt.Println(m)

        p1 := score{
              "High Score <",
              "Medium Score <",
              "Low Score ",
        }
        fmt.Println(p1)
}
