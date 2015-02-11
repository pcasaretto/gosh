package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-shellwords"
	"os"
	"reflect"
)

type BuiltinCollection int

func (c BuiltinCollection) Echo(in []string) {
	fmt.Println(in)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var builtins BuiltinCollection
	for {
		fmt.Print("-> ")
		if scanner.Scan() {
			var procAttr os.ProcAttr
			procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
			args, err := shellwords.Parse(scanner.Text())
			if err != nil {
				fmt.Fprintln(os.Stderr, "reading command line args", err)
				continue
			}
			command := args[0]
			value := reflect.ValueOf(builtins).MethodByName(command)
			if value.IsValid() {
				value.Call([]reflect.Value{reflect.ValueOf(args)})
			} else {
				process, err := os.StartProcess(args[0], args, &procAttr)
				if err != nil {
					fmt.Fprintln(os.Stderr, "starting process", err)
					continue
				}
				process.Wait()
			}
		} else {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
