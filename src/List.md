<div id="main" role="main" style="height: auto !important;">

<div class="header">

# list 套件

</div>

  

如果想連續地看待一組資料，可以使用 slice，優點是可以透過索引快速存取，透過 `append` 也可以附加元素，若偶而需要安插、刪除元素，可以透過切片等操作來實現。

然而，如果經常性地需要安插、刪除元素，透過 slice 實現缺乏效率時，Go 提供了 `container/list` 套件，可讓開發者基於雙向鏈結的 `list.List` 實作來達成需求。

想要建立 `list.List` 實例，可以透過 `list.New`，實例可使用的方法有：

``` prettyprint
func (l *List) Back() *Element
func (l *List) Front() *Element
func (l *List) Init() *List
func (l *List) InsertAfter(v interface{}, mark *Element) *Element
func (l *List) InsertBefore(v interface{}, mark *Element) *Element
func (l *List) Len() int
func (l *List) MoveAfter(e, mark *Element)
func (l *List) MoveBefore(e, mark *Element)
func (l *List) MoveToBack(e *Element)
func (l *List) MoveToFront(e *Element)
func (l *List) PushBack(v interface{}) *Element
func (l *List) PushBackList(other *List)
func (l *List) PushFront(v interface{}) *Element
func (l *List) PushFrontList(other *List)
func (l *List) Remove(e *Element) interface{}
```

從 `PushBack`、`PushFront` 方法的參數型態 `interface{}` 就能知道，`list.List` 可以保存任意型態的資料，它們會傳回 `*Element`，`Element` 是個結構，公開的欄位有 `Value`，公開的方法為 `Next` 與 `Prev`：

``` prettyprint
type Element struct {
    Value interface{}
}

func (e *Element) Next() *Element

func (e *Element) Prev() *Element
```

因此，若你保留傳回的 `*Element`，可以透過 `Value` 取得放入 `list.List` 的值，必要時也可以透過 `Next` 或 `Prev` 方法，往後探尋下一元素或往前探尋前一元素，`Next` 與 `Prev` 方法傳回的也是 `*Element`，因此隨時可以往前探尋元素前或後全部的清單。

`Back`、`Front` 方法，分別傳回 `list.List` 最後、最前一個元素，因此，若要從清單頭走訪至尾，基本的模式就是：

``` prettyprint
package main

import (
    "fmt"
    "container/list"
)

func printAll(lt *list.List) {
    for e := lt.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}

func main() {
    lt := list.New()
    for i := 1; i <= 10; i++ {
        lt.PushBack(i)
    }

    printAll(lt)
}
```

你可能會有問題，`Element` 的 `Value` 型態是 `interface{}`，那麼想操作保存的元素值上的欄位、方法時，不就要知道型態嗎？這裡仍然需要透過型態斷言：

補充（Go 1.18+）：Go 已支援泛型，不過 `container/list` 本身仍是舊有 API 風格（實務上可視為 `any` / `interface{}` 容器），因此若你需要型別安全的清單結構，常見做法是自行包一層泛型型別。

``` prettyprint
package main

import (
    "fmt"
    "container/list"
)

type Person struct {
    Name string
    Age  int
}

func printAllPerson(persons *list.List) {
    for e := persons.Front(); e != nil; e = e.Next() {
        p := e.Value.(*Person)
        fmt.Printf("姓名：%s\t年齡：%d\n", p.Name, p.Age)
    }
}

func main() {
    persons := list.New()

    persons.PushBack(&Person{"Irene", 12})
    persons.PushBack(&Person{"Justin", 45})
    persons.PushBack(&Person{"Monica", 42})

    printAllPerson(persons)
}
```

你可能還會有其他問題，例如 `list.List` 怎麼不支援索引？要怎麼進行排序等？…唔…`list.List` 提供的方法怎麼這麼少？

嚴格來說，不會直接使用 `list.List` 來保存資料，而是如果某資料結構底層需要雙向鏈結的特性，可以透過 `list.List` 來實現。例如，實現一個 `PersonQueue`：

``` prettyprint
package main

import (
    "fmt"
    "container/list"
)

type Person struct {
    Name string
    Age  int
}

type PersonQueue struct {
    list *list.List
}

func NewPersonQueue() *PersonQueue {
    return &PersonQueue{list.New()}
}

func (q *PersonQueue) Len() int {
    return q.list.Len()
}

func (q *PersonQueue) Offer(p *Person) {
    q.list.PushBack(p)
}

func (q *PersonQueue) Peek() *Person {
    if q.list.Len() == 0 {
        return nil
    }

    e := q.list.Remove(q.list.Front())
    return e.(*Person)
}

func main() {
    q := NewPersonQueue()

    q.Offer(&Person{"Irene", 12})
    q.Offer(&Person{"Justin", 45})
    q.Offer(&Person{"Monica", 42})

    for p := q.Peek(); p != nil; p = q.Peek() {
        fmt.Printf("姓名：%s\t年齡：%d\n", p.Name, p.Age)
    }
}
```

因此，並不是 `list.List` 不常用，而是你可能很少自行實現資料結構（都拿別人寫好的來用？）；另一種說法「每當想使用 `list.List` 時，都該思考一下是否優先使用 slice。」的說法也不是完全正確…

若想使用 `list.List`，應該問的是，你的資料結構在實現上需要雙向鏈結的特性嗎？例如，也許你會需要有個具索引的資料結構，同時底層實現必須是雙向鏈結（像是 Java 的 `LinkedList`）？那麼就可以考慮透過 `list.List` 來實現。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
