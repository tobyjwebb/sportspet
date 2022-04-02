package board

func New(board string) *chessBoard {
	return &chessBoard{[]byte(board)}
}

type chessBoard struct {
	status []byte
}

func (c *chessBoard) MovePiece(from, to string) {
	// TODO: validate movement
	piece := c.status[coordsToPos(from)]
	c.status[coordsToPos(to)] = piece
	c.status[coordsToPos(from)] = byte(' ')
}

func (c *chessBoard) State() string {
	return string(c.status)
}

func coordsToPos(coords string) int {
	// TODO: validate coords
	byteCoords := []byte(coords)
	letter := byteCoords[0]
	letterOffset := int(letter) - int('a')
	number := byteCoords[1]
	numberOffset := int(number) - int('1')
	return letterOffset + 8*numberOffset
}
