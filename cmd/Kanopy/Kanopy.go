package main

import (
	//app "TimeTracker/pkg/App"
	cfg "Kanopy/pkg/Config"
	lg "Kanopy/pkg/Loging"
	mv "Kanopy/pkg/Movie"
)

func main() {
	// Start logger
	logger := lg.GetLogger()
	cfg := cfg.GetConfigApp()

	logger.Info("Start main application")
	// Initialize configuration
	//app.ShowApp()
	logger.Info("End without problem", cfg.Listen.Port)
	mv.GetTest()
}
