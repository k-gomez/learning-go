package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("This is not a zip file")
	notAZip := bytes.NewReader(data)
	_, err := zip.NewReader(notAZip, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
