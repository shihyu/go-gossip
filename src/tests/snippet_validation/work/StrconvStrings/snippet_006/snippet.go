package mypackage

import (
    "testing"
    "strings"
)

func plusAppend() string {
    c := ""
    for i := 0; i < 100000; i++ {
        c += "test"
    }
    return c
}

func buliderAppend() string {
    var b strings.Builder
    for i := 0; i < 100000; i++ {
        b.WriteString("test")
    }
    return b.String()
}

func BenchmarkPlusAppend(b *testing.B) {
    for i := 0; i < b.N; i++ {
        plusAppend()
    }
}

func BenchmarkBuilderAppend(b *testing.B) {
    for i := 0; i < b.N; i++ {
        buliderAppend()
    }
}
