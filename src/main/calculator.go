package main
import (
  "fmt"
  "errors"
)

type Calculator struct {
 //  op1, op2 int
}

func (Calculator) Mult(x,y int) int {
  return  x * y
}

func (Calculator) Add(x,y int) int {
  return  x + y
}

func (Calculator) Div(x,y int) (int, error) {
  if y==0 {
     return -1, errors.New(fmt.Sprintf("could not divide by zero"))
  }
  return  x/y, nil
}

func main() {
  c := Calculator {}
  fmt.Println("Mult : ", c.Mult(2,4))
  fmt.Println("Add : ", c.Add(2,4))
  // fmt.Println("Div : ", c.Div(4,3))
}