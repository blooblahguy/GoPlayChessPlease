package main

import (
	"fmt"
	"gpcp/pieces"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var blackPawn = "images/black-pawn.png"
var whitePawn = "images/white-pawn.png"
var blackRook = "images/black-rook.png"
var whiteRook = "images/white-rook.png"
var blackKnight = "images/black-knight.png"
var whiteKnight = "images/white-knight.png"
var blackBishop = "images/black-bishop.png"
var whiteBishop = "images/white-bishop.png"
var blackQueen = "images/black-queen.png"
var whiteQueen = "images/white-queen.png"
var blackKing = "images/black-king.png"
var whiteKing = "images/white-king.png"

var square = ebiten.NewImage(squareSize, squareSize)

type HighlightImage struct {
	image     *ebiten.Image
	positionX int
	positionY int
	enabled   bool
}

var Highlight = HighlightImage{image: ebiten.NewImage(squareSize, squareSize)}

func BlendColors(color1 color.RGBA, color2 color.RGBA, alpha float64) color.RGBA {
	// Clamp alpha to the range [0.0, 1.0]
	if alpha < 0.0 {
		alpha = 0.0
	} else if alpha > 1.0 {
		alpha = 1.0
	}

	// Blend each color channel
	blend := func(c1, c2 uint8) uint8 {
		return uint8(float64(c1)*(1-alpha) + float64(c2)*alpha)
	}

	return color.RGBA{
		R: blend(color1.R, color2.R),
		G: blend(color1.G, color2.G),
		B: blend(color1.B, color2.B),
		A: blend(color1.A, color2.A),
	}
}

func OpenPNG(filePath string) (image.Image, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after reading

	// Decode the PNG file
	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode PNG: %w", err)
	}

	return img, nil
}

var PieceAssets = map[int]*ebiten.Image{
	pieces.Pawn | pieces.Black:   load(blackPawn),
	pieces.Pawn | pieces.White:   load(whitePawn),
	pieces.Rook | pieces.Black:   load(blackRook),
	pieces.Rook | pieces.White:   load(whiteRook),
	pieces.Knight | pieces.Black: load(blackKnight),
	pieces.Knight | pieces.White: load(whiteKnight),
	pieces.Bishop | pieces.Black: load(blackBishop),
	pieces.Bishop | pieces.White: load(whiteBishop),
	pieces.Queen | pieces.Black:  load(blackQueen),
	pieces.Queen | pieces.White:  load(whiteQueen),
	pieces.King | pieces.Black:   load(blackKing),
	pieces.King | pieces.White:   load(whiteKing),
}

func load(file string) *ebiten.Image {
	image, err := OpenPNG(file)
	if err != nil {
		fmt.Println(file)
		panic(err)
	}
	img := ebiten.NewImageFromImage(image)

	return img
}
