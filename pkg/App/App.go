package App

import (
	cfg "Kanopy/pkg/Config"
	lg "Kanopy/pkg/Loging"
	th "Kanopy/pkg/Themes"
	"fmt"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/canvas"
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

	application := app.New()
	application.Settings().SetTheme(th.GetMyTheme())
	w := application.NewWindow("Kanopy")
	w.Resize(fyne.NewSize(float32(cfg.GetConfigApp().Height), float32(cfg.GetConfigApp().Width)))

	

	movies := widget.NewButton("Фильмы", func ()  {fmt.Println(123)})
	music := widget.NewButton("Музыка", func ()  {fmt.Println(123)})
	anime := widget.NewButton("Аниме", func ()  {fmt.Println(123)})
	books := widget.NewButton("Книги", func ()  {fmt.Println(123)})
	manga := widget.NewButton("Манга", func ()  {fmt.Println(123)})

	t := canvas.NewRectangle(color.RGBA{29, 11, 80, 255})
	t.Resize(fyne.NewSize(200, 800))
	t.Move(fyne.NewPos(-10, -10))

	leftmenu := container.NewWithoutLayout(
		movies,
		music,
		anime,
		books,
		manga,
		//t,
	)

	almen := container.NewWithoutLayout(
		t, 
		leftmenu,
	)

	movies.Move(fyne.NewPos(49, 200))
	movies.Resize(fyne.NewSize(70, 30))
	music.Move(fyne.NewPos(49, 250))
	music.Resize(fyne.NewSize(70, 30))
	anime.Move(fyne.NewPos(49, 300))
	anime.Resize(fyne.NewSize(70, 30))
	books.Move(fyne.NewPos(49, 350))
	books.Resize(fyne.NewSize(70, 30))
	manga.Move(fyne.NewPos(49, 400))
	manga.Resize(fyne.NewSize(70, 30))


	w.SetContent(almen)
	

	w.ShowAndRun()
}
