package main

import "fmt"

func main(){
    var name string
    name = "Johnny Appleseed"
    fmt.Println("Initial name: " + name)

    pointer := &name
    *pointer = "Hello"
    fmt.Println("Name changed.")
    fmt.Println("Dereferenced pointer to name: " + *pointer)
    fmt.Println("Address of name and address of dereferenced pointer:")
    fmt.Println(&name)
    fmt.Println(&*pointer)
}
