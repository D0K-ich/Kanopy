package main

import (
	cfg "TimeTracker/pkg/Config"
	lg "TimeTracker/pkg/Loging"
	app "TimeTracker/pkg/App"
)

func main() {
	// Start logger
	logger := lg.GetLogger()
	cfg := cfg.GetConfig()

	logger.Info("Start main application")
	// Initialize configuration
	app.ShowApp()
	logger.Info("End without problem", cfg.Listen.Port)

}
