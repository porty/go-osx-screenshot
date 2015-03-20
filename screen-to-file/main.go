package main

import (
	"fmt"
	"os"
	"strings"

	screenshot "github.com/porty/go-osx-screenshot"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify file to save to")
		return
	}
	filename := os.Args[1]
	var format screenshot.SaveFormat
	if strings.HasSuffix(strings.ToLower(filename), ".png") {
		format = screenshot.FormatPng
	} else {
		format = screenshot.FormatJpeg
	}

	err := screenshot.SaveScreenshotToFile(filename, format)
	if err != nil {
		panic(err)
	}
	fmt.Println("Screenshot saved to " + filename)
}
