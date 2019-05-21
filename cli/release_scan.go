package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    response, err := http.Get("https://github.com/raynear/SimpleCoin/releases")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

    document.Find("a").Each(func(i int, s *goquery.Selection) {
        if(s.Text() == "Latest release") {
            s.Parent().Parent().Parent().Children().Find("a").Each(func(j int, s2 *goquery.Selection) {
                bVal,exist := s2.Attr("title")
                if (exist) {
                    fmt.Println(bVal)
                }
            })
        }
    })
}
