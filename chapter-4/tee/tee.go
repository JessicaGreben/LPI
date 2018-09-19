package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var path string
var append bool

func main() {
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	writeToStdOut(text)
	if err := writeToFile(text, path); err != nil {
		fmt.Print(err)
	}
}

func writeToStdOut(text string) {
	// Write to stdout
	fmt.Println(text)
}

func writeToFile(text string, filePath string) error {
	var wrOnlyOrAppend = os.O_WRONLY
	switch {
	case append:
		wrOnlyOrAppend = os.O_APPEND
	}

	fd, err := os.OpenFile(path, wrOnlyOrAppend|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer fd.Close()

	if _, err := fd.WriteString(text); err != nil {
		return err
	}
	return nil
}

func init() {
	flag.StringVar(&path, "path", "./output.txt", "Path to file to write to. Overwrites content.")
	flag.BoolVar(&append, "a", false, "Appends content to file.")
}
