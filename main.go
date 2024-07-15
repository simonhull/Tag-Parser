package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/unitnotes/audiotag"
)

func main() {

	filePath := flag.String("file", "", "path to the file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a file path using the -file flag")
		return
	}

	audioTetadata, err := parseAudioTag(*filePath)
	if err != nil {
		return
	}

	var rawMetaData map[string]interface{} = audioTetadata.Raw()

	// fmt.Printf("Title: %s\n", audioTetadata.Title())
	// publisher, series, series sequence, isbn

	fmt.Printf("Author: %v\n", audioTetadata.Composer()+audioTetadata.AlbumArtist())
	fmt.Printf("Title: %v\n", audioTetadata.Album())
	fmt.Printf("Subtitle: %v\n", rawMetaData["Subtitle"])
	fmt.Printf("Publisher: %v\n", rawMetaData["cprt"])
	fmt.Printf("Published Year: %d\n", audioTetadata.Year())
	fmt.Printf("Narrator(s): %v\n", audioTetadata.Composer())
	fmt.Printf("Genre: %v\n", audioTetadata.Genre())
	fmt.Printf("Series: %v\n", rawMetaData["\xa9mvn"])
	fmt.Printf("Series Sequence: %v\n", rawMetaData["series-part"])
	fmt.Printf("Language: %v\n", rawMetaData["LANGUAGE"])
	fmt.Printf("ISBN: %v\n", rawMetaData["ISBN"])
	fmt.Printf("ASIN: %v\n", rawMetaData["ASIN"])
	fmt.Printf("Cover: %v\n", rawMetaData["covr"])

	f, _ := os.Open(*filePath)
	fi, _ := f.Stat()

	// Extract file extension
	ext := filepath.Ext(*filePath)

	fmt.Println("File size:", fi.Size(), "bytes")

	// Print the file extension
	fmt.Println("File extension:", ext)

	fmt.Printf("Duration: %d\n", audioTetadata.Duration())

	// fmt.Printf("Chapter list: %v\n", rawMetaData["chpl"])

	var ChapterList = rawMetaData["chpl"].([]audiotag.Chapter)

	for index, chapter := range ChapterList {
		if index == len(ChapterList)-1 {
			chapter.EndTime = strconv.Itoa(audioTetadata.Duration())
		}
		fmt.Println(chapter)
	}
	// fmt.Printf("Raw Data: %v\n", rawMetaData)
}
