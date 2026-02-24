package main

import "fmt"

type Stack[T any] struct {
    elems []T
}

func (s *Stack[T]) Push(v T) {
    s.elems = append(s.elems, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elems) == 0 {
        var zero T
        return zero, false
    }
    i := len(s.elems) - 1
    v := s.elems[i]
    s.elems = s.elems[:i]
    return v, true
}

func main() {
    var s Stack[string]
    s.Push("Go")
    s.Push("1.26")
    v, _ := s.Pop()
    fmt.Println(v) // 1.26
}
