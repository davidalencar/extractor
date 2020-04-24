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

func readTxtFile(txtFile io.ReadCloser) {

	scanner := bufio.NewScanner(txtFile)

	for scanner.Scan() {

		line := scanner.Text()

		if line[0:1] == "1" {
			fmt.Println(line[3:17])
		}
		if line[0:1] == "2" {
			fmt.Println(line)
			break
		}
	}
}

func extractTxtFileFrom(zippedFileName string) {

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
		readTxtFile(ftxt)
	}

}

func retriveAllZippedFiles(dir string) []string {

	files, err := filepath.Glob("../govdata/*.zip")
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func main() {

	files := retriveAllZippedFiles("../govdata/*.zip")

	bar := progressbar.New(len(files))

	for _, fileName := range files {

		bar.Clear()
		bar.Describe(fmt.Sprintf("Extracting from %q ", filepath.Base(fileName)))
		bar.RenderBlank()
		extractTxtFileFrom(fileName)
		bar.Add(1)

	}

	fmt.Println()
}
