package main

import "fmt"

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}
