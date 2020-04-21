package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"path/filepath"
)

// ReadFile ziped data
func ReadFile(fileName string) {
	r, err := zip.OpenReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		ftxt, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(ftxt)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			break
		}
	}
}

func main() {

	files, err := filepath.Glob("../govdata/*")
	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range files {
		fmt.Printf("Extracting from '%s'", fileName)
		fmt.Printf("\n")

		ReadFile(fileName)
	}

}
