package mymath

import "testing"

func TestSomething(t *testing.T) {
    t.Skip()
}

func TestAdd(t *testing.T) {
    if Add(1, 2) == 3 {
        t.Log("mymath.Add PASS")
    } else {
        t.Error("mymath.Add FAIL")
    }
}
