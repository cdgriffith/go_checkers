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
var playerChoice string
// Cant jump onto 0s
// 3s represent blank spaces you can jump too
var board [8][8] int
var playing bool


func start() {

	fmt.Println("Lets play checkers!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("X or O? ")
	text, _ := reader.ReadString('\n')

	if strings.Contains(strings.ToLower(text), "x") {
		playerChoice = "X"
	} else {
		playerChoice = "O"
	}

	fmt.Println("You are playing as", playerChoice)

}

func converter(a int) string {
	if a == 0 {
		return "="
	} else if a == 1 {
		return playerChoice
	} else if a == 2 {
		if playerChoice == "X" {
			return "O"
		} else {
			return "X"
		}
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
	lat := map[string]int{
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
	l, lok := lat[strings.ToLower(string(cvt[1]))]
	//fmt.Printf("%s translated to %d%d\n",a, l, h)

	err := false
	if ! hok || ! lok {
		err = true
	}

	return [2]int{l,h}, err

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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func playerTurn(){
	printBoard()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your move (ex: a3 b4): ")
	text, _ := reader.ReadString('\n')
	if stringInSlice(strings.ToLower(text), []string{"quit", "exit", "stop"}){
		playing = false
		fmt.Println("Quitting...")
		return
	}
	text = string([]rune(text)[0:5])

	match, _ := regexp.MatchString("[a-h][1-8] [a-h][1-8]", text)
	if ! match{
		fmt.Printf("Bad input, expected [a-h][1-8] [a-h][1-8], got %s\n", text)
		playerTurn()
		return
	}


	pos := strings.Split(text, " ")
	mov1, err1 := posToBoard(pos[0])
	mov2, err2 := posToBoard(pos[1])
	if err1 || err2 {
		fmt.Println("Invalid move, buddy")
		playerTurn()
		return
	}

	if validMove(1, mov1, mov2){
		board[mov1[0]][mov1[1]] = 3
		board[mov2[0]][mov2[1]] = 1
		if mov2[0] - mov1[0] == -2  { // Capture
			l := (mov1[0] + mov2[0]) / 2
			h := (mov1[1] + mov2[1]) / 2
			fmt.Println(board[l][h])
			board[l][h] = 3
		}

	} else {
		fmt.Println("INVALUD MOVE!")
		playerTurn()
	}


}

func compTurn(){

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

	start()
	if playerChoice == "O" {
		compTurn()
	}

	playing = true
	for playing {
		playerTurn()
		if ! playing {
			break
		}
		compTurn()
	}

	fmt.Println("It's been fun!")

}
