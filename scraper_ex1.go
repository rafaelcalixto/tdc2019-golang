package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
    "strings"
)

var (
    path     []string
    startURL string
)

func ProcessLinks() {
    dict := map[string]bool{}

    for _, link := range path {
        link = strings.TrimSpace(link)
        if len(link) < 10 { continue }
        if link[:4] == "http" && link != startURL && link[:len(link)-1] != startURL  {
            dict[link] = true
        }
    }
    path = path[:0]
    for k, _ := range dict {
        path = append(path, k)
    }
}

func GetLinks(resp *http.Response) {
    tags := html.NewTokenizer(resp.Body)

    for {
        _ = tags.Next()
        token := tags.Token()

        if token.Type == html.ErrorToken { break }

        if token.DataAtom == atom.A && token.Type == html.StartTagToken {
            for _, attr := range token.Attr {
                if attr.Key == "href" { path = append(path, attr.Val) }
            }
        }
    }
}

func main() {
    startURL = "https://panoramacrypto.com.br"
    ans, _ := http.Get(startURL)
    defer ans.Body.Close()
    GetLinks(ans)
    ProcessLinks()

    fmt.Println(path)
}
