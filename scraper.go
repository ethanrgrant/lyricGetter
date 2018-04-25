package main

import (
	"fmt"
    //"strings"

	// import third party libraries
	"github.com/PuerkitoBio/goquery"
)

type SongLyrics struct {
	Lyrics string
	Song   string
}

func getLyrics(artistName string, songName string, songLyricCh chan SongLyrics, finishCh chan bool) {
	songLyricCh <- SongLyrics{Song: songName, Lyrics: "These are lyrics"}
	finishCh <- true
}

func GetSongList(artistName string, songNum * int, songLyricCh chan SongLyrics, finishCh chan bool) {
	finalUrl := "http://www.songfacts.com/artist-" + artistName + ".php"
	doc, err := goquery.NewDocument(finalUrl)
	if err != nil {
		fmt.Println("ERROR: failed to get songs for " + artistName + "!")
		return
	}

    *songNum = doc.Find(".songullist-orange ").Find("a").Each(func(i int, s *goquery.Selection) {}).Length()

	doc.Find(".songullist-orange ").Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the lyrics
		go getLyrics(artistName, s.Text(), songLyricCh, finishCh)
	})
}
