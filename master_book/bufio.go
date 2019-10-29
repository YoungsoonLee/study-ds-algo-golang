package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		fmt.Println("> ", scan.Text())
	}
}
