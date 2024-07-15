package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unitnotes/audiotag"
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
