package main
import "testing"

func TestAdder(t *testing.T){
  var actual = Adder(4,5)
  if actual != 9 {
    t.Errorf("Expected 9 got %v", actual)
  }
}
