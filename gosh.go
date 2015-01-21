package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("-> ")
		if scanner.Scan() {
			var procAttr os.ProcAttr
			procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
			process, err := os.StartProcess(scanner.Text(), []string{"-G"}, &procAttr)
			if err != nil {
				fmt.Fprintln(os.Stderr, "starting process", err)
			}
			process.Wait()
		} else {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
