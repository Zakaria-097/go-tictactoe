package game

import (
	"fmt"
	"strings"

	"github.com/zs5460/art"
)

type Position string

const (
	TL Position = "tl"
	TM          = "tm"
	TR          = "tr"
	ML          = "ml"
	MM          = "mm"
	MR          = "mr"
	BL          = "bl"
	BM          = "bm"
	BR          = "br"
)

type Player struct {
	playerTurn  bool
	Player1Move string
	Player2Move string
}

type TicTacToe struct {
	Board        [][]string
	player       *Player
	PositionsMap map[Position]*string
}

// NewTicTacToe Do something.
func NewTicTacToe() *TicTacToe {
	out := &TicTacToe{
		// create 3 by 3 tic-tac-toe board
		Board: [][]string{
			{"{_}", "{_}", "{_}"},
			{"{_}", "{_}", "{_}"},
			{"{_}", "{_}", "{_}"},
		},
		PositionsMap: make(map[Position]*string),
		player:       &Player{playerTurn: true, Player1Move: "{X}", Player2Move: "{O}"},
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

func (t *TicTacToe) Cast(m string, pos Position) {
	if !t.CheckPosition(pos) {
		t.player.playerTurn = !t.player.playerTurn
		*t.PositionsMap[pos] = m
		t.latest()
	}
}

// latest state of the board.
func (t *TicTacToe) latest() {
	for i := range t.Board {
		fmt.Printf("\n	%s\n", strings.Join(t.Board[i], " "))
	}
}

func (t *TicTacToe) CheckPosition(p Position) (occupied bool) {
	occupied = *t.PositionsMap[p] != "{_}"
	if occupied {
		fmt.Printf("\nPosition %v is Already Occupied. Try Again.\n", p)
		t.latest()
	}
	return occupied
}

func (t *TicTacToe) Play(Input1 string, Input2 string, finished bool) {
	for !finished {
		switch {
		case t.player.playerTurn: //player1's turn
			var p1_Input Position
			for p1_Input == "" {
				fmt.Printf("\n\n%v> Enter Position Code: ", Input1)
				fmt.Scanln(&p1_Input)
			}
			t.Cast(t.player.Player1Move, p1_Input)
		case !t.player.playerTurn: //player2's turn
			var p2_Input Position
			for p2_Input == "" {
				fmt.Printf("\n\n%v> Enter Position Code: ", Input2)
				fmt.Scanln(&p2_Input)
			}
			t.Cast(t.player.Player2Move, p2_Input)
		}

		firstLine := *t.PositionsMap[TL] + *t.PositionsMap[TM] + *t.PositionsMap[TR]
		secondLine := *t.PositionsMap[ML] + *t.PositionsMap[MM] + *t.PositionsMap[MR]
		thirdLine := *t.PositionsMap[BL] + *t.PositionsMap[BM] + *t.PositionsMap[BR]
		firstRow := *t.PositionsMap[TL] + *t.PositionsMap[ML] + *t.PositionsMap[BL]
		secondRow := *t.PositionsMap[TM] + *t.PositionsMap[MM] + *t.PositionsMap[BM]
		thirdRow := *t.PositionsMap[TR] + *t.PositionsMap[MR] + *t.PositionsMap[BR]
		diagonal1 := *t.PositionsMap[TL] + *t.PositionsMap[MM] + *t.PositionsMap[BR]
		diagonal2 := *t.PositionsMap[TR] + *t.PositionsMap[MM] + *t.PositionsMap[BL]

		player1Won := t.DidPlayer1Win(firstLine, secondLine, thirdLine, firstRow, secondRow, thirdRow, diagonal1, diagonal2)
		player2Won := t.DidPlayer2Win(firstLine, secondLine, thirdLine, firstRow, secondRow, thirdRow, diagonal1, diagonal2)

		switch {
		case player1Won:
			fmt.Printf("\n\n Congratulations: %v, You win!\n\n", Input1)
			finished = true
		case player2Won:
			fmt.Printf("\n\n Congratulations: %v, You win!\n\n", Input2)
			finished = true
		case t.Draw():
			fmt.Printf("\n\n Draw! Thanks for playing, %v, and %v \n\n", Input1, Input2)
			finished = true
		}
	}
}

func (t *TicTacToe) Player1WinCondition() string {
	return t.player.Player1Move + t.player.Player1Move + t.player.Player1Move
}

func (t *TicTacToe) DidPlayer1Win(conds ...string) bool {
	for _, cond := range conds {
		if cond == t.Player1WinCondition() {
			return true
		}
	}
	return false
}

func (t *TicTacToe) Draw() bool {

	var count int
	for k := range t.PositionsMap {
		if *t.PositionsMap[k] == "{X}" || *t.PositionsMap[k] == "{O}" {
			count++
		}
		if count == 9 {
			return true
		}
	}
	return false
}

func (t *TicTacToe) Player2WinCondition() string {
	return t.player.Player2Move + t.player.Player2Move + t.player.Player2Move
}

func (t *TicTacToe) DidPlayer2Win(conds ...string) bool {
	for _, cond := range conds {
		if cond == t.Player2WinCondition() {
			return true
		}
	}
	return false
}

func Start() {
	fmt.Println(art.String("\nTicTacToe\n"))

	fmt.Println("\nThis program is a basic TicTacToe Implementation in Golang")

	//inital variables
	var Input1 string
	var Input2 string
	var finished bool

	//Board Position Codes
	boardCodes := func() {
		fmt.Println("\n\n                   Board Position Codes")
		fmt.Println("\n	              {tl}", "{tm}", "{tr}")
		fmt.Println("\n	              {ml}", "{mm}", "{mr}")
		fmt.Println("\n                      {bl}", "{bm}", "{br}")
		fmt.Println()
	}

	for Input1 == "" {
		print("\nPlayer 1: Enter your name: ")
		fmt.Scanln(&Input1)
	}

	for Input2 == "" {
		print("\nPlayer 2: Enter your name: ")
		fmt.Scanln(&Input2)
	}

	boardCodes()

	fmt.Printf("\n\n		%v is X", Input1)
	fmt.Printf("		%v is O\n", Input2)

	game := NewTicTacToe()

	game.Play(Input1, Input2, finished)
}
