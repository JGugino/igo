package main

import "github.com/JGugino/igo"

func main() {
	windows := igo.Windows{
		Title:  "Example window",
		Width:  800,
		Height: 600,
		FPS:    60,
	}

	windows.CreateGameWindow()
}
