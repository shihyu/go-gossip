<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 結構入門

</div>

  

有些資料會有相關性，例如，一個 XY 平面上的點可以使用 (x, y) 座標來表示；名稱、郵件位址、電話可能代表著一張名片上的資訊。將相關聯的資料組織在一起，對於資料本身的可用性或者是程式碼的可讀性，都會有所幫助。

# struct 組織資料

Go 語言中有 `struct`，可以用來將相關的資料組織在一起，如果你學過 C 語言，這對你應該不陌生。舉個例子來說，相對於個別地存取 `x`、`y` 變數：

``` prettyprint
package main

import "fmt"

func main() {
    x := 10
    y := 20
    fmt.Printf("{%d %d}\n", x, y) // {10 20}

    x = 20
    y = 30
    fmt.Printf("{%d %d}\n", x, y) // {20 30}
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

若 `x` 與 `y` 變數，相當於 XY 平面上的 (x, y) 座標，那麼將之組織在一起同時存取會比較好：

``` prettyprint
package main

import "fmt"

func main() {
    point := struct{ x, y int }{10, 20}
    fmt.Printf("{%d %d}\n", point.x, point.y) // {10 20}

    point.x = 20
    point.y = 30

    fmt.Printf("{%d %d}\n", point.x, point.y) // {20 30}
}
```

實際上，`fmt.Println` 可以直接處理 `struct`，因此，上面的例子，可以直接使用 `fmt.Println(point)` 來得到相同的顯示結果。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

在上面的例子中，`struct` 定義了一個結構，當中包括了 `x` 與 `y` 兩個值域（field），接著馬上用它來建立了一個實例，依順序指定了 `x` 與 `y` 的值是 `10` 與 `20`，可以看到，想要存取結構的值域，可以運過點運算子（`.`）。

# 基於結構定義新型態

上面的例子中，建立了一個匿名型態的結構，你可以使用 `type` 基於 `struct` 來定義新型態，例如：

``` prettyprint
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    point1 := Point{10, 20}
    fmt.Println(point1) // {10 20}

    point2 := Point{Y: 20, X: 30}
    fmt.Println(point2) // {30 20}
} 
```

在上面基於結構定義了新型態 `Point`，留意到名稱開頭的大小寫，若是大寫的話，就可以在其他套件中存取，這點對於結構的值域也是成立，大寫名稱的值域，才可以在其他套件中存取。在範例中也可以看到，建立並指定結構的值域時，可以直接指定值域名稱，而不一定要按照定義時的順序。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如果一開始不知道結構的值域數值為何，可以使用 `var` 宣告即可，那麼值域會依型態而有適當的預設值。例如：

``` prettyprint
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {  
    var point Point
    fmt.Println(point)      // {0 0}           
}
```

`point` 並不是參考，`point` 的位置開始，有一片可以儲存結構的空間，可以使用 `&` 來取得 `point` 的位址值，`point` 的位址值無法改變。

# 結構與指標

如果你建立了一個結構的實例，並將之指定給另一個結構變數，那麼會進行值域的複製。例如：

``` prettyprint
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {  
    point1 := Point{X: 10, Y: 20}
    point2 := point1

    point1.X = 20

    fmt.Println(point1)  // {20, 20}
    fmt.Println(point2)  // {10 20}
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

這對於函式的參數傳遞也是一樣的：

``` prettyprint
package main

import "fmt"

type Point struct {
    X, Y int
}

func changeX(point Point) {
    point.X = 20
    fmt.Println(point)
}

func main() {
    point := Point{X: 10, Y: 20}

    changeX(point)     // {20 20}
    fmt.Println(point) // {10 20}
}
```

`point` 的位置開始儲存了結構，可以對 `point` 使用 `&` 取值，將位址值指定給指標，因此若指定或傳遞結構時，不是想要複製值域，可以使用指標。例如：

``` prettyprint
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    point1 := Point{X: 10, Y: 20}
    point2 := &point1

    point1.X = 20

    fmt.Println(point1) // {20, 20}
    fmt.Println(point2) // &{20 20}
}
```

注意到 `point2 := &point1` 多了個 `&`，這取得了 `point1` 實例的指標值，並傳遞給 `point2`，`point2` 的型態是 `*Point`，也就是相當於 `var point2 *Point = &point1`，因此，當你透過 `point1.X` 改變了值，透過 `point2` 就能取得對應的改變。

類似地，也可以在傳遞參數給函式時使用指標：

``` prettyprint
package main

import "fmt"

type Point struct {
    X, Y int
}

func changeX(point *Point) {
    point.X = 20
    fmt.Printf("&{%d %d}\n", point.X, point.Y)
}

func main() {
    point := Point{X: 10, Y: 20}

    changeX(&point)    // &{20 20}
    fmt.Println(point) // {20 20}
}
```

可以看到在 Go 語言中，即使是指標，也可以直接透過點運算子來存取值域，這是 Go 提供的語法糖，`point.X` 在編譯過後，會被轉換為 `(*point).X`。

你也可以透過 `new` 來建立結構實例，這會傳回結構實例的位址：

``` prettyprint
package main

import "fmt"

type Point struct {
    X, Y int
}

func default_point() *Point {
    point := new(Point)
    point.X = 10
    point.Y = 10
    return point
}

func main() {
    point := default_point()
    fmt.Println(point) // &{10 10}
}
```

在這邊，`point` 是個指標，也就是 `*Point` 型態，儲存了結構實例的位址。

結構的值域也可以是指標型態，也可以是結構自身型態之指標，因此可實現鏈狀參考，例如：

``` prettyprint
package main

import "fmt"

type Point struct {
    X, Y int
}

type Node struct {
    point *Point
    next  *Node
}

func main() {
    node := new(Node)

    node.point = &Point{10, 20}
    node.next = new(Node)

    node.next.point = &Point{10, 30}

    fmt.Println(node.point)      // &{10 20}
    fmt.Println(node.next.point) // &{10 30}
}
```

`$T{}` 的寫法與 `new(T)` 是等效的，使用 `&Point{10, 20}` 這類的寫法，可以同時指定結構的值域。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
