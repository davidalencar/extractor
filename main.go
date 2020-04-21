package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"path/filepath"

	"github.com/schollz/progressbar"
)

// ReadFile opens the zipped file to read the text file inside
func ReadFile(fileName string) {

	r, err := zip.OpenReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		ftxt, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(ftxt)

		for scanner.Scan() {
			//fmt.Println(scanner.Text())

		}
	}
}

func main() {

	files, err := filepath.Glob("../govdata/*.zip")
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.New(len(files))

	for _, fileName := range files {

		ReadFile(fileName)
		bar.Clear()
		bar.Describe(fmt.Sprintf("Extracting from %q	", filepath.Base(fileName)))
		bar.Add(1)

	}

	fmt.Println()
}
