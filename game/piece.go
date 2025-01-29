package game

type piece struct {
	sign  string
	taken bool
}

var pieces = make(map[int]*piece)

func newPiece(id int, sign string, taken bool) {
	pieces[id] = &piece{
		sign:  sign,
		taken: taken,
	}
}

func createPieces() {
	newPiece(1, "X", false)
	newPiece(2, "O", false)
	newPiece(3, "@", false)
	newPiece(4, "#", false)
}
