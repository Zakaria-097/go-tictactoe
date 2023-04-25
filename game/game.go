package game

import (
	"fmt"
	"strings"

	"github.com/zs5460/art"
	"golang.org/x/exp/slices"
)

func Start() {
	fmt.Println(art.String("\n\ngo-TicTacToe\n\n"))

	fmt.Println("\nThis program is a basic TicTacToe Implementation in Golang")

	//inital variables
	var player1Name string
	var player2Name string
	var finished bool

	for player1Name == "" {
		print("\nPlayer 1: Enter your name: ")
		fmt.Scanln(&player1Name)
	}

	for player2Name == "" {
		print("\nPlayer 2: Enter your name: ")
		fmt.Scanln(&player2Name)
	}

	// show board position codes
	showBoardCodes()

	fmt.Printf("\n\n		%v is X", player1Name)
	fmt.Printf("		%v is O\n", player2Name)

	game := NewTicTacToe(player1Name, player2Name)

	game.Play(finished)
}

func (t *TicTacToe) Play(gameFinished bool) {
	for !gameFinished {

		player1Turn := t.player.playerTurn

		switch {
		case player1Turn: //player1's turn
			var p1_Input Position
			for p1_Input == "" {
				fmt.Printf("\n\n%v> Enter Position Code: ", t.player.player1Name)
				fmt.Scanln(&p1_Input)
			}
			t.Cast(t.player.Player1Move, p1_Input)
		case !player1Turn: //player2's turn
			var p2_Input Position
			for p2_Input == "" {
				fmt.Printf("\n\n%v> Enter Position Code: ", t.player.player2Name)
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

		didPlayer1Win := t.DidPlayer1Win(firstLine, secondLine, thirdLine, firstRow, secondRow, thirdRow, diagonal1, diagonal2)
		didPlayer2Win := t.DidPlayer2Win(firstLine, secondLine, thirdLine, firstRow, secondRow, thirdRow, diagonal1, diagonal2)

		gameFinished = t.HasGameFinished(didPlayer1Win, didPlayer2Win)
	}
}

func (t *TicTacToe) Cast(m string, pos Position) {

	switch {
	case !t.ValidPosition(pos):

		// show board codes
		showBoardCodes()

		// log error
		fmt.Printf("\n\n Unrecognized position code %q \n\n Please choose a valid board position code.", pos)

	case t.isPositionAvailable(pos):
		// cast symbol
		*t.PositionsMap[pos] = m

		// switch player turns
		t.player.playerTurn = !t.player.playerTurn

		// show board
		t.showBoardLatest()

	default:
		t.showBoardLatest()
		fmt.Printf("\nError: That position is already taken by %q. Try again.\n", strings.TrimSuffix(fmt.Sprint(*t.PositionsMap[pos])[1:], "}"))
	}
}

func (t *TicTacToe) isPositionAvailable(pos Position) bool {
	return *t.PositionsMap[pos] == "{_}"
}

func (t *TicTacToe) ValidPosition(pos Position) bool {
	return slices.Contains(t.PositionCodes, pos)
}

// showBoardLatest will display latest state of the board.
func (t *TicTacToe) showBoardLatest() {
	for i := range t.Board {
		fmt.Printf("\n\n	%s\n", strings.Join(t.Board[i], " "))
	}
}

// showBoardCodes will display board's position codes.
func showBoardCodes() {
	fmt.Println("\n\n                   Board Position Codes")
	fmt.Println("\n	              {tl}", "{tm}", "{tr}")
	fmt.Println("\n	              {ml}", "{mm}", "{mr}")
	fmt.Println("\n                      {bl}", "{bm}", "{br}")
	fmt.Println()
}
