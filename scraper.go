package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "strings"

    // import third party libraries
    "github.com/PuerkitoBio/goquery"
)

type AristSongLyrics struct {
    Artist string
    Lyrics string
    Song string
}

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

// Gets all urls associated with the given url
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
            //fmt.Println(tok)

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
        }
    }
}


func GetSongList(artistName string) {
    finalUrl := "http://www.songfacts.com/artist-" + artistName + ".php"

/*
    artistSong := Artist{Artist: artistName
    defer func() {
       outArtistSongChan <-  artistSong
    }*/
    doc, err := goquery.NewDocument(finalUrl)
    if err != nil {
        fmt.Println("ERROR: failed to get songs for " + artistName + "!")
        return
    }
    doc.Find(".songullist-orange").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
        band := s.Find("a").Text()
        fmt.Printf("%s: %s\n", i, band)
    })
}
