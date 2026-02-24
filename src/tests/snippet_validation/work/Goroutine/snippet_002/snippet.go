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

func tortoise(totalStep int) {
    for step := 1; step <= totalStep; step++ {
        fmt.Printf("烏龜跑了 %d 步...\n", step)
    }
}

func hare(totalStep int) {
    flags := [...]bool{true, false}
    step := 0
    for step < totalStep {
        isHareSleep := flags[random(1, 10)%2]
        if isHareSleep {
            fmt.Println("兔子睡著了zzzz")
        } else {
            step += 2
            fmt.Printf("兔子跑了 %d 步...\n", step)
        }
    }
}

func main() {
    totalStep := 10

    go tortoise(totalStep)
    go hare(totalStep)

    time.Sleep(5 * time.Second) // 給予時間等待 Goroutine 完成
}
