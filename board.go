package main

import (
	"strings"
	"strconv"
	"fmt"
)

var board [8][8] int

func PopulateNewBoard(){
	board = [8][8]int{
		{0,2,0,2,0,2,0,2},
		{2,0,2,0,2,0,2,0},
		{0,2,0,2,0,2,0,2},
		{3,0,3,0,3,0,3,0},
		{0,3,0,3,0,3,0,3},
		{1,0,1,0,1,0,1,0},
		{0,1,0,1,0,1,0,1},
		{1,0,1,0,1,0,1,0}}
}

func PrintBoard(){
	fmt.Println(BoardAsString())
}

func BoardAsString() string {
	pieces := map[int]string {
		0: "=",
		1: "X",
		2: "O",
		3: " ",
		7: "%",
		8: "0",
	}
	boardString := ""
	for i := 0; i < 8; i++ {
		boardString += strconv.Itoa(8 - i)
		for j := 0; j < 8; j++ {
			boardString += " | " + pieces[board[i][j]]
		}
		boardString += " |\n"

	}
	boardString += "    A   B   C   D   E   F   G   H "
	return boardString
}

func OnBoard(vert int, hoz int) bool{
	if vert > 7 || vert < 0 || hoz > 7 || hoz < 0 {
		return false
	}
	return true
}


func PosToBoard(a string) ([2]int, bool) {
	cvt := []rune(strings.ToLower(a))
	hoz := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
	}
	vert := map[string]int{
		"8": 0,
		"7": 1,
		"6": 2,
		"5": 3,
		"4": 4,
		"3": 5,
		"2": 6,
		"1": 7,
	}
	h, hok := hoz[strings.ToLower(string(cvt[0]))]
	v, vok := vert[strings.ToLower(string(cvt[1]))]

	err := false
	if ! hok || ! vok {
		err = true
	}

	return [2]int{v,h}, err

}


func EnemyOnBoard(player int) bool{
	enemy := 2
	if player == 2 {
		enemy = 1
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == enemy || board[i][j] == enemy + 6{
				return true
			}
		}
	}
	return false
}