package main

import (
	"gpcp/pieces"
	"math"
	"strconv"
	"strings"
)

var StartFEN string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

type Board struct {
	Squares [64]int
}

func (b *Board) SetBoard() {
	b.Squares[0] = pieces.White | pieces.Bishop
	b.Squares[63] = pieces.Black | pieces.Queen
	b.Squares[7] = pieces.Black | pieces.Knight
}

func (b *Board) LoadPositionFromFEN(fen string) {
	// pieceTypeFromSymbol := map[string]int{"k": pieces.King, "p": pieces.Pawn, "n": pieces.Knight, "b": pieces.Bishop, "r": pieces.Rook, "q": pieces.Queen}

	file := 0
	rank := 7

	symbols := strings.Split(fen, "")

	for _, symbol := range symbols {
		if symbol == "/" {
			file = 0
			rank--
		} else {
			if skipValue, err := strconv.Atoi(symbol); err == nil {
				file += skipValue
			} else {
				// var pieceColor int
				// if strings.ToUpper(symbol) == symbol {
				// 	pieceColor = pieces.White
				// } else {
				// 	pieceColor = pieces.Black
				// }

				// var pieceType = pieceTypeFromSymbol[strings.ToLower(symbol)]
				b.Squares[rank*8+file] = pieces.PieceMap[symbol]

				file++
			}
		}
	}
}

func (b *Board) GetRankAndFileFromIndex(index int) (int, int) {
	rank := math.Abs(float64(index/8) - 7) // flip the biard
	file := index % 8

	return int(rank), file
}

// func (b *Board) GetPieceCodeFromIndex(index int) int{

// }
