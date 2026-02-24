<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 位元組構成的字串

</div>

  

在〈[認識預定義型態](http://openhome.cc/Gossip/Go/PreDeclaredType.html)〉中略略談過字串，表面看來，用雙引號（`"`）或反引號（\`）括起來的文字就是字串，預設型態為 `string`，實際在 Go 中，字串是由唯讀的 UTF-8 編碼位元組所組成。

# 字串入門

先從簡單的開始，在 Go 原始碼中，如果你撰寫 `"Go語言"` 這麼一段文字，那麼會產生一個字串，預設型態為 `string`，字串是唯讀的，一旦建立，就無法改變字串內容。

使用 `string` 宣告變數若無指定初值，預設是空字串 `""`，可以使用 `+` 對兩個字串進行串接，由於字串是唯讀的，因此實際上串接的動作，會產生新的字串，如果想比較兩個字串的相等性，可以使用 `==`、`!=`、`<`、`<=`、`>`、`>=` 依字典順序比較。

``` go
package main

import "fmt"

func main() {
    text1 := "Go語言"
    text2 := "Cool"
    var text3 string
    fmt.Println(text1 + text2) // Go語言Cool
    fmt.Printf("%q\n", text3)  // ""
    fmt.Println(text1 > text2) // true
}
```

上面的例子中，由於使用 `fmt.Println` 顯示空字串時看不到什麼，因此改用 `fmt.Printf`，並使用 `%q` 來脫離無法顯示的字元。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

使用 `""` 時不可換行，如果你的字串想要換行，方法之一是分兩個字串，並用 `+` 串接。例如：

``` go
    text := "Go語言" +
            "Cool" 
```

另一個方式是以重音符 \` 定義字串，例如：

``` go
package main

import "fmt"

func main() {
    text := `Go語言
                 Cool`
    fmt.Printf("%q\n", text) // "Go語言\n                  Cool"
}
```

使用 \` 定義的字串，會完全保留換行與空白，因此，在上頭你可以看到被保留的換行與空白字元，如果使用 `fmt.Println(text)`，顯示時也會看到對應的換行與空白。使用 \` 定義的字串，也不會轉譯字元，例如：

``` go
package main

import "fmt"

func main() {
    text := `Go語言\nCool`
    fmt.Println(text)  // Go語言\nCool
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

在這個例子中可以看到，使用 \` 時，不會對 `\n` 做轉譯的動作，因此，你會直接看到顯示了「Go語言\nCool」。

在 Go 中可以使用的轉譯有：

- `\a`：U+0007，警示或響鈴
- `\b`：U+0008，倒退（backspace）
- `\f`：U+000C，饋頁（form feed）
- `\n`：U+000A，換行（newline）
- `\r`：U+000D，歸位（carriage return）
- `\t`：U+0009，水平 tab
- `\v`：U+000b，垂直 tab
- `\\`：U+005c，反斜線（backslash）
- `\"`：U+0022，雙引號
- `\ooo`：位元組表示，o 為八進位數字
- `\xhh`：位元組表示，h 為十六進位數字
- `\uhhhh`：Unicode 點點表示，使用四個 16 進位數字
- `\Uhhhhhhhh`：Unicode 點點表示，使用八個 16 進位數字

# 唯讀位元組片段

那麼，想知道一個字串的長度該怎麼做呢？Go 中有個 `len` 函式，當它作用於字串時，結果可能會令一些從其他程式語言，像是 Java 過來的人感到訝異：

``` go
package main

import "fmt"

func main() {
    fmt.Println(len("Go語言")) // 8
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

顯示的結果是 8 而不是 4，給個提示，Go 的字串實作使用 UTF-8，是的！`len` 傳回的是位元組長度，因為 Go 的字串，本質上是 UTF-8 編碼後的位元組組成，如果你使用 `fmt.Printf("%x", "Go語言")`，會顯示 476fe8aa9ee8a880，47 是「G」的位元組以 16 進位數字表示的結果，6f 是 o，e8aa9e 是「語」的三個位元組分別以 16 進位數字表示的結果，e8a880 是「言」。

不單是如此，Go 中可以使用 `[]` 與索引來取得字串的位元組資料，是的，位元組！傳回的型態是 `byte`（`uint8`），`"Go語言"[0]` 取得的是 G 的位元組資料，`"Go語言"[1]` 取得的是 o 的位元組資料，`"Go語言"[2]` 呢？取得的是「語」的 UTF-8 實作中，第一個位元組資料，也就是 e8。可以用以下這個程式片段來印證：

``` go
package main

import "fmt"

func main() {
    text := "Go語言"
    for i := 0; i < len(text); i++ {
        fmt.Printf("%x ", text[i])
    }
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

雖然還沒正確介紹 `for` 迴圈，不過程式應該很清楚，用迴圈遞增的 `i` 值來取得指定索引處的資料，結果是顯示「47 6f e8 aa 9e e8 a8 80 」。

這個位元組序列是怎麼決定的？當你寫下 `"Go語言"`，你的 .go 原始碼檔案是什麼編碼呢？是的！UTF-8，Go 就是從這當中取得 `"Go語言"` 位元組序列，每個位元組就是 UTF-8 的一個碼元（code unit）。

雖說字串是唯讀的位元組片段，不過，實際的位元組是隱藏在字串底層，如果你想取得，必須轉為 `[]byte`，例如：

``` go
package main

import "fmt"

func main() {
    text1 := "Go語言"
    bs := []byte(text1)
    bs[0] = 103
    text2 := string(bs)
    fmt.Println(text1) // Go語言
    fmt.Println(text2) // go語言
}
```

注意，你不是真的取得字串底層的位元組資料，只是取得複本，因此，在範例中可以看到，雖然對 `text2` 的位元組做了修改，`text1` 是不受影響的，記得，字串是唯讀的，一旦建立，沒有方式可以改變其內容。

# string 與索引

實際上，Go 的字串支援片段操作，slice 操作時的索引是針對位元組，然而，傳回的型態還是 `string`，例如，`"Go語言"[0:2]`，傳回 `"Go"`，因為指定要切割出索引 0 開始，索引 2 結束（但不包括 2）的部份，也就是 47 與 6f 這兩個位元組，但是以 `string` 傳回。

那麼，如果是 `"Go語言"[2:3]` 呢？嗯，傳回的字串是 `"\xe8"`！這是什麼？事實上，Go 中的字串可以是任意位元組片段，因此，你可以如下定義字串：

``` go
package main

import "fmt"

func main() {
    text := "\x47\x6f\xe8\xaa\x9e\xe8\xa8\x80"
    fmt.Println(text)  // Go語言
}
```

片段操作時，如果省略冒號之後的數字，則預設取得至字串尾端的子字串，例如 `"Go語言"[3:]` 會傳回 `"\xaa\x9e\xe8\xa8\x80"` 的字串，如果省略冒號之前的數字，預設從索引 0 開始，例如 `"Go語言"[:2]` 會取得 `"Go"` 的字串，也就是 `"\x47\x6f"` 的字串，如果是 `"Go語言"[:]`，那麼就是取得全部字串內容了。

[`strings` 套件](https://pkg.go.dev/strings/) 中有不少字串可用的方法，想做字串操作時，可以多加利用，不過要看清楚是針對什麼在操作。例如 `strings.Index`：

``` go
package main

import "fmt"
import "strings"

func main() {
    text := "Go語言"
    fmt.Printf("%d\n", strings.Index(text, "言"))  // 5
}
```

傳回的索引值是 5 而不是 3，這是因為 `"言"` 的第一個位元組，是在 `"Go語言"` UTF-8 編碼後的位元組組成中第 5 個索引位置。

問題來了，如果對於 `"Go語言"`，想逐一取得 `'G'`、`'o'`、`'語'`、`'言'` 該怎麼辦？當然不能用 `text[n]`，這只會取得第 n 個位元組，可以將字串型態轉換為 `[]rune` ：

``` go
package main

import "fmt"

func main() {
    text := "Go語言"
    cs := []rune(text)
    fmt.Printf("%c\n", cs[2]) // 語
    fmt.Println(len(cs))      // 4
}
```

字串型態轉換為 `[]rune` 時，會將 UTF-8 編碼的位元組，轉換為 Unicode 碼點，在這個例子中可以看到，`cs[2]` 確實地取得了第三個文字「語」，而 `len` 也確實取得數量 4。

如〈[認識預定義型態](http://openhome.cc/Gossip/Go/PreDeclaredType.html)〉中談過的，在 Go 中並沒有字元對應的型態，只有碼點的概念，`rune` 為 `int32` 的別名，可用來儲存 Unicode 碼點（code point），如果使用 `fmt.Printf("%d\n", cs[2])`，會顯示 35486，這就是「語」的 Unicode 碼點，35486 的 16 進位表示是 8a9e，因此，如果你寫 `'\u8a9e'`，也會得到一個 `rune` 代表著「語」，`fmt.Printf("%c", '\u8a9e')` 也會顯示「語」，當然，直接寫 `'語'` 也是可以得到一個 `rune`。

想從 `rune` 得到一個 `string`，可以直接寫 `string('語')` 就可以了。如果想以 `rune` 為單位來走訪字串，而不是以位元組走訪，可以使用 `for range`，例如：

``` go
package main

import "fmt"

func main() {
    text := "Go語言"
    for index, runeValue := range text {
        fmt.Printf("%#U 位元起始位置 %d\n", runeValue, index)
    }
}
```

可以看到，`for range` 可以同時取得每個 `rune` 在字串中的位元起始位置，以及 `rune` 值，`%U` 可以用 16 進位顯示 `rune`，如果是 `%#U`，還會一併顯示碼點的列印形式。

這個程式的執行結果會顯示：

``` go
U+0047 'G' 位元起始位置 0
U+006F 'o' 位元起始位置 1
U+8A9E '語' 位元起始位置 2
U+8A00 '言' 位元起始位置 5
```

總而言之，Go 的字串是由 UTF-8 編碼的位元組構成，在〈[Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)〉談到了這麼設計的理由是，「字元」的定義太模稜兩可了，Go 為了避免模稜兩可，就將字串定義為 UTF-8 編碼的位元組構成，而 `rune` 用於儲存碼點。

PS. 這大概也是為何，我會整理出〈[亂碼 1/2](http://openhome.cc/Gossip/Encoding/)〉的原因 … XD

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
