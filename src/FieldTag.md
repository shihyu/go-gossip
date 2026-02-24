<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 結構欄位標籤

</div>

  

對於 JSON 或 XML 等具有結構性的資料，在 Go 中經常會使用 `struct` 定義資料結構，例如，底下這個程式可以將簡單的結構轉為 JSON：

``` go
package main

import (
    "fmt"
    "reflect"
    "strings"
)

type Customer struct {
    Name string
    City string
}

func ToJSON(obj interface{}) string {
    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)

    var b []string  
    for i, n := 0, t.NumField(); i < n; i++ {
        f := t.Field(i)
        b = append(b, fmt.Sprintf(`"%s": "%s"`, f.Name, v.FieldByName(f.Name)))
    }

    return fmt.Sprintf("{%s}", strings.Join(b, ","))
}

func main() {
    cust := Customer{"Justin", "Kaohsiung"}
    // 顯示 {"Name": "Justin","City": "Kaohsiung"}
    fmt.Println(ToJSON(cust))
}
```

然而，Go 的慣例中，公開的結構欄位名稱通常是大寫的，如果你的 JSON 要求的是小寫的欄位名稱，或者是其他名稱，可以使用欄位標籤（Field tag）。例如：

``` go
package main

import (
    "fmt"
    "reflect"
    "strings"
)

type Customer struct {
    Name string `name`
    City string `city`
}

func ToJSON(obj interface{}) string {
    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)
    var b []string

    for i, n := 0, t.NumField(); i < n; i++ {
        f := t.Field(i)
        b = append(b, fmt.Sprintf(`"%s": "%s"`, f.Tag, v.FieldByName(f.Name)))
    }

    return fmt.Sprintf("{%s}", strings.Join(b, ","))
}

func main() {
    cust := Customer{"Justin", "Kaohsiung"}
    // 顯示 {"name": "Justin","city": "Kaohsiung"}
    fmt.Println(ToJSON(cust))
}
```

欄位標籤可以在反射時，使用 `Field` 的 `Tag` 來取得，雖然欄位標籤可以是任意格式字串，然而慣例上，會由 `key: "value"` 的格式組成，符合此格式的話，可以使用 `Tag` 的 `Lookup` 來查找 `value`，它傳回兩個值，第一個值是 `value`，第二個值指出是否有對應的名稱，例如：

``` go
package main

import (
    "fmt"
    "reflect"
    "strings"
)

type Customer struct {
    Name string `json:"name"`
    City string `json:"city"`
}

func ToJSON(obj interface{}) string {
    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)
    var b []string

    for i, n := 0, t.NumField(); i < n; i++ {
        f := t.Field(i)
        fv, _ := f.Tag.Lookup("json")
        b = append(b, fmt.Sprintf(`"%s": "%s"`, fv, v.FieldByName(f.Name)))
    }

    return fmt.Sprintf("{%s}", strings.Join(b, ","))
}

func main() {
    cust := Customer{"Justin", "Kaohsiung"}
    // 顯示 {"name": "Justin","city": "Kaohsiung"}
    fmt.Println(ToJSON(cust))
}
```

實際上，如果要將結構轉為 JSON 格式字串，可以使用 `encoding/json`，例如：

``` go
package main

import (
    "encoding/json"
    "fmt"
)

type Customer struct {
    Name string `json:"name"`
    City string `json:"city"`
}

func main() {
    cust := Customer{"Justin", "Kaohsiung"}
    b, _ := json.Marshal(cust)
    // 顯示 {"name": "Justin","city": "Kaohsiung"}
    fmt.Println(string(b))
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
