package pieces

const (
	None = iota
	King
	Pawn
	Knight
	Bishop
	Rook
	Queen

	White = 8
	Black = 16
)

var PieceMap = map[string]int{
	"p": Pawn | Black,
	"P": Pawn | White,
	"r": Rook | Black,
	"R": Rook | White,
	"n": Knight | Black,
	"N": Knight | White,
	"b": Bishop | Black,
	"B": Bishop | White,
	"q": Queen | Black,
	"Q": Queen | White,
	"k": King | Black,
	"K": King | White,
}
