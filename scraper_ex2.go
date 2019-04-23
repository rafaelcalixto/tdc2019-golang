package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
    "strings"
    "sync"
)

var (
    startURL  string
    linkList []string
    wg sync.WaitGroup
)

func ProcessLinks(links []string) ([]string) {
    dict := map[string]bool{}

    for _, link := range links {
        link = strings.TrimSpace(link)
        if len(link) < 10 { continue }
        if link[:4] == "http" && link != startURL && link[:len(link)-1] != startURL  {
            dict[link] = true
        }
    }
    links = links[:0]
    for k, _ := range dict {
        links = append(links, k)
    }
    fmt.Println("Total links found:", len(links))
    return links
}

func GetLinks(t *html.Tokenizer, links []string) ([]string) {
    for {
        _ = t.Next()
        token := t.Token()

        if token.Type == html.ErrorToken { break }
        if token.DataAtom == atom.A && token.Type == html.StartTagToken {
            for _, attr := range token.Attr {
                if attr.Key == "href" { links = append(links, attr.Val) }
            }
        }
    }
    return links
}

func Scraper(c chan []string, url string) {
    fmt.Println("going check -> " + url)
    var linkList []string
    defer wg.Done()
    ans, _ := http.Get(url)
    defer ans.Body.Close()
    tags := html.NewTokenizer(ans.Body)
    c <- ProcessLinks(GetLinks(tags, linkList))
}

func CountLinks(links []string) {
    newChan := make(chan []string, 30)
    for _, l := range links {
        wg.Add(1)
        go Scraper(newChan, l)
    }
    wg.Wait()
    close(newChan)
}

func main() {
    var startLinks []string
    chanLinks := make(chan []string, 30)
    startURL = "https://panoramacrypto.com.br"
    wg.Add(1)
    go Scraper(chanLinks, startURL)
    wg.Wait()
    close(chanLinks)
    startLinks = <-chanLinks
    CountLinks(startLinks)
}
