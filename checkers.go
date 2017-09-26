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
	} else if a == 7 {
		return "%"
	} else if a == 8 {
		return "0"
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

func validMove(player int, mov1 [2]int, mov2 [2]int, capOnly bool) bool {

	if ! onBoard(mov1[0], mov1[1]) || ! onBoard(mov2[0], mov2[1]){
		return false
	}

	enemy := 2
	if player == 2 {
		enemy = 1
	}


	if board[mov1[0]][mov1[1]] != player && board[mov1[0]][mov1[1]] != player + 6 {
		fmt.Println("Starting posistion invalid", board[mov1[0]][mov1[1]])
		return false
	}
	if board[mov2[0]][mov2[1]] != 3 {
		fmt.Println("Moving posistion invalid", board[mov2[0]][mov2[1]])
		return false
	}

	hoz := mov1[1] - mov2[1]

	if mov2[0] - mov1[0] == move(player, "vert") && math.Abs(float64(hoz)) == 1 && ! capOnly{
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

	vert :=  move(player, "vert")
	vertCap := move(player, "vertCap")

	if captureChecks(mov1, mov2,vert, vertCap , capOnly, enemy) {
		return true
	} else if board[mov1[0]][mov1[1]] == player + 6 && captureChecks(mov1, mov2, -vert,  -vertCap, capOnly, enemy){
		return true
	}

	return false
}

func captureChecks(mov1 [2]int, mov2 [2]int, vert int, vertCap int, capOnly bool, enemy int) bool{
	hoz := mov1[1] - mov2[1]

	if mov2[0] - mov1[0] == vert && math.Abs(float64(hoz)) == 1 && ! capOnly{
		return true
	}

	if mov2[0] - mov1[0] == vertCap {
		if hoz == -2 && onBoard(mov1[0] + vert, mov1[1] + 1){
			if board[mov1[0] + vert ][mov1[1] + 1] == enemy {
				return true
			}
		} else if hoz == 2  && onBoard(mov1[0] + vert, mov1[1] - 1) {
			if board[mov1[0] + vert ][mov1[1] - 1] == enemy {
				return true
			}
		}
	}
	return false
}

func onBoard(vert int, hoz int) bool{
	if vert > 7 || vert < 0 || hoz > 7 || hoz < 0 {
		return false
	}
	return true
}



func enemyLeft(player int) bool{
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


func playerTurn(player int, capOnly bool, lastPos [2]int){
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
	} else if strings.HasPrefix(strings.ToLower(text), "skip"){
		if capOnly{
			fmt.Println("Womp womp, you have to take the capture")
			playerTurn(player, capOnly, lastPos)
			return
		}
		fmt.Println("Skipping turn")
		return
	}

	text = string([]rune(text)[0:5])

	match, _ := regexp.MatchString("[a-h][1-8] [a-h][1-8]", text)
	if ! match{
		fmt.Printf("Bad input, expected [a-h][1-8] [a-h][1-8], got %s\n", text)
		playerTurn(player, capOnly, lastPos)
		return
	}


	pos := strings.Split(text, " ")
	mov1, err1 := posToBoard(pos[0])
	mov2, err2 := posToBoard(pos[1])
	if capOnly && mov1 != lastPos {
		fmt.Println("You have to move the same piece you just used last, numskull")
		playerTurn(player, capOnly, lastPos)
		return
	}

	if err1 || err2 {
		fmt.Println("Invalid move, buddy")
		playerTurn(player, capOnly, lastPos)
		return
	}

	if validMove(player, mov1, mov2, capOnly){
		king := false
		if board[mov1[0]][mov1[1]] == player + 6 {
			king = true
		} else if (player == 1 && mov2[0] == 0) || (player == 2 && mov2[0] == 7){
			king = true
			fmt.Println("King me!")
		}
		board[mov1[0]][mov1[1]] = 3
		if king{
			board[mov2[0]][mov2[1]] = player + 6
		} else {
			board[mov2[0]][mov2[1]] = player
		}


		if math.Abs(float64(mov2[0] - mov1[0])) == 2 { // Capture
			l := (mov1[0] + mov2[0]) / 2
			h := (mov1[1] + mov2[1]) / 2
			board[l][h] = 3
			if ! enemyLeft(player){
				playing = false
				printBoard()
				fmt.Printf("\nCongraulations Player %d, you've won!\n", player)
				return
			}
			fmt.Println("Captured!")
			if validMove(player, mov2, [2]int{mov2[0] + move(player, "vertCap"), mov2[1] + 2}, true) ||
				validMove(player, mov2, [2]int{mov2[0] + move(player, "vertCap"), mov2[1] - 2}, true){
					fmt.Println("Anoter capture possible, please go again")
					playerTurn(player, true, mov2)
			} else if king && (validMove(player, mov2, [2]int{mov2[0] - move(player, "vertCap"), mov2[1] + 2}, true) ||
				validMove(player, mov2, [2]int{mov2[0] - move(player, "vertCap"), mov2[1] - 2}, true)){
				fmt.Println("Anoter capture possible, please go again")
				playerTurn(player, true, mov2)
			}
		}

	} else {
		fmt.Println("INVALUD MOVE!")
		playerTurn(player, capOnly, lastPos)
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

	fmt.Println("Welcome to GO Checkers! ")
	fmt.Println("You play by specifiying which peice to move, and the posistion to move it too")
	fmt.Println("You can string captures, the game will let you know if you must take the next capture")
	fmt.Println("\n You can also 'skip' or 'quit'")
	fmt.Println("Have fun!")
	playing = true
	for playing {
		playerTurn(1, false, [2]int{0, 0})
		if ! playing {
			break
		}
		playerTurn(2, false, [2]int{0, 0})
	}

}
