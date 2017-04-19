package main

import (
	"statseventrelay/app"
	"statseventrelay/log"
)

func main() {
	if err := app.Execute(); err != nil {
		log.WithError(err).Error("error starting application")
	}
}
