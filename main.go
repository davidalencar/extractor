package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"log"
	"path/filepath"

	"github.com/schollz/progressbar"
)

func ReadTxtFile(txtFile io.ReadCloser) {

	scanner := bufio.NewScanner(txtFile)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		break
	}
}

func ExtractTxtFileFrom(zippedFileName string) io.ReadCloser {

	r, err := zip.OpenReader(zippedFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		ftxt, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		return ftxt
	}
	return nil
}

func RetriveAllZippedFiles(dir string) []string {

	files, err := filepath.Glob("../govdata/*.zip")
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func main() {

	files := RetriveAllZippedFiles("../govdata/*.zip")

	bar := progressbar.New(len(files))

	for _, fileName := range files {

		txtFile := ExtractTxtFileFrom(fileName)
		if txtFile != nil {
			ReadTxtFile(txtFile)
		}

		bar.Clear()
		bar.Describe(fmt.Sprintf("Extracting from %q ", filepath.Base(fileName)))
		bar.Add(1)

	}

	fmt.Println()
}
