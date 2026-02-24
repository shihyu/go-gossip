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

func tortoise(totalStep int, goal chan string) {
    for step := 1; step <= totalStep; step++ {
        fmt.Printf("烏龜跑了 %d 步...\n", step)
    }
    goal <- "烏龜"
}

func hare(totalStep int, goal chan string) {
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
    goal <- "兔子"
}

func main() {
    goal := make(chan string)

    totalStep := 10

    go tortoise(totalStep, goal)
    go hare(totalStep, goal)

    fmt.Printf("%s 抵達終點\n", <-goal)
    fmt.Printf("%s 抵達終點\n", <-goal)
}
