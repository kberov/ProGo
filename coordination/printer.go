package main

import . "fmt"

func Printfln(template string, values ...any) {
	Printf(template+"\n", values...)
}
