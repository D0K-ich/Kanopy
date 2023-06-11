package main

import (
	lg "TimeTracker/pkg/Loging"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	logger := lg.GetLogger()
	logger.Info("Running")
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


	logger.Info("Testing")
}
