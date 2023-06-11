package main

import (
	lg "TimeTracker/pkg/Loging"
	cfg "TimeTracker/pkg/Config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Start logger
	logger := lg.GetLogger()
	logger.Info("Start main application")
	// Initialize configuration
	cfg := cfg.GetConfig()
	
	
	a := app.New()
	w := a.NewWindow("TimeTracker")
	w.Resize(fyne.NewSize(800, 600))

	testConteiner := container.NewAppTabs(
		container.NewTabItem("Tab1", widget.NewLabel("test")),
		container.NewTabItem("Tab1", widget.NewLabel("test")),
		container.NewTabItem("Tab1", widget.NewLabel("test")),
	)
	testConteiner.SetTabLocation(container.TabLocationLeading)
	w.SetContent(testConteiner)

	w.ShowAndRun()

	logger.Info("End without problem", cfg.Listen.Port)

}
