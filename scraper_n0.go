package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func main() {
    ans, _ := http.Get("https://api.kraken.com/0/public/Assets")
    body, _ := ioutil.ReadAll(ans.Body)

    cryptos := string(body)
    cryptosMAP := make(map[string]interface{})

    _ = json.Unmarshal([]byte(cryptos), &cryptosMAP)
    ans.Body.Close()

    for key, _ := range cryptosMAP {
        fmt.Println(key)
    }
}
