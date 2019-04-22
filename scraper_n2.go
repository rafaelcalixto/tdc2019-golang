package main

import (
    "fmt"
    "net/http"
    erh "errhandler"
)

var (
    url string
)

func main() {
    url = "https://panoramacrypto.com.br"
    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)

    resp, err := client.Do(req)
    erh.Inspect(err, "client.Do", true)

    fmt.Println(resp)
}
