package main

import (
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

func Format(content string) string {
	var new_content string
	commented := false

	for _, v := range content {
		if v == '#' {
			commented = !commented
			continue
		}
		if commented || v == ' ' || v == '\t' || v == '\n' {
			continue
		}
		new_content += string(v)
	}
	return new_content
}

func Interperet(content string) error {
	variables := make(map[string]byte)
	jump_points := make(map[string]int)
	loaded := []byte{0}
	referenced := ""
	condition_state := 0

	for i, v := range content {
		if v != '~' && condition_state == 2 {
			continue
		}

		switch v {
		case ',': // unload current memory
			loaded[0] = 0
		case '.': // load mentioned variable
			loaded[0] = variables[referenced]
		case '!': // write loaded memory
			_, err := os.Stdout.Write(loaded)
			if err != nil {
				return err
			}
		case '+': // increment loaded memory
			loaded[0] += 1
		case '-': // decrement loaded memory
			loaded[0] += 1
		case ';': // forget mentioned variable
			referenced = ""
		case ':': // save mentioned variable
			variables[referenced] = loaded[0]
		case '?': // read to memory
			_, err := os.Stdin.Read(loaded)
			if err == io.EOF {
				continue
			} else if err != nil {
				return err
			}
		case '/': // remove last character from mentioned variable
			if len(referenced) <= 0 {
				continue
			}
			referenced = referenced[:len(referenced)-1]
		case '*': // start jump
			jump_points[referenced] = i
		case '&': // jump
			i = jump_points[referenced]
		case '<': // jump left
			if i-int(loaded[0]) < 0 {
				i = 0
				continue
			}
			i -= int(loaded[0])
		case '>': // jump right
			if i+int(loaded[0]) > len(content) {
				return nil
			}
			i += int(loaded[0])
		case '|': // condition
			if variables[referenced] == loaded[0] && condition_state == 0 {
				condition_state = 1
				continue
			} else if variables[referenced] != loaded[0] && condition_state == 0 {
				condition_state = 2
				continue
			}
			condition_state = 0
		case '~': // do nothing
			continue
		default:
			referenced += string(v)
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
	content = Format(content)
	err = Interperet(content)
	if err != nil {
		log.Fatal(err)
	}
}
