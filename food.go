package main

import "fmt"

func main() {
  var food string
  
  fmt.Println("What would you like to eat?")
  fmt.Scan(&food)
  
  fmt.Printf("Sure, we can have %v.\n", food)
  
}
