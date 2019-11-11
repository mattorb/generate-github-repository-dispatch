package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

func stringFromFile(filename string) (string, error) {
	buf := bytes.NewBuffer(nil)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return "", err
	}

	io.Copy(buf, f)
	return string(buf.Bytes()), nil
}

func stringFromConsole() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimRight(text, "\n"), err
}
