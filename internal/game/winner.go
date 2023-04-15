package game

import "fmt"

func (t *TicTacToe) Player1WinCondition() string {
	return t.player.Player1Move + t.player.Player1Move + t.player.Player1Move
}
func (t *TicTacToe) Player2WinCondition() string {
	return t.player.Player2Move + t.player.Player2Move + t.player.Player2Move
}

func (t *TicTacToe) DidPlayer1Win(conds ...string) bool {
	for _, cond := range conds {
		if cond == t.Player1WinCondition() {
			return true
		}
	}
	return false
}
func (t *TicTacToe) DidPlayer2Win(conds ...string) bool {
	for _, cond := range conds {
		if cond == t.Player2WinCondition() {
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
	}
	return count == 9
}

func (t *TicTacToe) HasGameFinished(player1Won, player2Won bool) bool {
	switch {
	case player1Won:
		fmt.Printf("\n\n Congratulations: %v, You win!\n\n", t.player.player1Name)
		return true
	case player2Won:
		fmt.Printf("\n\n Congratulations: %v, You win!\n\n", t.player.player2Name)
		return true
	case t.Draw():
		fmt.Printf("\n\n Draw! Thanks for playing, %v, and %v \n\n", t.player.player1Name, t.player.player2Name)
		return true
	}

	return false
}
