<div id="main" role="main" style="height: auto !important;">

<div class="header">

# heap 套件

</div>

  

如果在收集元素的過程中，想要一併排序，方式之一是使用[堆積排序](https://openhome.cc/Gossip/AlgorithmGossip/HeapSort.htm)，對於這個需求，Go 提供了 `heap` 套件作為實現上的輔助。

`heap` 套件提供的是最小堆積樹演算，底層的資料結構必須實現 `heap.Interface`：

``` go
type Interface interface {
    sort.Interface
    Push(x interface{}) 
    Pop() interface{} 
}
```

也就是說，除了實現 `sort.Interface` 的 `Len`、`Less`、`Swap` 方法之外，還要實現 `Push` 與 `Pop` 的行為，在 [`heap` 的 Go 官方文件說明](https://pkg.go.dev/container/heap/) 有個簡單範例：

``` go
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

實現了 `heap.Interface` 的資料結構，就可以透過 `heap` 套件中的 `Init`、`Push`、`Pop` 等函式來進行操作：

``` go
h := &IntHeap{2, 1, 5}
heap.Init(h)
heap.Push(h, 3)
fmt.Printf("minimum: %d\n", (*h)[0])
for h.Len() > 0 {
    fmt.Printf("%d ", heap.Pop(h))
}
```

在 `Push`、`Pop` 過程中有關堆積樹的調整，就都由 `heap.Push`、`heap.Pop` 等函式來處理了。

官方文件提供的範例是可以簡單示範 `heap` 套件的使用，不過，一下子使用 `heap.Xxx`，一下子又是使用 `h.Xxx` 的混合風格，看來蠻怪的，可以來改變一下：

``` go
package main

import (
    "container/heap"
    "fmt"
)

// IntSlice 實現了 heap.Interface
type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *IntSlice) Push(x interface{}) {
    *s = append(*s, x.(int))
}

func (s *IntSlice) Pop() interface{} {
    old := *s
    n := len(old)
    x := old[n-1]
    *s = old[0 : n-1]
    return x
}

// IntHeap 封裝了 IntSlice
type IntHeap struct {
    elems IntSlice
}

// 實現相關函式或方法時，透過 heap 提供的函式
func NewIntHeap(numbers ...int) *IntHeap {
    h := &IntHeap{IntSlice(numbers)}
    heap.Init(&(h.elems))
    return h
}

func (h *IntHeap) Push(n int) {
    heap.Push(&(h.elems), n)
}

func (h *IntHeap) Pop() int {
    return heap.Pop(&(h.elems)).(int)
}

func (h *IntHeap) Len() int {
    return len(h.elems)
}

// 一律透過 h 來操作
func main() {
    h := NewIntHeap(2, 1, 5)
    h.Push(3)
    for h.Len() > 0 {
        fmt.Printf("%d ", h.Pop())
    }
}
```

官方文件提供的範例中，還有個 `PriorityQueue` 的實現，類似地，該範例是簡單示範，混合了兩種操作風格，你也可以試著自行把 `heap.Xxx` 的操作給封裝起來。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
