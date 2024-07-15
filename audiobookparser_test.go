package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"testing"
)

func TestParseM4B(t *testing.T) {
	filePath := "Eldest - Christopher Paolini.m4b" // replace with your sample m4b file path

	audioTetadata, err := parseAudioTag(filePath)
	if err != nil {
		t.Fatalf("Failed to parse M4B file: %v", err)
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

	f, _ := os.Open(filePath)
	fi, _ := f.Stat()

	// Extract file extension
	ext := filepath.Ext(filePath)

	fmt.Println("File size:", fi.Size(), "bytes")

	// Print the file extension
	fmt.Println("File extension:", ext)

	fmt.Printf("Duration: %d\n", audioTetadata.Duration())

	fmt.Printf("Chapter list: %v\n", rawMetaData["chpl"])
	fmt.Printf("Chapter list: %v\n", rawMetaData)

}

func TestParseMP3(t *testing.T) {
	filePath := "sample.mp3" // replace with your sample m4b file path

	audioTetadata, err := parseAudioTag(filePath)
	if err != nil {
		t.Fatalf("Failed to parse M4B file: %v", err)
	}

	var rawMetaData map[string]interface{} = audioTetadata.Raw()

	// fmt.Printf("Title: %s\n", audioTetadata.Title())
	// publisher, series, series sequence, isbn

	fmt.Printf("Author: %v\n", audioTetadata.Composer()+audioTetadata.AlbumArtist())
	fmt.Printf("Title: %v\n", audioTetadata.Album())
	fmt.Printf("Subtitle: %v\n", rawMetaData["TIT3"])
	fmt.Printf("Publisher: %v\n", rawMetaData["TCOP"])
	fmt.Printf("Published Year: %d\n", audioTetadata.Year())
	fmt.Printf("Narrator(s): %v\n", audioTetadata.Composer())
	fmt.Printf("Genre: %v\n", audioTetadata.Genre())
	fmt.Printf("Series: %v\n", rawMetaData["TXXX_4"])
	fmt.Printf("Series Sequence: %v\n", rawMetaData["TXXX_5"])
	fmt.Printf("Language: %v\n", rawMetaData["TXXX_3"])
	fmt.Printf("ISBN: %v\n", rawMetaData["ISBN"])
	fmt.Printf("ASIN: %v\n", rawMetaData["TXXX_0"])
	fmt.Printf("Cover: %v\n", rawMetaData["APIC"])

	f, _ := os.Open(filePath)
	fi, _ := f.Stat()

	// Extract file extension
	ext := filepath.Ext(filePath)

	fmt.Println("File size:", fi.Size(), "bytes")

	// Print the file extension
	fmt.Println("File extension:", ext)

	fmt.Printf("Duration: %d\n", rawMetaData["TLEN"])

	fmt.Printf("Chapter list: %v\n", rawMetaData["chpl"])
	fmt.Printf("Chapter list: %v\n", rawMetaData)

}

func Float32ToBytes(f float32) []byte {
	// Convert float32 to uint32
	bits := math.Float32bits(f)
	// Create a byte slice to hold the bytes
	bytes := make([]byte, 4)
	// Convert uint32 to bytes (little-endian)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}
