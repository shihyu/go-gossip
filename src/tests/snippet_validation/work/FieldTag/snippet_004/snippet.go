package main

import (
    "encoding/json"
    "fmt"
)

type Customer struct {
    Name string `json:"name"`
    City string `json:"city"`
}

func main() {
    cust := Customer{"Justin", "Kaohsiung"}
    b, _ := json.Marshal(cust)
    // 顯示 {"name": "Justin","city": "Kaohsiung"}
    fmt.Println(string(b))
}
