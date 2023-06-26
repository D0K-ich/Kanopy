package App

import (
	cfg "Kanopy/pkg/Config"
	lg "Kanopy/pkg/Loging"
	th "Kanopy/pkg/Themes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	config *cfg.ConfigApp
	loger  *lg.Logger
}

func NewApp(config *cfg.ConfigApp, logger *lg.Logger) App {
	return App{
		config: config,
		loger:  logger,
	}
}

func ShowApp() {

	a := app.New()
	a.Settings().SetTheme(th.GetMyTheme())
	w := a.NewWindow("TimeTracker")
	w.Resize(fyne.NewSize(float32(cfg.GetConfigApp().Height), float32(cfg.GetConfigApp().Width)))

	testConteiner := container.NewAppTabs(
		container.NewTabItem("Tab1", widget.NewLabel("test")),
		container.NewTabItem("Tab2", widget.NewLabel("test")),
		container.NewTabItem("Tab3", widget.NewLabel("test")),
		container.NewTabItem("Tab3", widget.NewLabel("test2")),
	)
	testConteiner.SetTabLocation(container.TabLocationLeading)

	w.SetContent(testConteiner)

	w.ShowAndRun()
}
