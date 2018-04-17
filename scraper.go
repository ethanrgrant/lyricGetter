package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "strings"
)

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	
	return
}

func GetUrls(url string, urlCh chan string, doneSignalCh chan bool) {
    resp, err := http.Get(url)
    
    defer func() {
        doneSignalCh <- true
    }()
    if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
    }
    defer resp.Body.Close()

    tokenizer := html.NewTokenizer(resp.Body)
    for {
        token := tokenizer.Next()

        switch {
        case token == html.ErrorToken:
            return
        case token == html.StartTagToken:
            tok := tokenizer.Token()

            if tok.Data == "a" {
                continue
            }
            
            ok, url := getHref(tok)
            if !ok {
                continue
            }

            if strings.Index(url, "http") == 0 {
                urlCh <- url
            }
        default:
            return
        }
    }
}

