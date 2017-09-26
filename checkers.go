package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math"
	"regexp"
)

var playing bool

func getPossiblePos(player int) (int, int) {
	vert := -1
	if player == 2 {
		vert = 1
	}
	vertCap := -2
	if player == 2 {
		vertCap = 2
	}
	return vert, vertCap
}

func validMove(player int, mov1 [2]int, mov2 [2]int, capOnly bool) bool {

	if ! OnBoard(mov1[0], mov1[1]) || ! OnBoard(mov2[0], mov2[1]) {
		return false
	}

	enemy := 2
	if player == 2 {
		enemy = 1
	}

	if board[mov1[0]][mov1[1]] != player && board[mov1[0]][mov1[1]] != player+6 {
		fmt.Println("Starting posistion invalid", board[mov1[0]][mov1[1]])
		return false
	}
	if board[mov2[0]][mov2[1]] != 3 {
		fmt.Println("Moving posistion invalid", board[mov2[0]][mov2[1]])
		return false
	}

	if captureChecks(player, mov1, mov2, capOnly, enemy, false) {
		return true
	} else if board[mov1[0]][mov1[1]] == player+6 && captureChecks(player, mov1, mov2, capOnly, enemy, true) {
		return true
	}

	return false
}

func captureChecks(player int, mov1 [2]int, mov2 [2]int, capOnly bool, enemy int, kingFlip bool) bool {
	hoz := mov1[1] - mov2[1]
	vert, vertCap := getPossiblePos(player)
	if kingFlip {
		vert = -vert
		vertCap = -vertCap
	}

	if mov2[0]-mov1[0] == vert && math.Abs(float64(hoz)) == 1 && ! capOnly {
		return true
	}

	if mov2[0]-mov1[0] == vertCap {
		if hoz == -2 && OnBoard(mov1[0]+vert, mov1[1]+1) {
			if board[mov1[0]+vert ][mov1[1]+1] == enemy {
				return true
			}
		} else if hoz == 2 && OnBoard(mov1[0]+vert, mov1[1]-1) {
			if board[mov1[0]+vert ][mov1[1]-1] == enemy {
				return true
			}
		}
	}
	return false
}

func playerTurn(player int, capOnly bool, lastPos [2]int) {
	PrintBoard()
	reader := bufio.NewReader(os.Stdin)

	rep := "X"
	if player == 2 {
		rep = "O"
	}

	fmt.Printf("Player %d (%s) move (ex: a3 b4): \n", player, rep)
	text, _ := reader.ReadString('\n')
	if strings.HasPrefix(strings.ToLower(text), "quit") {
		playing = false
		fmt.Println("Quitting...")
		return
	} else if strings.HasPrefix(strings.ToLower(text), "skip") {
		if capOnly {
			fmt.Println("Womp womp, you have to take the capture")
			playerTurn(player, capOnly, lastPos)
			return
		}
		fmt.Println("Skipping turn")
		return
	}

	text = string([]rune(text)[0:5])

	match, _ := regexp.MatchString("[a-h][1-8] [a-h][1-8]", text)
	if ! match {
		fmt.Printf("Bad input, expected [a-h][1-8] [a-h][1-8], got %s\n", text)
		playerTurn(player, capOnly, lastPos)
		return
	}

	pos := strings.Split(text, " ")
	mov1, err1 := PosToBoard(pos[0])
	mov2, err2 := PosToBoard(pos[1])
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

	if validMove(player, mov1, mov2, capOnly) {
		king := false
		if board[mov1[0]][mov1[1]] == player+6 {
			king = true
		} else if (player == 1 && mov2[0] == 0) || (player == 2 && mov2[0] == 7) {
			king = true
			fmt.Println("King me!")
		}
		board[mov1[0]][mov1[1]] = 3
		if king {
			board[mov2[0]][mov2[1]] = player + 6
		} else {
			board[mov2[0]][mov2[1]] = player
		}

		if math.Abs(float64(mov2[0]-mov1[0])) == 2 { // Capture
			l := (mov1[0] + mov2[0]) / 2
			h := (mov1[1] + mov2[1]) / 2
			board[l][h] = 3
			if ! EnemyOnBoard(player) {
				playing = false
				PrintBoard()
				fmt.Printf("\nCongraulations Player %d, you've won!\n", player)
				return
			}
			fmt.Println("Captured!")
			_, vertCap := getPossiblePos(player)
			if validMove(player, mov2, [2]int{mov2[0] + vertCap, mov2[1] + 2}, true) ||
				validMove(player, mov2, [2]int{mov2[0] + vertCap, mov2[1] - 2}, true) {
				fmt.Println("Anoter capture possible, please go again")
				playerTurn(player, true, mov2)
			} else if king && (validMove(player, mov2, [2]int{mov2[0] - vertCap, mov2[1] + 2}, true) ||
				validMove(player, mov2, [2]int{mov2[0] - vertCap, mov2[1] - 2}, true)) {
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

	fmt.Println("Welcome to GO Checkers! ")
	fmt.Println("You play by specifiying which peice to move, and the posistion to move it too")
	fmt.Println("You can string captures, the game will let you know if you must take the next capture")
	fmt.Println("\n You can also 'skip' or 'quit'")
	fmt.Println("Have fun!")
	PopulateNewBoard()

	playing = true
	for playing {
		playerTurn(1, false, [2]int{0, 0})
		if ! playing {
			break
		}
		playerTurn(2, false, [2]int{0, 0})
	}

}
