package main

import (
  "testing"
  "../main"
)

var  calculator  = main.Calculator{}

func TestMult(t *testing.T) {
 // var v = calculator.Mult(1,2) 
  v := calculator.Mult(1,2) 
  if v != 2 {
    t.Error("Expected 2, got ", v)
 }

}


func TestAdd(t *testing.T) {
  v := calculator.Add(1,2) 
  if v != 3 {
    t.Error("Expected 3, got ", v)
 }

}

func TestDiv(t *testing.T) {
  v,e := calculator.Div(4,2) 
  if v != 2 || e != nil {
    t.Error("Expected 2, got ", v)
 }

}


func TestDivByZero(t *testing.T) {
  v,e := calculator.Div(4,0) 
  if e == nil {
    t.Error("Exception expected, got ", v)
 }

}
