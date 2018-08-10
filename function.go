package main

import "fmt"
import "strings"

func main(){
  fmt.Println(arithmetic(100, 200))
  fmt.Println(compareStrings("apple", "orange"))
}

func arithmetic(a float32, b float32) (float32, float32){
  var sum float32
  var average float32

  sum = a + b
  average = sum / 2

  return sum, average
}

func compareStrings(first, second string) int{
  return strings.Compare(first, second)
}
