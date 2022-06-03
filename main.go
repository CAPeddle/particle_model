// An implementation of Conway's Game of Life.
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"CAPeddle/particle_model/environment"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

// Game implements ebiten.Game interface.
type Game struct {
	pixels []byte
	world  *environment.Environment
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	g.world.Update()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	g.pixels = make([]byte, screenWidth*screenHeight*4)

	g.world.Draw(g.pixels)
	screen.ReplacePixels(g.pixels)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{
		world: environment.NewEnvironment(screenWidth, screenHeight),
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Game of Life (Ebiten Enhanced)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
