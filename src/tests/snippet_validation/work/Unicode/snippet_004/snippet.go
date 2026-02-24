package main

import (
    "fmt"
    "unicode"
)

func allNumbers(s string) bool {
    for _, r := range []rune(s) {
        if !unicode.Is(unicode.Number, r) {
            return false
        }
    }
    return true
}

func main() {
    // true
    fmt.Println(allNumbers("²³¹¼½¾𝟏𝟐𝟑𝟜𝟝𝟞𝟩𝟪𝟫𝟬𝟭𝟮𝟯𝟺𝟻𝟼㉛㉜㉝ⅠⅡⅢⅣⅤⅥⅦⅧⅨⅩⅪⅫⅬⅭⅮⅯⅰⅱⅲⅳⅴⅵⅶⅷⅸⅹⅺⅻⅼⅽⅾⅿ"))
}
