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
	loaded := []byte{0}
	mentioned_variable := ""
	var condition_stack []byte // figure it out

	for i, v := range content {
		switch v {
		case ',': // unload current memory
			loaded[0] = 0
			fmt.Println("unloaded memory")
		case '.': // load mentioned variable
			loaded[0] = variables[mentioned_variable]
			fmt.Println("loaded: ", loaded[0], " from variable: ", mentioned_variable)
		case '!': // write loaded memory
			_, err := os.Stdout.Write(loaded)
			if err != nil {
				return err
			}
			fmt.Println("wrote memory: ", loaded[0])
		case '+': // increment loaded memory
			loaded[0] += 1
			fmt.Println("incremented memory to: ", loaded[0])
		case '-': // decrement loaded memory
			loaded[0] += 1
			fmt.Println("decremented memory to: ", loaded[0])
		case ';': // forget mentioned variable
			mentioned_variable = ""
			fmt.Println("forgot mentioned variable")
		case ':': // save mentioned variable
			variables[mentioned_variable] = loaded[0]
			fmt.Println("saved: ", loaded[0], " to: ", mentioned_variable)
		case '?': // read to memory
			_, err := os.Stdin.Read(loaded)
			if err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}
			fmt.Println("read: ", loaded[0])
		case '/': // remove last character from mentioned variable
			mentioned_variable = mentioned_variable[:len(mentioned_variable)-1]
			fmt.Println("mentioned: ", mentioned_variable)
		case '<': // go back by an amount
			if i-int(loaded[0]) < 0 {
				i = 0
				continue
			}
			i -= int(loaded[0])
		case '>': // go forward by an amount
			if i+int(loaded[0]) >= len(content) {
				return nil
			}
			i += int(loaded[0])
		case '{': // condition start
		case '}': // condition end
		default:
			mentioned_variable += string(v)
			fmt.Println("mentioned: ", mentioned_variable)
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
