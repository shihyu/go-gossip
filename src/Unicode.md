<div id="main" role="main" style="height: auto !important;">

<div class="header">

# unicode 套件

</div>

  

`unicode`、`unicode/utf8`、`unicode/utf16` 是用來判斷、處理 Unicode 以及 UTF-8、UTF-16 編碼的套件，在使用這些套件之前，要先知道的是，Go 認為「字元」的定義過於模糊，在 Go 中使用 `rune` 儲存 Unicode 碼點（Code point），而 Go 中字串是 UTF-8 編碼的位元組組成。

[`unicode`](https://pkg.go.dev/unicode/) 套件主要用來判斷 Unicode 碼點的特性（properties），在 Unicode 規範中，每個碼點會被指定某些特性，具有相同特性的一組碼點構成一個集合，以便於理解、判斷這組碼點。

例如，[General Category](https://en.wikipedia.org/wiki/Template:General_Category_(Unicode)) 特性有 Letter/L 代表字母、Number/N 代表數字等，在 Go 的 [unicode 套件文件的 Variables](https://pkg.go.dev/unicode/#pkg-variables) 一開頭，列出的就是這類特性的變數：

``` go
var (
    ...
    Digit  = _Nd // 十進位數字的集合

    Letter = _L  // 字母集合
    L      = _L
    ...
    Number = _N  // 數字集合
    N      = _N
    ...
}
```

每個變數的型態都是 `*RangeTable`，由碼點的範圍等欄位組成：

``` go
type RangeTable struct {
    R16         []Range16   // 用 uint16 記錄碼點低位至高位
    R32         []Range32   // 記錄 R16 無法表示的範圍，用 uint32 記錄碼點低位至高位
    LatinOffset int 
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

碼點範圍表可以在 [tables.go](https://go.dev/src/unicode/tables.go) 找到。舉例來說，字母集合的碼點範圍：

``` go
var _L = &RangeTable{
    R16: []Range16{
        {0x0041, 0x005a, 1},
        {0x0061, 0x007a, 1},
        {0x00aa, 0x00b5, 11},
        很長的清單...
```

透過指定 `RangeTable`，就可以簡單地判斷碼點是否有某特性，例如，`²³¹¼½¾𝟏𝟐𝟑𝟜𝟝𝟞𝟩𝟪𝟫𝟬𝟭𝟮𝟯𝟺𝟻𝟼㉛㉜㉝ⅠⅡⅢⅣⅤⅥⅦⅧⅨⅩⅪⅫⅬⅭⅮⅯⅰⅱⅲⅳⅴⅵⅶⅷⅸⅹⅺⅻⅼⅽⅾⅿ` 都是數字：

``` go
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
```

Unicode 將希臘文、漢字等以[文字（Script）](https://en.wikipedia.org/wiki/Script_(Unicode))特性標示，在 Go 的 [unicode 套件文件的 Variables](https://pkg.go.dev/unicode/#pkg-variables) 第二組列出的變數清單，就是對應的 `RangeTable`，例如 `unicode.Han` 是正體中文、簡體中文，以及日、韓、越南文的全部漢字範圍。

另外還有一些其他特性，列在 Go 的 [unicode 套件文件的 Variables](https://pkg.go.dev/unicode/#pkg-variables) 第三組變數清單，例如 `unicode.White_Space` 代表被標示為空白特性的碼點，這包括了半形、全形、Tab 等。

如果想要使用多個 RangeTable，`可以透過 IsOneOf`：

``` go
func IsOneOf(ranges []*RangeTable, r rune) bool
```

`unicode` 也提供了一些常用的判斷函式：

``` go
func IsControl(r rune) bool
func IsDigit(r rune) bool
func IsGraphic(r rune) bool
func IsLetter(r rune) bool
func IsLower(r rune) bool
func IsMark(r rune) bool
func IsNumber(r rune) bool
func IsPrint(r rune) bool
func IsPunct(r rune) bool
func IsSpace(r rune) bool
func IsSymbol(r rune) bool
func IsTitle(r rune) bool
func IsUpper(r rune) bool
```

在大小寫或特定轉換上，有以下的函式：

``` go
func To(_case int, r rune) rune
func ToLower(r rune) rune
func ToTitle(r rune) rune
func ToUpper(r rune) rune
```

基本上，這可以應付大多數語言的轉換，像是全形字母的大小寫或首字母大寫等，`To` 可使用的常數有：

``` go
const (
    UpperCase = iota
    LowerCase
    TitleCase
    MaxCase
)
```

例如，`unicode.To(unicode.UpperCase, rune('ａ'))` 可以得到 `'Ａ'`。

# unicode/utf8、unicode/utf16 套件

`unicode/utf8` 套件提供的函式，主要是進行 `rune` 與 UTF-8 編碼之間的處理。例如驗證是否為合法的 UTF-8 `[]byte` 或字串：

``` go
func Valid(p []byte) bool
func ValidString(s string) bool
```

驗證 `rune` 可否編碼為 UTF-8：

``` go
func ValidRune(r rune) bool
```

在 `rune` 與 UTF-8 編碼之間轉換：

``` go
func DecodeLastRune(p []byte) (r rune, size int)
func DecodeLastRuneInString(s string) (r rune, size int)
func DecodeRune(p []byte) (r rune, size int)
func DecodeRuneInString(s string) (r rune, size int)
func EncodeRune(p []byte, r rune) int
```

`unicode/utf16` 主要是進行 `rune` 與 UTF-16 編碼之間的處理，只不過目前函式只有幾個：

``` go
func Decode(s []uint16) []rune
func DecodeRune(r1, r2 rune) rune
func Encode(s []rune) []uint16
func EncodeRune(r rune) (r1, r2 rune)
func IsSurrogate(r rune) bool
```

UTF-8 編碼下，碼元（code unit）是 8 個位元，Go 中使用 `byte` 也就是 `uint8` 來儲存，UTF-16 編碼下，碼元（code unit）是 16 個位元，Go 中使用 `uint16` 來儲存。

來看個簡單的範例，使用 `unicode/utf8` 與 `unicode/utf16` 套件來顯示「Hello, 世界」的 UTF-16 碼元：

``` go
package main

import (
    "fmt"
    "unicode/utf8"
    "unicode/utf16"
)

func main() {    
    b := []byte("Hello, 世界")

    for len(b) > 0 {
        r, size := utf8.DecodeRune(b)
        u16 := utf16.Encode([]rune{r})
        fmt.Printf("%#U:\n  Code unit %04X\n", r, u16)
        b = b[size:]
    }
}
```

顯示結果如下：

``` go
U+0048 'H':       
  Code unit [0048]
U+0065 'e':       
  Code unit [0065]
U+006C 'l':       
  Code unit [006C]
U+006C 'l':       
  Code unit [006C]
U+006F 'o':
  Code unit [006F]
U+002C ',':
  Code unit [002C]
U+0020 ' ':
  Code unit [0020]
U+4E16 '世':
  Code unit [4E16]
U+754C '界':
  Code unit [754C]
```

Unicode 碼點號碼與碼元顯示剛好一樣對吧？這就是為什麼常有人會亂說「Unicode 使用 16 位元儲存」的原因之一吧！… XD

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
