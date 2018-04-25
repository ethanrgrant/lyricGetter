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

func buildLyricUrl(artistName string, songName string) (string) {
    return "https://genius.com/" + artistName + "-" + songName + "-lyrics"
}

func buildSongList(artistName string) (string) {
	return "http://www.songfacts.com/artist-" + artistName + ".php"
}

func getLyrics(artistName string, songName string, songLyricCh chan SongLyrics, finishCh chan bool) {
    songLyrics := SongLyrics{Song: songName, Lyrics: "" }
    defer func() { finishCh <- true; songLyricCh <- songLyrics }()
    doc, err := goquery.NewDocument(buildLyricUrl(artistName, songName))
    if err != nil {
        songLyrics.Lyrics = "Failed to get lyrics for this song"
        return
    }
    lyrics := ""
    doc.Find(".lyrics").Find("a").Each(func(i int, s *goquery.Selection) {
       lyrics += s.Text()
    })
    songLyrics.Lyrics = lyrics
}

func GetSongList(artistName string, songNum * int, songLyricCh chan SongLyrics, finishCh chan bool) {
	doc, err := goquery.NewDocument(buildSongList(artistName))
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
