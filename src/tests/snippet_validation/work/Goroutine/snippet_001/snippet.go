package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max-min) + min
}

func main() {
    flags := [...]bool{true, false}
    totalStep := 10
    tortoiseStep := 0
    hareStep := 0
    fmt.Println("龜兔賽跑開始...")
    for tortoiseStep < totalStep && hareStep < totalStep {
        tortoiseStep++
        fmt.Printf("烏龜跑了 %d 步...\n", tortoiseStep)
        isHareSleep := flags[random(1, 10)%2]
        if isHareSleep {
            fmt.Println("兔子睡著了zzzz")
        } else {
            hareStep += 2
            fmt.Printf("兔子跑了 %d 步...\n", hareStep)
        }
    }
}
