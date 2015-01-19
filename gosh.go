package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}
		process, err := os.StartProcess(scanner.Text(), nil, &procAttr)
		if err != nil {
			continue
		}
		process.Wait()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
