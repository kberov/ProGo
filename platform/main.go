package main

import (
	"os"
	"platform/config"
	"platform/logging"
)

func main() {
	var cfg config.Configuration
	var err error
	cfg, err = config.Load("config.json")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	var logger logging.Logger = logging.NewDefaultLogger(cfg)
	writeMessage(logger, cfg)
}

func writeMessage(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if ok {
		message, ok := section.GetString("message")
		if ok {
			logger.Info(message)
		} else {
			logger.Panic("Cannot find configuration setting")
		}
	} else {
		logger.Panic("Config section not found")
	}
}
