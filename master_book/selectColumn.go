package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("useage: selectColumn column <file1> [<file2> [...<fileN]]\n")
		os.Exit(1)
	}

	temp, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Column value is not an integer: ", temp)
		return
	}

	column := temp
	if column < 0 {
		fmt.Println("Invalied Column number!")
		os.Exit(1)
	}

	for _, filename := range args[2:] {
		fmt.Println("\n\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}

		}
	}

}
