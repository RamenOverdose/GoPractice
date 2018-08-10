package main

import "fmt"

type person struct{
  name string
  hair string
  houseNum int
  married bool
}

func main(){
  var p1 = person{"bob", "brown", 10101, true}
  fmt.Println(p1.name)
  fmt.Println(p1.houseNum)
  fmt.Println(p1.married)

  pointer := &p1
  fmt.Println(*pointer)
  fmt.Println(pointer.houseNum)
  pointer.houseNum = 12345
  fmt.Println(p1.houseNum)

  stuff()
}

func stuff(){
  var a, b, c = 139, 714, 441
  var p1, p2, p3 = &a, &b, &c

  var pSlice[]*int
  pSlice = append(pSlice, p1)
  fmt.Println(*pSlice[0])
  pSlice = append(pSlice, p2)
  fmt.Println(*pSlice[1])
  pSlice = append(pSlice, p3)
  fmt.Println(*pSlice[2])

  //prints addresses because pSlice is slice of pointers
  fmt.Println(pSlice)
}
