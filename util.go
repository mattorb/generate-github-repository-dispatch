package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func stringFromFile(filename string) string {
	buf := bytes.NewBuffer(nil)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.Copy(buf, f)
	return string(buf.Bytes())
}
