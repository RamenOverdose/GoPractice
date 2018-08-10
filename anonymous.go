package main

import "fmt"
import "strings"

const myString = "maybe"

func main(){
  var input string
  var input2 string
  fmt.Print("Enter a string: ")
  fmt.Scanln(&input)
  fmt.Print("Enter a second string: ")
  fmt.Scanln(&input2)

  switch (strings.Compare(input, input2)){
    case 1:
      fmt.Println(input + " is greater than " + input2)
      break
    case -1:
      fmt.Println(input + " is less than " + input2)
      break
    default:
      fmt.Println("YOUR TWO STRINGS ARE EQUAL!")
  }
}
