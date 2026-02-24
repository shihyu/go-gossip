package main

import "os"

func main() {
    buf := make([]byte, 5);
    os.Stdout.Write([]byte("輸入五個數字："))
    os.Stdin.Read(buf)
    os.Stdout.Write(buf)
}   
