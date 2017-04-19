package app

import (
	"os"
	"os/signal"
	"syscall"

	"statseventrelay/config"
	"statseventrelay/log"

	"github.com/sirupsen/logrus"
)

// Execute application
func Execute() error {
	return rootCliCmd.Execute()
}

// Setup
func Setup() {
	// Create a new logger
	l := log.New(log.Config{
		Level:        config.LogLevel(),
		Format:       config.LogFormat(),
		File:         config.LogFile(),
		Console:      config.LogConsole(),
		LogstashType: config.LogstashType(),
	})
	// Always log these fields
	l.PersistentFields(logrus.Fields{
		"version":   Version(),
		"buildTime": BuildTime(),
	})
	// Set global logger
	log.SetGlobalLogger(l)
}

// Wait for OS signal
func wait() os.Signal {
	sigC := make(chan os.Signal, 1)
	signal.Notify(
		sigC,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	signal := <-sigC
	return signal
}

// Runs the application
func run() error {
	log.Info("application start")
	defer log.Info("application exit")
	// TODO: Stuff
	// Wait OS signal to exit
	signal := wait()
	log.WithField("signal", signal).Debug("received os signal")
	return nil
}
