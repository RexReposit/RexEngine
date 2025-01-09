package src

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	screenWidth, screenHeight int32
	windowTitle               string
}

func NewApp(w, h int32, wTitle string) *App {
	return &App{screenWidth: w, screenHeight: h, windowTitle: wTitle}
}

func (a *App) Run(node *Node) {
	rl.InitWindow(a.screenWidth, a.screenHeight, node.Name)
	rl.SetTargetFPS(60)

	rl.InitAudioDevice()

	node.Init()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		node.Process()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
