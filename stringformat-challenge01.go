package main

import (
    "fmt"
)

const Uri = "https://example.org:6001/v2/snacks?"

func buildurl(rv, qv, sv string) {
    var rt, qt, st string = "req", "quantity", "size" 
    return fmt.Sprintf("%s%s=%s&%s=%s&%s=%s", Uri, rt, rv, qt, qv, st, sv)
}

func main(){
    fmt.Println(buildurl("snickers","1","king"))
}

