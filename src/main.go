package main

import (
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
		if commented || v == ' ' || v == '\t' || v == '\n' { // rempve commented code & white spaces
			continue
		}
		new_content += string(v)
	}
	return new_content
}

type Reference struct {
	ByteValue byte
	GotoIndex int
}

func Interperet(content string) error {
	references := make(map[string]Reference)
	main_byte := []byte{0}
	referenced := ""
	condition_state := 0
	index := 0

	for index < len(content) {

		if content[index] != '}' && condition_state == 2 { // while false, skip everything besides the if statement end
			index++
			continue
		}

		switch content[index] {
		case ',': // unload current memory
			main_byte[0] = 0
		case '.': // load mentioned variable
			main_byte[0] = references[referenced].ByteValue
		case '!': // write main_byte memory
			_, err := os.Stdout.Write(main_byte)
			if err != nil {
				return err
			}
		case '+': // increment main_byte
			main_byte[0] += 1
		case '-': // decrement main_byte
			main_byte[0] -= 1
		case ';': // forget mentioned variable
			referenced = ""
		case ':': // save mentioned variable
			references[referenced] = Reference{main_byte[0], references[referenced].GotoIndex}
		case '?': // read to memory
			_, err := os.Stdin.Read(main_byte)
			if err != nil {
				return err
			}
		case '/': // remove last character from mentioned variable
			if len(referenced) > 0 {
				referenced = referenced[:len(referenced)-1]
			}
		case '*': // start goto
			references[referenced] = Reference{references[referenced].ByteValue, index}
		case '&': // goto
			index = references[referenced].GotoIndex
			continue
		case '<': // decrease jump distance
			if index-int(main_byte[0]) < 0 {
				index = 0
				continue
			}
			index -= int(main_byte[0])
			continue
		case '>': // goto right
			if index+int(main_byte[0]) > len(content) {
				return nil
			}
			index += int(main_byte[0])
			continue
		case '{': // condition check
			if references[referenced].ByteValue == main_byte[0] && condition_state == 0 {
				condition_state = 1
			} else if references[referenced].ByteValue != main_byte[0] && condition_state == 0 {
				condition_state = 2
			}
		case '}': // condition end
			condition_state = 0
		default: // append to reference
			if content[index] != '~' { // do nothing / no opp
				referenced += string(content[index])
			}
		}
		index++
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
