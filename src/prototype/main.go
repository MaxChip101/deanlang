package main

import (
	"fmt"
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

func DebugOperation(debug bool, message string) {
	if debug {
		log.Println(message)
	}
}

func Interperet(content string, debug bool) error {
	references := make(map[string]Reference)
	main_byte := []byte{0}
	referenced := ""
	condition_skip := false
	index := 0

	for index < len(content) {

		if content[index] != '}' && condition_skip { // while false, skip everything besides the if statement end
			DebugOperation(debug, fmt.Sprintf("condition is false, skipping: '%c' at index: %d", content[index], index))
			index++
			continue
		}

		switch content[index] {
		case ',': // unload current memory
			DebugOperation(debug, fmt.Sprintf("unloaded memory: (%d : '%c') at index: %d", main_byte[0], main_byte[0], index))
			main_byte[0] = 0
		case '.': // load referenced variable
			main_byte[0] = references[referenced].ByteValue
			DebugOperation(debug, fmt.Sprintf("loaded reference byte: (%d : '%c') from reference: \"%s\" at index: %d", main_byte[0], main_byte[0], referenced, index))
		case '!': // write main_byte memory
			_, err := os.Stdout.Write(main_byte)
			if err != nil {
				return err
			}
			DebugOperation(debug, fmt.Sprintf("wrote byte: (%d : '%c') at index: %d", main_byte[0], main_byte[0], index))
		case '+': // increment main_byte
			main_byte[0] += 1
			DebugOperation(debug, fmt.Sprintf("incremented main byte by 1 to (%d : '%c') at index: %d", main_byte[0], main_byte[0], index))
		case '-': // decrement main_byte
			main_byte[0] -= 1
			DebugOperation(debug, fmt.Sprintf("decremented main byte by 1 to (%d : '%c') at index: %d", main_byte[0], main_byte[0], index))
		case ';': // forget mentioned variable
			DebugOperation(debug, fmt.Sprintf("cleared the reference: \"%s\" at index: %d", referenced, index))
			referenced = ""
		case ':': // save mentioned variable
			references[referenced] = Reference{main_byte[0], references[referenced].GotoIndex}
			DebugOperation(debug, fmt.Sprintf("saved byte: (%d : '%c') to reference: \"%s\" at index: %d", main_byte[0], main_byte[0], referenced, index))
		case '?': // read to memory
			_, err := os.Stdin.Read(main_byte)
			if err != nil {
				return err
			}
			DebugOperation(debug, fmt.Sprintf("read: '%c' at index: %d", main_byte[0], index))
		case '/': // remove last character from mentioned variable
			DebugOperation(debug, fmt.Sprintf("subtracted last character from reference: \"%s\" to: \"%s\" at index: %d", referenced, referenced[:len(referenced)-1], index))
			if len(referenced) > 0 {
				referenced = referenced[:len(referenced)-1]
			}
		case '*': // start goto
			references[referenced] = Reference{references[referenced].ByteValue, index}
			DebugOperation(debug, fmt.Sprintf("created goto point with a label of: \"%s\" at index: %d", referenced, index))
		case '&': // goto
			DebugOperation(debug, fmt.Sprintf("jumping to goto point with a label of: \"%s\" at index: %d", referenced, index))
			index = references[referenced].GotoIndex
			DebugOperation(debug, fmt.Sprintf("jumped to index: %d", index))
			continue
		case '<': // decrease jump distance
			DebugOperation(debug, fmt.Sprintf("jumping backward by: \"%d\" at index: %d", main_byte[0], index))
			if index-int(main_byte[0]) < 0 {
				index = 0
				continue
			}
			index -= int(main_byte[0])
			DebugOperation(debug, fmt.Sprintf("jumped to index: %d", index))
			continue
		case '>': // goto right
			DebugOperation(debug, fmt.Sprintf("jumping forward by: \"%d\" at index: %d", main_byte[0], index))
			if index+int(main_byte[0]) > len(content) {
				return nil
			}
			index += int(main_byte[0])
			DebugOperation(debug, fmt.Sprintf("jumped to index: %d", index))
			continue
		case '{': // condition check
			if references[referenced].ByteValue != main_byte[0] && !condition_skip {
				DebugOperation(debug, fmt.Sprintf("condition is false (main_byte: (%d : '%c') != reference: \"%s\" = %c ) at index: %d", main_byte[0], main_byte[0], referenced, references[referenced].ByteValue, index))
				condition_skip = true
			} else {
				DebugOperation(debug, fmt.Sprintf("condition is true (main_byte: (%d : '%c') == reference: \"%s\" = %c ) at index: %d", main_byte[0], main_byte[0], referenced, references[referenced].ByteValue, index))
			}
		case '}': // condition end
			DebugOperation(debug, fmt.Sprintf("condition ended at index: %d", index))
			condition_skip = false
		default: // append to reference
			if content[index] != '~' { // do nothing / no opp
				DebugOperation(debug, fmt.Sprintf("added: '%c' to the reference: \"%s\" at index: %d", content[index], referenced, index))
				referenced += string(content[index])
			} else {
				DebugOperation(debug, fmt.Sprintf("no opperation at index: %d", index))
			}
		}
		index++
	}
	return nil
}

func Bake(content string) {

}

func Emit(content string) {

}

func Help() {
	fmt.Println("flags:\n --help : prints information about the interpereter\n --debug : debugs a deanlang script")
	os.Exit(0)
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

	arg_len := len(args) - start

	debug_mode := false
	input_file := ""
	compile_method := "run"
	fully_compile := false

	if arg_len <= 1 {
		Help()
	}

	for _, flag := range args {
		switch flag {
		case "--debug":
			debug_mode = true
		case "--bake":
			compile_method = "bake"
		case "--emit":
			compile_method = "emit"
		case "--native":
			fully_compile = true
		case "--help":
			Help()
		default:
			input_file = flag
		}
	}

	content, err := GetContent(input_file)
	if err != nil {
		log.Fatal(err)
	}
	content = Format(content)

	switch compile_method {
	case "run":
		err = Interperet(content, debug_mode)
		if err != nil {
			log.Fatal(err)
		}
	case "bake":

	}
}
