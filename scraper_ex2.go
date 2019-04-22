package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
    "strings"
)

var (
    startURL  string
    linkList []string
)

func ProcessLinks() {
    dict := map[string]bool{}

    for _, link := range linkList {
        link = strings.TrimSpace(link)
        if len(link) < 10 { continue }
        if link[:4] == "http" && link != startURL && link[:len(link)-1] != startURL  {
            dict[link] = true
        }
    }
    linkList = linkList[:0]
    for k, _ := range dict {
        linkList = append(linkList, k)
    }
}

func GetLinks(startURL string) {
    ans, _ := http.Get(startURL)
    defer ans.Body.Close()
    tags := html.NewTokenizer(ans.Body)

    for {
        _ = tags.Next()
        token := tags.Token()

        if token.Type == html.ErrorToken { break }

        if token.DataAtom == atom.A && token.Type == html.StartTagToken {
            for _, attr := range token.Attr {
                if attr.Key == "href" { linkList = append(linkList, attr.Val) }
            }
        }
    }
    ProcessLinks()
}

func CountLinks(listPath []string) {
    var newLinksList []string

    for _, link := range listPath {
        go GetLinks(link)
    }
}

func UpdateValue(l *[]string) {
    *l = linkList
    linkList = linkList[:0]
}

func main() {
    var startLinks []string
    startURL = "https://panoramacrypto.com.br"
    GetLinks(startURL)
    UpdateValue(&startLinks)

    CountLinks(startLinks)
}
