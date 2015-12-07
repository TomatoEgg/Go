package main

import "fmt" //We can do it in one-line

//Or if we have multiple imports, they can be grouped
//import (
//  "fmt"
//)


//Returning a closure for generating fibonacci sequence
func fibi() func() int {
  a, b := 1, 1 //initial values
  
  return func() int {
    next := a //keep a copy
    a, b = b, a + b //generate new values
    return next
  }
}

func main() {
  //Print the first 10 numbers from fibonacci sequence
  i := 0
  fibiFunc := fibi()
  for ;i < 10;i++ {
    fmt.Printf("fibi(%d)=%d, ", i, fibiFunc())
  }
}

