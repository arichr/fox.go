package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Cat written in Go.")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open %s: %s\n", os.Args[1], err)
		return
	}

	reader := bufio.NewReader(file)
	filesize := reader.Size()

	lines := 1
	line_ending := "None"
	for buffered := reader.Size(); buffered > 0; buffered-- {
		read_byte, err := reader.ReadByte()
		if err != nil {
			break
		}

		if read_byte == '\n' {
			line_ending = "LF (Unix)"
			lines += 1
		} else if read_byte == '\r' {
			next_byte, _ := reader.ReadByte()
			if next_byte == '\n' {
				line_ending = "CRLF (Windows)"
			} else {
				line_ending = "CR (Apple)"
			}
			lines += 1
		}
	}

	fmt.Printf("> \033[43;1mFile size:\033[49m %d bytes (%dK)", filesize, filesize/1024)
	fmt.Print("    \033[44;1mLines:\033[49m\t", lines)
	fmt.Printf("    \033[45;1mLine ending:\033[49m %s\033[0m <\n", line_ending)

	file.Close()

	file, err = os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open %s: %s\n", os.Args[1], err)
		return
	}

	scanner := bufio.NewScanner(file)
	line := 1

	line_format := fmt.Sprintf("\033[40;1m%%%dd\033[0m ", len(strconv.Itoa(lines)))
	for ; scanner.Scan(); line++ {
		fmt.Printf(line_format, line)
		fmt.Println(scanner.Text())
	}
	if lines+1 > line {
		fmt.Printf(line_format+"\n", line)
	}
}
