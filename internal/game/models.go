package game

// location on board
type Position string

// nought or cross
type Symbol *string

const (
	TL Position = "tl"
	TM Position = "tm"
	TR          = "tr"
	ML          = "ml"
	MM          = "mm"
	MR          = "mr"
	BL          = "bl"
	BM          = "bm"
	BR          = "br"
)

type Player struct {
	player1Name string
	player2Name string

	Player1Move string
	Player2Move string

	playerTurn bool
}

type TicTacToe struct {
	Board        [][]string
	player       *Player
	PositionsMap map[Position]Symbol
}

// NewTicTacToe will initialise the board and assign noughts and crosses to users
func NewTicTacToe(player1Name, player2Name string) *TicTacToe {

	out := &TicTacToe{
		// create 3 by 3 board
		Board: [][]string{
			{"{_}", "{_}", "{_}"},
			{"{_}", "{_}", "{_}"},
			{"{_}", "{_}", "{_}"},
		},
		PositionsMap: make(map[Position]Symbol),
		player: &Player{
			player1Name: player1Name,
			player2Name: player2Name,
			playerTurn:  true,
			Player1Move: "{X}",  // player 1 will use crosses
			Player2Move: "{O}"}, // player 2 will use noughts
	}

	out.PositionsMap[TL] = &out.Board[0][0]
	out.PositionsMap[TM] = &out.Board[0][1]
	out.PositionsMap[TR] = &out.Board[0][2]
	out.PositionsMap[ML] = &out.Board[1][0]
	out.PositionsMap[MM] = &out.Board[1][1]
	out.PositionsMap[MR] = &out.Board[1][2]
	out.PositionsMap[BL] = &out.Board[2][0]
	out.PositionsMap[BM] = &out.Board[2][1]
	out.PositionsMap[BR] = &out.Board[2][2]

	return out
}
