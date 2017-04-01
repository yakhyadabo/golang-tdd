package main

import (
  "testing"
)

var calculator = NewCalculator()

func TestAdd(t *testing.T) {
  v := calculator.Add(1,2) 
  if v != 3 {
    t.Error("Expected 3, got ", v)
 }

}

func TestSum(t *testing.T) {
  v := calculator.Sum(1,2,3,4) 
  if v != 10 {
    t.Error("Expected 10, got ", v)
 }

}

func TestMult(t *testing.T) {
 // var v = calculator.Mult(1,2) 
  v := calculator.Mult(1,2) 
  if v != 2 {
    t.Error("Expected 2, got ", v)
 }

}

func TestFact(t *testing.T) {
  v := calculator.Fact(3) 
  if v != 6 {
    t.Error("Expected 6, got ", v)
 }

}

func TestFactZero(t *testing.T) {
  v := calculator.Fact(0) 
  if v != 1 {
    t.Error("Expected 1, got ", v)
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
