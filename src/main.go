package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Interperet(content string) error {
	variables := make(map[string]byte)
	loaded := []byte{0}
	mentioned_variable := ""
	for _, v := range content {
		switch v {
		case ',': // unload current memory
			loaded[0] = 0
		case '.': // load mentioned variable
			loaded[0] = variables[mentioned_variable]
		case '!': // write loaded memory
			_, err := os.Stdout.Write(loaded)
			if err != nil {
				return err
			}
		case '+': // increment loaded memory
			loaded[0]++
		case '-': // decrement loaded memory
			loaded[0]--
		case ';': // forget mentioned variable
			mentioned_variable = ""
		case ':': // save mentioned variable
			variables[mentioned_variable] = loaded[0]
		case '?': // read to memory
			_, err := os.Stdin.Read(loaded)
			if err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}
		default:
			mentioned_variable += string(v)
		}
	}
	return nil
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeName := filepath.Base(exePath)

	args := os.Args

	var start int

	for i, arg := range args {
		if strings.Compare(arg, exeName) == 0 {
			start = i
		}
	}

	if len(args)-start <= 1 {
		log.Fatal("no input file")
	}
	content, err := GetContent(args[start+1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}
