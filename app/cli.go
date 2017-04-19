package app

import (
	"fmt"

	"statseventrelay/config"
	"statseventrelay/log"

	"github.com/spf13/cobra"
)

// Custom config file path
var configPath string

// Root CLI Command
var rootCliCmd = &cobra.Command{
	Use:   "statseventrelay",
	Short: "Consumes player events, augmenting and  placing them onto a work queue",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Bind Log Level Flag to Log Configuration
		config.BindLogLevelFlag(cmd.PersistentFlags().Lookup("log-level"))
		// Load custom configuration file if path given
		if configPath != "" {
			if err := config.Read(configPath); err != nil {
				log.WithError(err).Error("error reading configuration")
			}
		}
		Setup() // Run Setup
	},
	Run: func(*cobra.Command, []string) {
		if err := run(); err != nil {
			log.WithError(err).Error("application run error")
		}
	},
}

// Version CLI Command
var versionCliCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version and build time",
	Run: func(*cobra.Command, []string) {
		bt := BuildTime().Format("Monday January 2 2006 at 15:04:05 MST")
		fmt.Println(fmt.Sprintf("Version: %s", Version()))
		fmt.Println(fmt.Sprintf("Build Time: %s", bt))
	},
}

func init() {
	// Add CLI Flags
	rootCliCmd.PersistentFlags().StringVarP(
		&configPath,
		"config",
		"c",
		"",
		"Absolute path to configuration file")
	rootCliCmd.PersistentFlags().StringP(
		"log-level",
		"l",
		"",
		"Log Level (debug,info,warn,error)")
	// Add sub commands
	rootCliCmd.AddCommand(versionCliCmd)
}
