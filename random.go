package main

import "fmt"
import "math/rand"

func main(){
    var nums = []int32{1,1,1,1,1}

    for a := range nums{
      nums[a] = rand.Int31()
    }

    for i := range nums{
      fmt.Println(nums[i])
    }

}
