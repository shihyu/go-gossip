<div id="main" role="main" style="height: auto !important;">

<div class="header">

# ring 套件

</div>

  

對於環狀資料結構，Go 提供了 `container/ring` 套件，`Ring` 結構有 `Value` 欄位，可以使用 `New` 指定元素數量來建立實例，可用的方法有：

``` go
func (r *Ring) Do(f func(interface{}))  // 走訪每個元素並傳入 f
func (r *Ring) Len() int                // 元素數量
func (r *Ring) Link(s *Ring) *Ring      // 銜接另一個 Ring
func (r *Ring) Move(n int) *Ring        // 移動 n 個元素，n 可正或負
func (r *Ring) Next() *Ring             // 下一個鏈（也就是下一個元素）
func (r *Ring) Prev() *Ring             // 上一個鏈（也就是上一個元素）
func (r *Ring) Unlink(n int) *Ring      // 解除指定數量的 Ring，傳回被解除的子鏈
```

因為是環狀結構，每個元素都可視為一個鏈的開頭或結尾，因此 `Link` 等操作都傳回 `*Ring`。底下是個建立 `Ring` 並設值的簡單範例：

``` go
package main

import (
    "fmt"
    "container/ring"
)

func main() {
    numbers := ring.New(10)
    for i := 0; i < numbers.Len(); i++ {
        numbers.Value = i
        numbers = numbers.Next()
    }

    numbers.Do(func(n interface{}) {
        fmt.Printf("%d ", n.(int))
    })
}
```

[`ring` 的官方文件](https://pkg.go.dev/container/ring/)有相關方法的範例，這邊就不重複列出了，實際應用上，`ring` 可以用來管理有限筆數的歷史記錄、輪播等。

這邊的話拿來解一下 [約瑟夫問題（Josephus Problem）](https://openhome.cc/Gossip/AlgorithmGossip/) 好了：

``` go
package main

import (
    "fmt"
    "container/ring"
)

type Person struct {
    Number int
}

func main() {
    persons := ring.New(41)
    // 給每個人編號
    for i := 1; i <= persons.Len(); i++ {
        persons.Value = &Person{i}
        persons = persons.Next()    
    }

    persons = persons.Prev()

    // 最後只留下兩人
    for persons.Len() > 2 {
        for i := 1; i <= 2; i++ {
            persons = persons.Next()
        }
        // 報數 3 Out
        persons.Unlink(1)
    }

    fmt.Print("安全位置：")
    persons.Do(func(p interface{}) {
        person := p.(*Person)
        fmt.Printf("%d ", person.Number)
    })
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
