package main

import (
	"log/slog"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jhuggett/thingamabob/app"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

	game := app.NewApp(func(app *app.App) {
		ebiten.SetWindowSize(1200, 800)
		ebiten.SetWindowTitle("Design Library Example")
		ebiten.SetWindowDecorated(true)
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

		ebiten.SetCursorMode(ebiten.CursorModeVisible)

		// firstPage := NewFirstPage(
		// 	app,
		// )
		// app.Push(firstPage)

		sandboxPage := NewSandboxPage(
			app,
		)
		app.Push(sandboxPage)

		if err := ebiten.RunGame(app); err != nil {
			panic(err)
		}
	})
	game.Start()
}
