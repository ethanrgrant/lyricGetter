package main

import (
	"fmt"
	"os"
)

func constructUrl(artist string) string {
	return "https://genius.com/artists/" + artist
}

func processArtist(songNum int, finishCh chan bool, songLyricCh chan SongLyrics) {
    totalSeen := 0
    for {
        select {
        case songLyric := <-songLyricCh:
            fmt.Printf("\n%s \n %s", songLyric.Song, songLyric.Lyrics)
        case <-finishCh:
            // if all songs have been processed move on
            totalSeen++
            if totalSeen >= songNum {
                return
            }
        }
    }
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("please include one artists name as cmd line args")
	}

	artists := os.Args[1:]
	songLyricCh := make(chan SongLyrics)
	finishCh := make(chan bool)

	for _, artist := range artists {
        songNum := 0
		GetSongList(artist, &songNum, songLyricCh, finishCh)
        processArtist(songNum, finishCh, songLyricCh)
	}
}
