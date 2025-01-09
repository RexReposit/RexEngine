#Minimal sample

```Go
package main

import (
	engine "github.com/RexReposit/RexEngine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	engine.Node
	Position rl.Vector2
}

func (p *Player) Init() {
	p.Position = rl.Vector2{
		X: 10,
		Y: 10,
	}
}

func (p *Player) Process() {
	p.Position.X += 0.1
}

func main() {
	app := engine.NewApp(1366, 768, "MyApp")
	rootNode := engine.NewRootNode()

	player := Player{
		Position: rl.Vector2{},
	}

	rootNode.AddChild(&player)

	app.Run(rootNode)
}

```
