package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var squareSize int = 140
var pieceSize float64 = 128
var lightColor = color.RGBA{240, 240, 240, 255}
var darkColor = color.RGBA{20, 150, 75, 255}
var highlightColor = color.RGBA{0, 0, 0, 255}
var board = &Board{}

// squares
var whiteSquare = ebiten.NewImage(squareSize, squareSize)
var darkSquare = ebiten.NewImage(squareSize, squareSize)

// pieces

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// Get the current mouse position
		mouseX, mouseY := ebiten.CursorPosition()

		file := mouseX / squareSize
		rankRaw := mouseY / squareSize
		rank := int(math.Abs(float64(rankRaw) - 7))

		Highlight.enabled = true
		Highlight.positionX = file * squareSize
		Highlight.positionY = rankRaw * squareSize

		// rank, file := board.GetRankAndFileFromIndex(k)

		if file >= 0 && file < 8 && rank >= 0 && rank < 8 {
			fmt.Printf("mouse clicked %d, %d", mouseX, mouseY)
			fmt.Println()
			fmt.Printf("clicked on file: %d and rank: %d", file, rank)
			fmt.Println()

			// index := rank*8+file

			// board.Squares[index]
		}
	}
	// Write your game's logical update.
	// fmt.Println("update")
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")

	CreateGraphicalBoard(screen)
	// fmt.Println("darw")
	// Write your game's rendering.
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 8 * squareSize, 8 * squareSize
}

func main() {
	// create assets
	whiteSquare.Fill(lightColor)
	darkSquare.Fill(darkColor)

	Highlight.image.Fill(highlightColor)

	// pieces
	board.LoadPositionFromFEN(StartFEN)

	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(8*squareSize, 8*squareSize)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func CreateGraphicalBoard(screen *ebiten.Image) {
	for file := 0; file < 8; file++ {
		for rank := 0; rank < 8; rank++ {
			var isLightSquare bool = (file+rank)%2 != 0

			positionX := float64(file * squareSize)
			positionY := float64(rank * squareSize)

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(positionX, positionY)

			if isLightSquare {
				square.Fill(lightColor)
				screen.DrawImage(square, opts)
			} else {
				square.Fill(darkColor)
				screen.DrawImage(square, opts)
			}
		}
	}

	// draw highlight square
	if Highlight.enabled {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(Highlight.positionX), float64(Highlight.positionY))
		rank := Highlight.positionX / squareSize
		file := Highlight.positionY / squareSize
		var isLightSquare bool = (rank+file)%2 != 0
		// fmt.Println(Highlight.positionX, Highlight.positionY)
		if isLightSquare {
			Highlight.image.Fill(BlendColors(lightColor, highlightColor, 0.1))
		} else {
			Highlight.image.Fill(BlendColors(darkColor, highlightColor, 0.5))
		}

		screen.DrawImage(Highlight.image, opts)
	}

	// now let's populate pieces
	for k, v := range board.Squares {
		image := PieceAssets[v]
		rank, file := board.GetRankAndFileFromIndex(k)

		positionX := float64(file) * float64(squareSize)
		positionY := float64(rank) * float64(squareSize)

		scale := pieceSize / 128

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(scale, scale)
		opts.GeoM.Translate(positionX, positionY)

		if image != nil {
			screen.DrawImage(image, opts)
		}
	}

}
