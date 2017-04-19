package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"statseventrelay/config"
	"statseventrelay/log"
	"statseventrelay/pubsub/redis"

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
	// Context
	ctx, cancel := context.WithCancel(context.Background())
	// Pub/Sub Subscribtions
	ps := redis.New(redis.Config{
		Host: config.RedisHost(),
	})
	defer ps.Close()
	_, err := ps.Subscribe(ctx, "player:play", "player:stop")
	if err != nil {
		return err
	}
	// Wait OS signal to exit
	signal := wait()
	log.WithField("signal", signal).Debug("received os signal")
	cancel()
	return nil
}
