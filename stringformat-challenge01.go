package main

import (
    "fmt"
)

const uri = "https://example.org:6001/v2/snacks?"

func main() {
    var rt, qt, st string = "req", "quantity", "size" 
    var rv, qv, sv string = "snickers", "1", "king" 
    res := fmt.Sprintf("%s%s=%s&%s=%s&%s=%s", uri, rt, rv, qt, qv, st, sv)
    fmt.Println(res)
}

