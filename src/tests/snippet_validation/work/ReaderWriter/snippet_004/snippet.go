package main

import (
    "fmt"
    "io"
    "os"
)

func write(w io.Writer, buf []byte, n int) (err error) {
    nw, ew := w.Write(buf[:n])
    if ew != nil {
        return ew
    }
    if n != nw {
        return io.ErrShortWrite
    }
    return nil
}

func Copy(w io.Writer, r io.Reader) (err error) {
    buf := make([]byte, 32 * 1024)
    for {
        nr, er := r.Read(buf)
        if nr > 0 {
            err = write(w, buf, nr)
            if err != nil {
                return
            }
        }
        if er != nil {
            if er != io.EOF {
                err = er
            }
            return
        }
    }
}

func main() {
    fmt.Print("檔案來源：")
    var filename string
    fmt.Scanf("%s", &filename)

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    Copy(os.Stdout, f)
}
