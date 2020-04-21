package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/schollz/progressbar"
)

func main() {

	files, err := filepath.Glob("../govdata/*.zip")
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.New(len(files))

	for _, fileName := range files {

		//
		bar.Clear()
		bar.Describe(fmt.Sprintf("Extracting from %q	", filepath.Base(fileName)))
		bar.Add(1)

	}

	fmt.Println()
}
