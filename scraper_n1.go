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
    url = "https://thedevelopersconference.com.br/tdc/2019/florianopolis/trilha-go"

    ans, err := http.Get(url)
    erh.Inspect(err, "http.Get", true)

    fmt.Println(ans)
}
