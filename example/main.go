package main

func main() {
	igo := Igo
	window := igo.{
		Title:  "Example Game",
		Width:  800,
		Height: 600,
		FPS:    60,
	}

	window.CreateGameWindow()
}
