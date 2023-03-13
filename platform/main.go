package main

import (
	"os"
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
	/*
		var cfg config.Configuration

		if err := services.GetService(&cfg); err != nil {
			panic(err)
		}
		var logger logging.Logger

		if err := services.GetService(&logger); err != nil {
			panic(err)
		}
	*/
	// _ is a slice of results
	if _, err := services.Call(writeMessage); err != nil {
		println(err)
		os.Exit(1)
	}

	/*
		The main function defines an anonymous struct and resolves the services it
		requires by passing a pointer to the Populate function. The result is that the
		embedded Logger fields is populated using a service.  The Populate function
		skips the message field, but a value is defined when the struct is initialized.
	*/
	val := struct {
		message string
		logging.Logger
	}{
		message: "Hello from the struct",
	}
	_ = services.Populate(&val)
	val.Logger.Debug(val.message)
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
