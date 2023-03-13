package main

import (
	"platform/config"
	"platform/logging"
	"platform/services"
)

func main() {
	// Registering and Using Services
	services.RegisterDefaultServices()

	/*
		Resolving a service is done by passing a pointer to a variable whose type is an
		interface. In Listing 32-21, the GetService function is used to obtain
		implementations of the Repository and Logger interfaces, without needing to
		know which struct type will be used, the process by which it is created, or the
		service lifecycles.
	*/
	var cfg config.Configuration

	if err := services.GetService(&cfg); err != nil {
		panic(err)
	}
	var logger logging.Logger

	if err := services.GetService(&logger); err != nil {
		panic(err)
	}
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
