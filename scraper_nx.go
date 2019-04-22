package main

import (
    "fmt"
    "math/rand"
    "time"
    "net/url"
    "net/http"
    erh "errhandler"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

var proxies = []string{
    "http://138.68.150.51:8118",
    "http://82.206.132.106:80",
    "http://178.62.88.110:8118",
}

func RandomString(options []string) string {
    rand.Seed(time.Now().Unix())
    randNum := rand.Int() % len(options)
    return options[randNum]
}

func scrapeClient(targetURL string, proxyOptions []string, userAgents []string) (*http.Response, error) {
    selectedProxy := RandomString(proxyOptions)
    proxyURL, _ := url.Parse(selectedProxy)

    client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

    req, err := http.NewRequest("GET", targetURL, nil)
    erh.Inspect(err, "client.Do", true)

    req.Header.Set("User-Agent", RandomString(userAgents))

    resp, err := client.Do(req)
    erh.Inspect(err, "client.Do", true)

    if err != nil { return nil, err } else { return resp, nil }
}

func main() {
    var url string = "https://thedevelopersconference.com.br/tdc/2019/florianopolis/trilha-go"
    resp, _ := scrapeClient(url, proxies, userAgents)
    fmt.Println(resp)
}
