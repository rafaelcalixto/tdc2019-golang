package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    ans, _ := http.Get("https://panoramacrypto.com.br")
    body, _ := ioutil.ReadAll(ans.Body)

    tags := string(body)
    ans.Body.Close()

    fmt.Println(tags)
}
