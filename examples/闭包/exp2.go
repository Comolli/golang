package main

import "fmt"

func main() {
   var funs []func() int

   for i := 0; i < 10; i ++ {
      funs = append(funs, num(i))
   }
   fmt.Print(funs[0](), funs[2]())
}

func num(i int) func() int {
   return func() int {
      return i
   }
}
