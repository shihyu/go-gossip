<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 運算子

</div>

  

Go 語言中的運算子，大致上與 C 語系的語言中提供的運算子差不多，其中 `&`、`*` 也用來作為指標（Pointer）運算子。

# 算術運算子

算術運算子作用於數值，產生與第一個運算元相同型態的結果。`+`、`-`、`*`、`/` 四個運算子，可用於整數、浮點數與複數；`+` 也用於字串串接；`%` 餘除運算子，只用於整數，`&`、`|`、`^`、`&^` 位元運算子只用於整數，`<<`、`>>` 位移運算子只用於整數。

`+`、`-`、`*`、`/` 使用上應該沒什麼問題，主要就是注意運算的順序是先乘除後加減，必要時使用括號讓順序清楚，例如：

``` prettyprint
package main

import "fmt"

func main() {
    fmt.Println(1 + 2*3)         // 7
    fmt.Println(2 + 2 + 8/4)     // 6
    fmt.Println((2 + 2 + 8) / 4) // 3
    fmt.Println(10 % 3)          // 1
}
```

`%`運算子計算的結果是除法後的餘數，例如上頭 `10 % 3` 會得到餘數 1。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

對於遞增與遞減 1 的操作，Go 可以使用 `++` 與 `--` 的操作，不過，`++` 與 `--` 只能置於變數後方，而且是個陳述，因此，對於 `i := 1`，你可以在一行陳述中寫 `i++` 或 `i--`，不過，不能寫 `fmt.Println(i++)`，這樣就能避免是要先傳回 `i` 值再遞增 `i`，還是先遞增 `i` 再傳回 1 的問題。

在二進位運算上有 AND、OR、XOR 等運算，底下是 Go 中的一些例子：

``` prettyprint
package main

import "fmt"

func main() {
    fmt.Println("AND運算：")
    fmt.Printf("0 AND 0 %5d\n", 0&1)
    fmt.Printf("0 AND 1 %5d\n", 0&1)
    fmt.Printf("1 AND 0 %5d\n", 1&0)
    fmt.Printf("1 AND 1 %5d\n", 1&1)

    fmt.Println("\nOR運算：")
    fmt.Printf("0 OR 0 %6d\n", 0|0)
    fmt.Printf("0 OR 1 %6d\n", 0|1)
    fmt.Printf("1 OR 0 %6d\n", 1|0)
    fmt.Printf("1 OR 1 %6d\n", 1|1)

    fmt.Println("\nXOR運算：")
    fmt.Printf("0 XOR 0 %5d\n", 0^0)
    fmt.Printf("0 XOR 1 %5d\n", 0^1)
    fmt.Printf("1 XOR 0 %5d\n", 1^0)
    fmt.Printf("1 XOR 1 %5d\n", 1^1)

    fmt.Println("\nAND NOT運算：")
    fmt.Printf("0 AND NOT 0 %5d\n", 0&^0)
    fmt.Printf("0 AND NOT 1 %5d\n", 0&^1)
    fmt.Printf("1 AND NOT 0 %5d\n", 1&^0)
    fmt.Printf("1 AND NOT 1 %5d\n", 1&^1)
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

執行結果如下：

``` prettyprint
AND運算：
0 AND 0     0
0 AND 1     0
1 AND 0     0
1 AND 1     1

OR運算：
0 OR 0      0
0 OR 1      1
1 OR 0      1
1 OR 1      1

XOR運算：
0 XOR 0     0
0 XOR 1     1
1 XOR 0     1
1 XOR 1     0

AND NOT運算：
0 AND NOT 0     0
0 AND NOT 1     0
1 AND NOT 0     1
1 AND NOT 1     0
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

位元運算是逐位元運算，例如 10010001 與 01000001 作 AND 運算，是一個一個位元對應運算，答案就是 00000001。補數運算是將所有位元 0 變 1，1 變 0。例如 00000001 經補數運算就會變為 11111110。Go 的補數運算子是 `^`，例如：

``` prettyprint
package main

import "fmt"

func main() {
    number := 0
    fmt.Println(^number)  // -1
}
```

上面的程式片段會顯示 -1，因為 number 在記憶體中全部位元都是 0，經補數運算全部位元就都變成 1，這個數在電腦中用整數表示則是 -1。

`<<` 左移運算子會將所有位元往左移指定位數，左邊被擠出去的位元會被丟棄，而右邊補上 0；`>>` 右移運算則是相反，會將所有位元往右移指定位數，右邊被擠出去的位元會被丟棄，至於最左邊補上原來的位元，如果左邊原來是 0 就補0，1 就補 1。

``` prettyprint
package main

import "fmt"

func main() {
    number := 1
    fmt.Printf("2 的 0 次方: %d\n", number)        // 1
    fmt.Printf("2 的 1 次方: %d\n", number << 1)   // 2
    fmt.Printf("2 的 2 次方: %d\n", number << 2)   // 4
    fmt.Printf("2 的 3 次方: %d\n", number << 3)   // 8
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

實際來左移看看就知道為何可以如此作次方運算了：

``` prettyprint
00000001 -> 1 
00000010 -> 2 
00000100 -> 4 
00001000 -> 8
```

對於一個算術運算 `x = x op y`，可以寫成 `x op= y`，`op` 是指算術運算子，例如 `x = x + y`，可以寫成 `x += y`，這也就是所謂的指定運算子。

# 比較運算

數學上有大於、等於、小於的比較運算，Go 中也提供了這些運算子，它們有大於（`>`）、不小於（`>=`）、小於（`<`）、不大於（`<=`）、等於（`==`）以及不等於（`!=`），比較條件成立時用 `true` 表示，比較條件不成立用 `false` 表示。以下程式片段示範了幾個比較運算的使用：

``` prettyprint
package main

import "fmt"

func main() {
    fmt.Printf("10 >  5 結果 %t\n", 10 > 5)   // true
    fmt.Printf("10 >= 5 結果 %t\n", 10 >= 5)  // true
    fmt.Printf("10 <  5 結果 %t\n", 10 < 5)   // false
    fmt.Printf("10 <= 5 結果 %t\n", 10 <= 5)  // false
    fmt.Printf("10 == 5 結果 %t\n", 10 == 5)  // false
    fmt.Printf("10 != 5 結果 %t\n", 10 != 5)  // true
}
```

`==` 與 `!=` 只能用在 comparable 的運算元上，這有一套嚴格規則，Go 語言中哪些值是可以比較的，可以參考規格書中〈[Comparison operators](https://go.dev/ref/spec#Comparison_operators)〉的說明。

Go 中沒有 `?:` 三元條件運算子。

# 邏輯運算

在邏輯上有所謂的「且」、「或」與「反相」，在 Go 中提供對應的邏輯運算子（Logical operator），分別為 `&&`、`||`及 `!`。看看以下的程式片段會輸出什麼結果？

``` prettyprint
package main

import "fmt"

func main() {
    number := 75
    fmt.Println(number > 70 && number < 80)     // true
    fmt.Println(number > 80 || number < 75)     // false
    fmt.Println(!(number > 80 || number < 75))  // true
}
```

`&&` 與 `||` 有捷徑運算（Short-Circuit Evaluation）。因為 `&&` 只要其中一個為假，就可以判定結果為假，所以只要左運算元評估為 `false`，就會直接傳回 `false`，不會再去運算右運算元。因為 `||` 只要其中一個為真，就可以判定結果為真，所以只要左運算元評估為 `true`，就會直接傳回 `true`，就不會再去運算右運算元。

來舉個運用捷徑運算的例子，在 Go 中兩個整數相除，若除數為 0 會發生 integer divide by zero 的錯誤，以下運用 `&&` 捷徑運算避免了這個問題：

``` prettyprint
if(number2 != 0 && number1 / number2 > 1) {
    fmt.Println(number1 / number2)
}
```

在這個程式片段中，變數 number1 與 number2 都是 `int` 型態，如果 `number2` 為 0 的話，`&&` 左邊運算元結果就是 `false`，直接判斷整個 `&&` 的結果應是 `false`，不用再去評估右運算元，從而避免了 `number1 / number2` 而 `number2` 等於 `0` 時的除零錯誤。

# 指標

Go 語言中有指標（Pointer），你可以在宣告變數時於型態前加上 `*`，這表示建立一個指標，例如：

``` prettyprint
var i *int
```

這時 `i` 是個空指標，也就是值為 `nil`，上頭等同於 `var i *int = nil`，目前並沒有儲存任何位址，如果想讓它儲存另一個變數的記憶體位址，可以使用 `&` 取得變數位址並指定給 `i`，例如：

``` prettyprint
package main

import "fmt"

func main() {
    var i *int
    j := 1

    i = &j
    fmt.Println(i)  // 0x104382e0 之類的值
    fmt.Println(*i) // 1

    j = 10
    fmt.Println(*i) // 10

    *i = 20
    fmt.Println(j) // 20
}
```

`j` 的位置儲存了 1，那麼具體來說，`j` 的位置到底是在哪？這就是 `&` 取址運算的目的，`&j` 具體取得了 `j` 的位置，然後指定給 `i`。

如上所示，如果想存取指標位址處的變數儲存的值，可以使用 `*`，因而，你改變 `j` 的值，`*i` 取得的就是改變後的值，透過 `*i` 改變值，從 `j` 取得的也會是改變後的值。

其應用的實例之一是使用 `fmt.Scanf` 取得標準輸入時，例如：

``` prettyprint
package main

import "fmt"

func main() {
    var input int
    fmt.Printf("輸入數字")
    fmt.Scanf("%d", &input)
    fmt.Println(input)
}
```

這邊使用 `&input` 取出 `input` 的記憶體位址值，並傳入 `fmt.Scanf` 函式，函式中會取得使用者的標準輸入，並儲存至 `input` 變數的記憶體位址，因而，再度取得 `input` 的值時，就會是使用者輸入的值。

Go 雖然有指標，不過不能如同 C/C++ 那樣對指標做運算，之後有機會用到指標時，會再做相關說明。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
