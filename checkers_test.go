package main

import (
	"strings"
	"testing"
)

func PopulateTestBoard() {
	board = [8][8]int{
		{0, 2, 0, 3, 0, 3, 0, 2},
		{2, 0, 1, 0, 2, 0, 2, 0},
		{0, 2, 0, 1, 0, 2, 0, 2},
		{3, 0, 3, 0, 3, 0, 7, 0},
		{0, 8, 0, 8, 0, 8, 0, 3},
		{1, 0, 7, 0, 1, 0, 7, 0},
		{0, 1, 0, 3, 0, 1, 0, 3},
		{1, 0, 1, 0, 1, 0, 1, 0}}
}

func TestValidMoves(t *testing.T) {
	PopulateTestBoard()
	if !ValidMove(1, [2]int{2, 3}, [2]int{0, 5}, false, false) {
		t.Error("Regular 1 can't capture Regular 2")
	} else if !ValidMove(2, [2]int{1, 4}, [2]int{3, 2}, false, false) {
		t.Error("Regular 2 can't capture Regular 1")
	} else if !ValidMove(1, [2]int{5, 0}, [2]int{3, 2}, false, false) {
		t.Error("Regular 1 can't capture king 2")
	} else if !ValidMove(1, [2]int{5, 2}, [2]int{3, 0}, false, false) {
		t.Error("King 1 can't capture king 2")
	} else if !ValidMove(2, [2]int{4, 5}, [2]int{6, 7}, false, false) {
		t.Error("King 2 can't capture king 1")
	} else if !ValidMove(2, [2]int{2, 5}, [2]int{4, 7}, false, false) {
		t.Error("Regular 2 can't capture king 1")
	}

}

// Board Tests

func TestBoardAsString(t *testing.T) {
	PopulateNewBoard()
	boardString := BoardAsString()
	expected := `8 | = | O | = | O | = | O | = | O |
7 | O | = | O | = | O | = | O | = |
6 | = | O | = | O | = | O | = | O |
5 |   | = |   | = |   | = |   | = |
4 | = |   | = |   | = |   | = |   |
3 | X | = | X | = | X | = | X | = |
2 | = | X | = | X | = | X | = | X |
1 | X | = | X | = | X | = | X | = |
    A   B   C   D   E   F   G   H `

	if strings.Replace(boardString, " ", "", -1) != strings.Replace(expected, " ", "", -1) {
		t.Error("Board was not the expected string: ", boardString)
	}
}

func TestOnBoard(t *testing.T) {
	if !OnBoard(0, 0) {
		t.Error("0 0 is on board")
	} else if !OnBoard(7, 7) {
		t.Error(" 7 7 is on the board")
	} else if OnBoard(0, 8) {
		t.Error("0 8 is off the board")
	} else if OnBoard(8, 0) {
		t.Error("8 0 if off the board")
	} else if OnBoard(-1, 0) {
		t.Error("-1 0 is off the board")
	} else if OnBoard(0, -1) {
		t.Error("0 -1 if off the board")
	}
}

func TestPosToBoard(t *testing.T) {
	t1, _ := PosToBoard("a1")
	t2, _ := PosToBoard("H8")
	_, t3 := PosToBoard("j1")
	if t1 != [2]int{7, 0} {
		t.Error("a7 should equal 7,0")
	} else if t2 != [2]int{0, 7} {
		t.Error("H8 should equal 0,7")
	} else if !t3 {
		t.Error("j1 should raise an error")
	}

}

func TestEnemyOnBoard(t *testing.T) {
	PopulateNewBoard()
	if !EnemyOnBoard(1) {
		t.Error("There should be other players on the board")
	}
	board = [8][8]int{
		{0, 8, 0, 8, 0, 8, 0, 8},
		{8, 0, 8, 0, 8, 0, 8, 0},
		{0, 8, 0, 8, 0, 8, 0, 8},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0}}
	if !EnemyOnBoard(1) {
		t.Error("There should be other player's kings on the board")
	}
	board = [8][8]int{
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{3, 0, 3, 0, 3, 0, 3, 0},
		{0, 3, 0, 3, 0, 3, 0, 3},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0}}
	if EnemyOnBoard(1) {
		t.Error("Your enemy should be dead!")
	}
}
