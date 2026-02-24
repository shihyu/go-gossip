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
