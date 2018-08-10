package main

import "fmt"

type person struct{
    name string
    hair string
    houseNum int
    married bool
}

func main(){
    x := person{"Alice", "black", 9999, false}
    fmt.Println(x)
    fmt.Println(x.changePerson())
}

//sets person struct to bob
func (p1 person) changePerson() person{
    p1.name = "Bobby"
    p1.hair = "blonde"
    p1.houseNum = 9876
    p1.married = false
    return p1
}
