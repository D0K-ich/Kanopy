package main

import (
	app "TimeTracker/pkg/App"
	cfg "TimeTracker/pkg/Config"
	lg "TimeTracker/pkg/Loging"
)

func main() {
	// Start logger
	logger := lg.GetLogger()
	cfg := cfg.GetConfigApp()

	logger.Info("Start main application")
	// Initialize configuration
	app.ShowApp()
	logger.Info("End without problem", cfg.Listen.Port)
}
