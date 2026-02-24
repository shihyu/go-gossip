<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 介面組合

</div>

  

有時，可能會想要基於某個已定義的介面，並新增自己的行為，在 Go 中，這類似於結構中方法的查找，只要在定義介面時，內嵌想要的介面名稱就可以了。例如：

``` prettyprint
package main

import "fmt"

type ParentTester interface {
    ptest()
}

type ChildTester interface {
    ParentTester
    ctest()
}

type Subject struct {
    name string
}

func (s *Subject) ptest() {
    fmt.Printf("ptest %s\n", s)
}

func (s *Subject) ctest() {
    fmt.Printf("ctest %s\n", s)
}

func main() {
    var tester ChildTester = &Subject{"Test"}
    tester.ptest()
    tester.ctest()
}
```

在上面，`Subject` 必須實作 `ParentTester` 與 `ChildTest` 中定義的全部行為，其實例才可以被指定 `ChildTest`。你也可以介面中包含多個介面：

``` prettyprint
package main

import "fmt"

type SuperTester interface {
    stest()
}

type ParentTester interface {
    ptest()
}

type ChildTester interface {
    SuperTester
    ParentTester
    ctest()
}

type Subject struct {
    name string
}

func (s *Subject) stest() {
    fmt.Printf("stest %s\n", s)
}

func (s *Subject) ptest() {
    fmt.Printf("ptest %s\n", s)
}

func (s *Subject) ctest() {
    fmt.Printf("ctest %s\n", s)
}

func main() {
    var tester ChildTester = &Subject{"Test"}
    tester.stest()
    tester.ptest()
    tester.ctest()
}
```

如果多個介面間的行為重複定義了，就會出現 duplicate method 的錯誤。（這是個有爭議性的特性，因為許多人認為，實際上雖然在介面語法上確實重複定義了行為，然而就 Duck typing 的精神來看，結構上只要有實作行為就可以了，事實上在其他語言中，像是 Java 中，類似的情況並不會發生編譯錯誤，有關此議題，可參考 [golang/go 的 此 issue](https://github.com/golang/go/issues/6977)）。

雖然說這像是介面有了繼承方面的語法，然而更精確地說，應該是行為的內嵌，因此，只要是有實現相關行為，就算沒有被包含在某個介面中，也可以做介面轉換：

``` prettyprint
package main

import "fmt"

type SuperTester interface {
    stest()
}

type ParentTester interface {
    ptest()
}

type ChildTester interface {
    SuperTester
    ParentTester
    ctest()
}

type Tester interface {
    stest()
    ptest()
    ctest()
}

type Subject struct {
    name string
}

func (s *Subject) stest() {
    fmt.Printf("stest %s\n", s)
}

func (s *Subject) ptest() {
    fmt.Printf("ptest %s\n", s)
}

func (s *Subject) ctest() {
    fmt.Printf("ctest %s\n", s)
}

func main() {
    var ctester ChildTester = &Subject{"Test"}
    var tester Tester = ctester
    tester.stest()
    tester.ptest()
    tester.ctest()
}
```

有些文件會說，在介面有組合關係時，子介面的實例可以指定給父介面，反之就不行，這種說法不能說是錯，畢竟就上例來說，`ChildTester` 介面的實例，被指定給 `ParentTester` 介面時，從編譯器的角度來看，`ChildTester` 介面確實是有 `ParentTester` 介面的行為；反過來的話，`ParentTester` 介面被指定給 `ChildTester` 介面時，編譯器是看不到 `ParentTester` 介面上，會有 `ChildTester` 介面行為的，當然會發生錯誤。

更精確來說，Go 本身並非基於類別，沒有提供繼承語法，也就沒有父介面、子介面的概念，以上僅僅只是以行為的內嵌實現了繼承的概念，因而是就看不看得到相關的行為，來判斷是否可通過編譯。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
