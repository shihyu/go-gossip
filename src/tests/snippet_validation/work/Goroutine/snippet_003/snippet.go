package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max-min) + min
}

func tortoise(totalStep int, wg *sync.WaitGroup) {
    defer wg.Done()

    for step := 1; step <= totalStep; step++ {
        fmt.Printf("烏龜跑了 %d 步...\n", step)
    }
}

func hare(totalStep int, wg *sync.WaitGroup) {
    defer wg.Done()

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
    wg := new(sync.WaitGroup)
    wg.Add(2)

    totalStep := 10

    go tortoise(totalStep, wg)
    go hare(totalStep, wg)

    wg.Wait()
}
