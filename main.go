package main

import (
    "fmt"
    "os"
)

func constructUrl(artist string) (string)  {
    return "https://genius.com/artists/" + artist
}

func main() {

    if len(os.Args) == 1 {
        fmt.Println("please include one artists name as cmd line args")
    }

    artists := os.Args[1:]
    foundUrls := make([]string, 0)
    urlCh := make(chan string)
    finishCh := make(chan bool)
    for _, artist := range artists {
        cleanUrl := constructUrl(artist)
        fmt.Println(cleanUrl)
        go GetUrls(cleanUrl, urlCh, finishCh)
    }

    for artistIdx := 0; artistIdx < len(artists); {
        select {
        case url := <-urlCh:
            foundUrls = append(foundUrls, url) 
        case <- finishCh:
            artistIdx++
        }
    }

    for _, url := range foundUrls {
        fmt.Printf("url: %s", url)
    }
}
