# Game object
```Go
package game

import (
	"fmt"

	engine "github.com/RexReposit/RexEngine"
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	engine.Node // implplementation engine.Node interface
	Position    raylib.Vector2
}

// once initialized
func (p *Player) Init() {
	p.Position = raylib.Vector2{
		X: 10,
		Y: 10,
	}
}

// every frame
func (p *Player) Process() {
	p.Position.X += 0.1
	fmt.Println(p.Position)
}
```








# Init sample

```Go
package main

import (
	"app/game"

	engine "github.com/RexReposit/RexEngine"
)

func main() {
	app := engine.NewApp(1366, 768, "MyApp") // 1. Create a new game engine application

	rootNode := engine.NewRootNode()  // 2. Create a root node for the game
	rootNode.AddChild(&game.Player{}) // 3. Add a objects to the root node

	app.Run(rootNode) // 4. Run the game
}

```
