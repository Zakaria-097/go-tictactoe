package game_test

import (
	"testing"

	"github.com/Zakaria-097/go-tictactoe/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestNewTicTacToe(t *testing.T) {
	got := game.NewTicTacToe()
	expected := &game.TicTacToe{
		Board: [][]string{
			{"{_}", "{_}", "{_}"},
			{"{_}", "{_}", "{_}"},
			{"{_}", "{_}", "{_}"},
		},
	}
	assert.Equal(t, expected, got)
}

func TestDidPlayer1Win(t *testing.T) {
	tcs := map[string]struct {
		conds    string
		board    [][]string
		expected bool
	}{
		"test1": {
			conds:    "{X}{X}{X}",
			expected: true,
		},
		"test2": {
			conds:    "{O}{O}{X}",
			expected: false,
		},
		"test3": {
			conds:    "{X}{O}{X}",
			expected: false,
		},
	}
	for name, tc := range tcs {
		t.Run(name,
			func(t *testing.T) {
				game := game.NewTicTacToe()
				got := game.DidPlayer1Win(tc.conds)
				assert.Equal(t, tc.expected, got)
			})
	}
}
func TestDidPlayer2Win(t *testing.T) {
	tcs := map[string]struct {
		conds    string
		board    [][]string
		expected bool
	}{
		"test1": {
			conds:    "{O}{O}{O}",
			expected: true,
		},
		"test2": {
			conds:    "{X}{X}{O}",
			expected: false,
		},
		"test3": {
			conds:    "{O}{X}{O}",
			expected: false,
		},
	}
	for name, tc := range tcs {
		t.Run(name,
			func(t *testing.T) {
				game := game.NewTicTacToe()
				got := game.DidPlayer2Win(tc.conds)
				assert.Equal(t, tc.expected, got)
			})
	}
}

func TestCast(t *testing.T) {
	tcs := map[string]struct {
		pos   game.Position
		move  string
		board [][]string
	}{
		"p1 tl": {
			pos:  "tl",
			move: "{X}",
			board: [][]string{
				{"{X}", "{_}", "{_}"},
				{"{_}", "{_}", "{_}"},
				{"{_}", "{_}", "{_}"},
			},
		},
		"p2 tl": {
			pos:  "tl",
			move: "{O}",
			board: [][]string{
				{"{O}", "{_}", "{_}"},
				{"{_}", "{_}", "{_}"},
				{"{_}", "{_}", "{_}"},
			},
		},
		"p1 tm": {
			pos:  "tm",
			move: "{X}",
			board: [][]string{
				{"{_}", "{X}", "{_}"},
				{"{_}", "{_}", "{_}"},
				{"{_}", "{_}", "{_}"},
			},
		},
		"p2 tm": {
			pos:  "tm",
			move: "{O}",
			board: [][]string{
				{"{_}", "{O}", "{_}"},
				{"{_}", "{_}", "{_}"},
				{"{_}", "{_}", "{_}"},
			},
		},
		"p1 tr": {
			pos:  "tr",
			move: "{X}",
			board: [][]string{
				{"{_}", "{_}", "{X}"},
				{"{_}", "{_}", "{_}"},
				{"{_}", "{_}", "{_}"},
			},
		},
		"p2 tr": {
			pos:  "tr",
			move: "{O}",
			board: [][]string{
				{"{_}", "{_}", "{O}"},
				{"{_}", "{_}", "{_}"},
				{"{_}", "{_}", "{_}"},
			},
		},
	}
	for name, tc := range tcs {
		t.Run(name,
			func(t *testing.T) {
				game := game.NewTicTacToe()
				game.Cast(tc.move, tc.pos)
				got := game.Board
				assert.Equal(t, tc.board, got)
			})
	}
}

func TestDraw(t *testing.T) {
	tcs := map[string]struct {
		Position1 game.Position
		Position2 game.Position
		Value1    string
		Value2    string
		expected  bool
	}{
		"Not a DRaw": {
			Position1: "bm",
			Position2: "br",
			Value1:    "{O}",
			Value2:    "{X}",
			expected:  false, //p1 won Draw=false
		},

		"A Draw": {
			Position1: "bm",
			Position2: "br",
			Value1:    "{O}",
			Value2:    "{X}",
			expected:  true, //draw = true
		},
	}
	for name, tc := range tcs {
		t.Run(name,
			func(t *testing.T) {
				g := game.NewTicTacToe()
				*g.PositionsMap["tl"] = "{X}"
				*g.PositionsMap["tm"] = "{O}"
				*g.PositionsMap["tr"] = "{X}"
				*g.PositionsMap["ml"] = "{O}"
				*g.PositionsMap["mm"] = "{X}"
				*g.PositionsMap["mr"] = "{O}"
				*g.PositionsMap["bl"] = "{X}"
				*g.PositionsMap[tc.Position1] = tc.Value1
				*g.PositionsMap[tc.Position2] = tc.Value2
				print(g.Board)
				got := g.Draw()
				assert.Equal(t, tc.expected, got)
			})
	}
}

func TestCheckPosition(t *testing.T) {
	tcs := map[string]struct {
		test_Positon game.Position
		value        string
		expected     bool
	}{
		"tl occupied1": {
			test_Positon: "tl",
			value:        "{X}",
			expected:     true,
		},
		"tl occupied2": {
			test_Positon: "tl",
			value:        "{O}",
			expected:     true,
		},
		"tl free": {
			test_Positon: "tl",
			value:        "{_}",
			expected:     false,
		},
		"tm occupied1": {
			test_Positon: "tm",
			value:        "{X}",
			expected:     true,
		},
		"tm occupied2": {
			test_Positon: "tm",
			value:        "{O}",
			expected:     true,
		},
		"tm free": {
			test_Positon: "tm",
			value:        "{_}",
			expected:     false,
		},
	}
	for name, tc := range tcs {
		t.Run(name,
			func(t *testing.T) {
				g := game.NewTicTacToe()
				*g.PositionsMap[tc.test_Positon] = tc.value
				got := g.CheckPosition(tc.test_Positon)
				assert.Equal(t, tc.expected, got)
			})
	}
}

func TestPlay(t *testing.T) {
	//t.Run("playTL", game.playTL("X", "player1", "player1"))
}
