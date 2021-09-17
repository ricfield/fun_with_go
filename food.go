package main

import "fmt"

func main() {
  fmt.Println("What would you like to eat?")
  
  var food string
  fmt.Scan(&food)
  
  fmt.Printf("Sure, we can have %v.\n", food)
  
}
