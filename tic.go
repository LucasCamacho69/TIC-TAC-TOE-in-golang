package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var board = [][]string{
	{"0", "0", "0"},
	{"0", "0", "0"},
	{"0", "0", "0"},
}

var num_of_plays = 0
var finished bool
var player_turn bool
var player1 = "X"
var player2 = "O"


func main() {
	clear_terminal()
	fmt.Println("Starting....")
	time.Sleep(2 * time.Second)
	start_game()
}

func print_board(board [][]string) {
	clear_terminal()
	fmt.Println(board[0])
	fmt.Println(board[1])
	fmt.Println(board[2])
}

func clear_terminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
  cmd.Run()
}

func user_play(player bool) {

	if player == false {
		fmt.Print("Player 1 turn\n")
		get_user_input(player1)
	} else {
		fmt.Print("Player 2 turn\n")
		get_user_input(player2)
	}

	num_of_plays++
	player_turn = !player_turn 
}

func get_user_input(player string){
	var play int

	for {
		fmt.Print("Choose a position from 1 to 9 \n")
		_, err := fmt.Scan(&play)

		if err != nil {
			invalid_position()
			continue
		}

		if play < 1 || play > 9 {
			invalid_position()
			continue
		}

		played := make_play_in_board(play, player)

		if played == "invalid" {
			invalid_position()
			continue
		}
		break
	}
}

func invalid_position() {
	fmt.Println("Invalid position!")
	time.Sleep(1 * time.Second)
	print_board(board)
}

func make_play_in_board(position int, player string) string{
	position-- 

	i := position / 3 
	j := position % 3 

	if board[i][j] == "X" || board[i][j] == "O" {
		return "invalid"
	}
	board[i][j] = player

	return "valid"
}

func verify_win(board [][]string) string{
	//verify row winning
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && 
		board[i][1] == board[i][2] &&
		board[i][0] != "0" {
			return board[i][0]
		}
	}

	//verify col winning
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && 
		board[1][i] == board[2][i] &&
		board[0][i] != "0" {
			return board[0][i]
		}
	}

	//verify diagonal winning
	if board[1][1] != "0"{
		if board[0][0] == board[1][1] &&
		board[1][1] == board[2][2]  ||
		board[0][2] == board[1][1] &&
		board[1][1] == board[2][0]{
			return board [1][1]
		}
	}		
	return ""
}

func win_treatment(winner string) {
	clear_terminal()
	print_board(board)
	fmt.Printf("Congratulations '%s'!!!\n", winner)
	fmt.Printf("You won the game")
}

func start_game() {
	for finished == false {
		print_board(board)
		user_play(player_turn)
		
		maybe_winner := verify_win(board)

		if maybe_winner != ""{
			finished = !finished
			win_treatment(maybe_winner)
			break
		}

		if num_of_plays == 9 {
			print_board(board)
			fmt.Println("It's a tie!!!")
			break
		}
	}
}