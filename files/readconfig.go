package main

import (
	"encoding/json"
	"log"
	"os"
	//"strings"
)

type ConfigData struct {
	UserName           string
	AdditionalProducts []Product
}

var Config ConfigData

func init() {
	err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s.", err.Error())
	} else {
		Printfln("UserName: %s", Config.UserName)
		Products = append(Products, Config.AdditionalProducts...)
	}
}

func LoadConfig() (err error) {
	file, err := os.Open("config.json")
	if err == nil {
		defer func() { _ = file.Close() }()
		nameSlice := make([]byte, 5)
		_, _ = file.ReadAt(nameSlice, 15)
		Config.UserName = string(nameSlice)
		_, _ = file.Seek(45, 0)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config.AdditionalProducts) // this
		// Printfln("%s", data)
		// is the same as
		// Printfln(string(data))
		// SO!  String is just a sequence(slice) of bytes which is
		// interpreted/treated as string. Each byte or subsequence of bytes is
		// associated with a letter in the Unicode table.
	}
	return
}
