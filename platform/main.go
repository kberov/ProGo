package main

import "platform/logging"

func main() {
	var l logging.Logger = logging.NewDefaultLogger(logging.Info)
	writeMessage(l)
}

func writeMessage(l logging.Logger) {
	l.Info("\n\nCreating a Web Platform\n ")
}
