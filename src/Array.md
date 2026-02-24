<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 身為複合值的陣列

</div>

  

在 Go 中，陣列的長度固定，是個複合值，元素的型態及個數決定了陣列的型態，在記憶體中使用連續空間配置。

# 建立與存取陣列

建立陣列的方式是 `[n]type`，其中 `n` 為陣列的元素數量，`type` 是元素的型態。例如：

``` prettyprint
package main

import "fmt"

func main() {
    var scores [10]int
    scores[0] = 90
    scores[1] = 89
    fmt.Println(scores)      // [90 89 0 0 0 0 0 0 0 0]
    fmt.Println(len(scores)) // 10
}
```

在上面的程式中，建立了具有 10 個元素的陣列，可以用來儲存 `int` 型態的值，可透過 `scores` 變數指定索引來存取元素，`scores` 變數的型態為 `[10]int`，記得，長度也是陣列的型態的一部份，若一個陣列為 `[10]int`，而另一個陣列為 `[5]int`，這兩個陣列會是不同的型態，像上頭這樣宣告陣列，預設每個元素都會初始為 0。

陣列使用索引存取，如同其他語言的慣例，索引從 0 開始，`len` 函式可以取得陣列的長度，如果想在建立陣列時指定初始，可以如下：

``` prettyprint
package main

import "fmt"

func main() {
    arr1 := [3]int{1, 2, 3}
    arr2 := [5]int{1, 2, 3}
    arr3 := [...]int{1, 2, 3, 4, 5}
    fmt.Println(arr1) // [1 2 3]
    fmt.Println(arr2) // [1 2 3 0 0]
    fmt.Println(arr3) // [1 2 3 4 5]
}
```

在上頭可以看到，如果宣告的元素數量不足 `[]` 中指定的數量，那麼會自動給予初值，也可以使用 `...`，或者只寫 `[]`，讓編譯器自動判斷數量，如果宣告的元素數量超過 `[]` 中指定的數量，那麼會有 out of bounds 的編譯錯誤。

# 陣列指定與比較

在 Go 中，陣列指定會逐一複製值，例如：

``` prettyprint
package main

import "fmt"

func main() {
    arr1 := [...]int{1, 2, 3}
    arr2 := arr1
    fmt.Println(arr1) // [1 2 3]
    fmt.Println(arr2) // [1 2 3]
    arr1[0] = 10
    fmt.Println(arr1) // [10 2 3]
    fmt.Println(arr2) // [1 2 3]
}
```

在呼叫函式時若傳遞陣列給參數，或者是傳回陣列，也是做複製的動作。陣列可以使用 `==` 與 `!=` 進行比較，由於長度也是陣列型態的一部份，因此，只要長度與元素型態相同的陣列才可以做比較，如果將 `[3]int` 與 `[5]int` 做比較，會發生 mismatched types 編譯錯誤，同樣的，指定陣列給另一陣列時，也必須是相同型態的陣列。

# 巢狀陣列

Go 的陣列是線性的，如果想模擬多維，可以使用巢狀陣列。例如，建立一個二維陣列：

``` prettyprint
package main

import "fmt"

func main() {
    var arr [2][3]int
    fmt.Println(arr)   // [[0 0 0] [0 0 0]]
}
```

顯然地，第一個 `[]` 中數字指定了陣列中會有兩個 `[3]int` 陣列，因此，若要同時宣告陣列中的元素，可以如下：

``` prettyprint
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
```

上頭一口氣示範了幾種巢狀陣列的宣告方式，基本上後兩種應該是比較容易撰寫的，由於陣列的長度是型態的一部份，必須在宣告時指定，因此，就二維陣列來說，一定都是方陣。

# 走訪陣列

想要逐一走訪陣列的話，基本上可以使用 `for` 迴圈，例如：

``` prettyprint
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3}
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%d\n", arr[i])
    }
}
```

另一個方式是使用 `for range`：

``` prettyprint
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3}
    for index, element := range arr {
        fmt.Printf("%d: %d\n", index, element)
    }
}
```

在不需要索引的情況下，可以使用 `_` 忽略傳回的索引值，例如：

``` prettyprint
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3}
    for _, element := range arr {
        fmt.Printf("%d\n", element)
    }
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
