package main

import (
	"flag"
	"fmt"
	"github.com/simonhull/Tag-Parser/audiotag"
	"github.com/simonhull/Tag-Parser/go-mp4tag"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func parseAudioTag(filePath string) (audiotag.Metadata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Parse Error:", err.Error())
	}
	m, err := audiotag.ReadFrom(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(m.Format()) // The detected format.
	log.Print(m.Title())  // The title of the track (see Metadata interface for more details).

	return m, nil
}
func main() {

	filePath := flag.String("file", "", "path to the file")
	title := flag.String("title", "", "title")
	publisher := flag.String("copyright", "", "copyright")
	year := flag.Int("year", 0, "year")
	narrator := flag.String("narrator", "", "narrator")
	genre := flag.Uint("genre", 0, "genre")
	subtitle := flag.String("sub", "", "subtitle")
	language := flag.String("lang", "", "language")
	isbn := flag.String("isbn", "", "isbn")
	asin := flag.String("asin", "", "asin")
	mvn := flag.String("mvn", "", "mvn")
	seriesPart := flag.String("series-part", "", "seris-part")
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
	f.Close()

	mp4, err := mp4tag.Open(*filePath)
	if err != nil {
		panic(err)
	}
	defer mp4.Close()

	writeTags := &mp4tag.MP4Tags{}

	// Custom: map[string]string{
	// 	"Subtitle":    "Golang Tutorial",
	// 	"LANGUAGE":    "English",
	// 	"ISBN":        "English",
	// 	"ASIN":        "English",
	// 	"\xa9mvn":     "My Series",
	// 	"series-part": "Seris-Part",
	// },

	if *title != "" {
		writeTags.Title = *title
	}

	if *publisher != "" {
		writeTags.Copyright = *publisher
	}

	if *year != 0 {
		writeTags.Year = int32(*year)
	}

	if *genre != 0 {
		writeTags.Genre = mp4tag.Genre(*genre)
	}

	if *narrator != "" {
		writeTags.Composer = *narrator
	}

	writeTags.Custom = make(map[string]string)

	if *subtitle != "" {
		writeTags.Custom["Subtitle"] = *subtitle
	}

	if *language != "" {
		writeTags.Custom["LANGUAGE"] = *language
	}

	if *isbn != "" {
		writeTags.Custom["ISBN"] = *isbn
	}

	if *asin != "" {
		writeTags.Custom["ASIN"] = *asin
	}
	if *mvn != "" {
		writeTags.Custom["\xa9mvn"] = *mvn
	}

	if *seriesPart != "" {
		writeTags.Custom["series-part"] = *seriesPart
	}

	err = mp4.Write(writeTags, []string{})
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Raw Data: %v\n", rawMetaData)
}
