package App

import (
	cfg "TimeTracker/pkg/Config"
	lg "TimeTracker/pkg/Loging"
	"fyne.io/fyne/v2"
	apl "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	config *cfg.Config
	loger *lg.Logger
}

func NewApp(config *cfg.Config, logger *lg.Logger, ) (App) {
	return App{
		config: config,
		loger: logger,
	}
}

func ShowApp() {
	a := apl.New()
	w := a.NewWindow("TimeTracker")
	w.Resize(fyne.NewSize(float32(cfg.GetConfig().Height), float32(cfg.GetConfig().Width)))

	testConteiner := container.NewAppTabs(
		container.NewTabItem("Tab1", widget.NewLabel("test")),
		container.NewTabItem("Tab1", widget.NewLabel("test")),
		container.NewTabItem("Tab1", widget.NewLabel("test")),
	)
	testConteiner.SetTabLocation(container.TabLocationLeading)
	w.SetContent(testConteiner)

	w.ShowAndRun()
}