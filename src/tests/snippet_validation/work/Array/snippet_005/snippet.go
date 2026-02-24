package main

import "fmt"

func change(arr [3]int) [3]int {
    arr[0] = 10
    return arr
}

func main() {
    arr1 := [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
    fmt.Println(arr1) // [[1 2 3] [4 5 6]]

    arr2 := [...][3]int{[...]int{1, 2, 3}, [...]int{4, 5, 6}}
    fmt.Println(arr2) // [[1 2 3] [4 5 6]]

    arr3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
    fmt.Println(arr3) // [[1 2 3] [4 5 6]]

    arr4 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
    fmt.Println(arr4) // [[1 2 3] [4 5 6]]
}
