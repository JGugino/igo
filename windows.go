package igo

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Windows struct {
	Title  string
	Width  int
	Height int
	FPS    int
}

func (w *Windows) StartGame(clearColor color.RGBA) {
	rl.InitWindow(int32(w.Width), int32(w.Height), w.Title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(int32(w.FPS))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(clearColor)
		rl.EndDrawing()
	}
}
