package battles

type Battle struct {
	ID          string
	Board       string
	WhiteTeamID string
	BlackTeamID string
	MoveCount   int // With this, can tell turn
}

// type Move struct {
// 	Num int
// 	BattleID string
// 	SessionID string // ?
//  From Position
//  To Position
// 	Nick string // ?
// 	Piece byte
// 	Timestamp string
// }

// type Position struct {
// 	X int
// 	Y int
// }
