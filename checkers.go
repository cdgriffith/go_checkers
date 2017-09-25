package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
	"regexp"
)
// Cant jump onto 0s
// 3s represent blank spaces you can jump too
var board [8][8] int
var playing bool


func converter(a int) string {
	if a == 0 {
		return "="
	} else if a == 1 {
		return "X"
	} else if a == 2 {
		return "O"
	} else if a == 3 {
		return " "
	} else {
		return "G"
	}
}

func posToBoard(a string) ([2]int, bool) {
	cvt := []rune(a)
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
	//fmt.Printf("%s translated to %d%d\n",a, l, h)

	err := false
	if ! hok || ! vok {
		err = true
	}

	return [2]int{v,h}, err

}


func printBoard() {
	for i := 0; i < 8; i++ {
		printString := strconv.Itoa(8 - i)
		for j := 0; j < 8; j++ {
			printString += " | " + converter(board[i][j])
		}
		printString += " |"
		fmt.Println(printString)
	}
	fmt.Println("    A   B   C   D   E   F   G   H ")

}


func move(player int, move string) int{
	if player == 2 {
		if move == "vert" {
			return 1
		} else if move == "vertCap" {
			return 2
		} else {
			return 200
		}

	} else {
		if move == "vert" {
			return -1
		} else if move == "vertCap" {
			return -2
		} else {
			return 200
		}
	}
}

func validMove(player int, mov1 [2]int, mov2 [2]int) bool {

	enemy := 2
	if player == 2 {
		enemy = 1
	}


	if board[mov1[0]][mov1[1]] != player {
		fmt.Println("Starting posistion invalid", board[mov1[0]][mov1[1]])
		return false
	}
	if board[mov2[0]][mov2[1]] != 3 {
		fmt.Println("Moving posistion invalid", board[mov2[0]][mov2[1]])
		return false
	}

	hoz := mov1[1] - mov2[1]

	if mov2[0] - mov1[0] == move(player, "vert") && math.Abs(float64(hoz)) == 1 {
		return true
	}

	if mov2[0] - mov1[0] == move(player, "vertCap")  {
		if hoz == -2 {
			if board[mov1[0] + move(player, "vert") ][mov1[1] + 1] == enemy {
				return true
			}
		} else if hoz == 2 {
			if board[mov1[0] + move(player, "vert") ][mov1[1] - 1] == enemy {
				return true
			}
		}
	}

	return false
}


func playerTurn(player int){
	printBoard()
	reader := bufio.NewReader(os.Stdin)

	rep := "X"
	if player == 2{
		rep = "O"
	}

	fmt.Printf("Player %d (%s) move (ex: a3 b4): \n", player, rep)
	text, _ := reader.ReadString('\n')
	if strings.HasPrefix(strings.ToLower(text), "quit"){
		playing = false
		fmt.Println("Quitting...")
		return
	}
	text = string([]rune(text)[0:5])

	match, _ := regexp.MatchString("[a-h][1-8] [a-h][1-8]", text)
	if ! match{
		fmt.Printf("Bad input, expected [a-h][1-8] [a-h][1-8], got %s\n", text)
		playerTurn(player)
		return
	}


	pos := strings.Split(text, " ")
	mov1, err1 := posToBoard(pos[0])
	mov2, err2 := posToBoard(pos[1])
	if err1 || err2 {
		fmt.Println("Invalid move, buddy")
		playerTurn(player)
		return
	}

	if validMove(player, mov1, mov2){
		board[mov1[0]][mov1[1]] = 3
		board[mov2[0]][mov2[1]] = player
		if mov2[0] - mov1[0] == move(player, "vertCap")  { // Capture
			l := (mov1[0] + mov2[0]) / 2
			h := (mov1[1] + mov2[1]) / 2
			board[l][h] = 3
			fmt.Println("Captured!")
		}

	} else {
		fmt.Println("INVALUD MOVE!")
		playerTurn(player)
	}


}


func main() {
	board = [8][8]int{
		{0,2,0,2,0,2,0,2},
		{2,0,2,0,2,0,2,0},
		{0,2,0,2,0,2,0,2},
		{3,0,3,0,3,0,3,0},
		{0,3,0,3,0,3,0,3},
		{1,0,1,0,1,0,1,0},
		{0,1,0,1,0,1,0,1},
		{1,0,1,0,1,0,1,0}}

	playing = true
	for playing {
		playerTurn(1)
		if ! playing {
			break
		}
		playerTurn(2)
	}

	fmt.Println("It's been fun!")

}
