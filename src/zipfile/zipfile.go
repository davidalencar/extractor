package zipfile

import (
	"archive/zip"
	"bufio"
	"log"
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
