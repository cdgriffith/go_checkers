package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

var playing bool
var playerOneStuck bool
var playerTwoStuck bool

func GetPossiblePos(player int) (int, int) {
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

func ValidMove(player int, mov1 [2]int, mov2 [2]int, capOnly bool, noLog bool) bool {

	if !OnBoard(mov1[0], mov1[1]) || !OnBoard(mov2[0], mov2[1]) {
		return false
	}

	enemy := 2
	if player == 2 {
		enemy = 1
	}

	if board[mov1[0]][mov1[1]] != player && board[mov1[0]][mov1[1]] != player+6 {
		if !noLog {
			fmt.Println("Starting posistion invalid", board[mov1[0]][mov1[1]])
		}
		return false
	}
	if board[mov2[0]][mov2[1]] != 3 {
		if !noLog {
			fmt.Println("Moving posistion invalid", board[mov2[0]][mov2[1]])
		}
		return false
	}

	if CaptureChecks(player, mov1, mov2, capOnly, enemy, false) {
		return true
	} else if board[mov1[0]][mov1[1]] == player+6 && CaptureChecks(player, mov1, mov2, capOnly, enemy, true) {
		return true
	}

	return false
}

func CaptureChecks(player int, mov1 [2]int, mov2 [2]int, capOnly bool, enemy int, kingFlip bool) bool {
	hoz := mov1[1] - mov2[1]
	vert, vertCap := GetPossiblePos(player)
	if kingFlip {
		vert = -vert
		vertCap = -vertCap
	}

	if mov2[0]-mov1[0] == vert && math.Abs(float64(hoz)) == 1 && !capOnly {
		return true
	}

	if mov2[0]-mov1[0] == vertCap {
		if hoz == -2 && OnBoard(mov1[0]+vert, mov1[1]+1) {
			if board[mov1[0]+vert][mov1[1]+1] == enemy || board[mov1[0]+vert][mov1[1]+1] == enemy+6 {
				return true
			}
		} else if hoz == 2 && OnBoard(mov1[0]+vert, mov1[1]-1) {
			if board[mov1[0]+vert][mov1[1]-1] == enemy || board[mov1[0]+vert][mov1[1]-1] == enemy+6 {
				return true
			}
		}
	}
	return false
}

func ValidMoves(player int) [][2][2]int {
	var valids [][2][2]int
	vert, vertCap := GetPossiblePos(player)

	for i, row := range board {
		for j, element := range row {
			if element != player && element != player+6 {
				continue
			}
			checkMoves := [4][2]int{
				{i + vert, j + 1},
				{i + vert, j - 1},
				{i + vertCap, j + 2},
				{i + vertCap, j - 2},
			}
			currentPos := [2]int{i, j}
			for _, move := range checkMoves {
				if ValidMove(player, currentPos, move, false, true) {
					valids = append(valids, [2][2]int{currentPos, move})
				}
			}

			if element == player+6 {
				kingMoves := [4][2]int{
					{i - vert, j + 1},
					{i - vert, j - 1},
					{i - vertCap, j + 2},
					{i - vertCap, j - 2},
				}
				for _, move := range kingMoves {
					if ValidMove(player, currentPos, move, false, true) {
						valids = append(valids, [2][2]int{currentPos, move})
					}
				}
			}
		}
	}
	return valids
}

// For later, when AI or UI is added, fast way to find valid moves for a single piece
func ValidMovesForPiece(player int, piece [2]int)  [][2]int {
	var valids [][2]int

	for _, move := range ValidMoves(player) {
		if move[0] == piece {
			valids = append(valids, move[1])
		}
	}
	return valids
}


func PlayerTurn(player int, capOnly bool, lastPos [2]int) {
	PrintBoard()
	if len(ValidMoves(player)) == 0 {
		fmt.Println("No moves for you, bad luck!")
		if player == 1 {
			playerOneStuck = true
		} else {
			playerTwoStuck = true
		}
		if playerOneStuck && playerTwoStuck {
			fmt.Println("It's a draw!")
			playing = false
			return
		}

	} else {
		if player == 1 {
			playerOneStuck = false
		} else {
			playerTwoStuck = false
		}
	}

	reader := bufio.NewReader(os.Stdin)

	rep := "X"
	if player == 2 {
		rep = "O"
	}

	fmt.Printf("Player %d (%s) move (ex: a3 b4): \n", player, rep)
	text, _ := reader.ReadString('\n')
	fmt.Println()
	if strings.HasPrefix(strings.ToLower(text), "quit") {
		playing = false
		enemy := 2
		if player == 2 {
			enemy = 1
		}
		fmt.Printf("Player %d has quit, Player %d wins!\n", player, enemy)
		return
	}

	text = string([]rune(strings.ToLower(text))[0:5])

	match, _ := regexp.MatchString(`^[a-h][1-8] [a-h][1-8].*`, text)
	if !match {
		fmt.Printf("Bad input, expected [a-h][1-8] [a-h][1-8], got %s\n", text)
		PlayerTurn(player, capOnly, lastPos)
		return
	}

	pos := strings.Split(text, " ")
	mov1, err1 := PosToBoard(pos[0])
	mov2, err2 := PosToBoard(pos[1])
	if capOnly && mov1 != lastPos {
		fmt.Println("You have to move the same piece you just used last, numskull")
		PlayerTurn(player, capOnly, lastPos)
		return
	}

	if err1 || err2 {
		fmt.Println("Invalid move, buddy")
		PlayerTurn(player, capOnly, lastPos)
		return
	}

	if ValidMove(player, mov1, mov2, capOnly, false) {
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
			if !EnemyOnBoard(player) {
				playing = false
				PrintBoard()
				fmt.Printf("\nCongraulations Player %d, you've won!\n", player)
				return
			}
			fmt.Println("Captured!")
			_, vertCap := GetPossiblePos(player)
			if ValidMove(player, mov2, [2]int{mov2[0] + vertCap, mov2[1] + 2}, true, true) ||
				ValidMove(player, mov2, [2]int{mov2[0] + vertCap, mov2[1] - 2}, true, true) {
				fmt.Println("Anoter capture possible, please go again")
				PlayerTurn(player, true, mov2)
			} else if king && (ValidMove(player, mov2, [2]int{mov2[0] - vertCap, mov2[1] + 2}, true, true) ||
				ValidMove(player, mov2, [2]int{mov2[0] - vertCap, mov2[1] - 2}, true, true)) {
				fmt.Println("Anoter capture possible, please go again")
				PlayerTurn(player, true, mov2)
			}
		}

	} else {
		fmt.Println("INVALID MOVE!")
		PlayerTurn(player, capOnly, lastPos)
		return
	}

}

func main() {
	playerOneStuck = false
	playerTwoStuck = false
	fmt.Println("Welcome to GO Checkers!\n")
	fmt.Println("You play by specifiying which peice to move, and the posistion to move it too")
	fmt.Println("You can string captures, the game will let you know if you must take the next capture")
	fmt.Println("You can also 'quit' at anytime")
	fmt.Println("\nHave fun!\n")
	PopulateNewBoard()
	playing = true
	for playing {
		PlayerTurn(1, false, [2]int{0, 0})
		if !playing {
			break
		}
		PlayerTurn(2, false, [2]int{0, 0})
	}

}
