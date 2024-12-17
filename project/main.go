package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PieceType int
type Color int

const (
	Normal PieceType = iota
	Queen
)

const (
	Black Color = iota
	Red
	Empty
)

type Piece struct {
	Type PieceType
	Color Color
}

type Position struct {
	X, Y int
}

type Move struct {
	Start Position
	End Position
	Captures Position
}

type Board struct {
	Grid [8][8]*Piece
	Turn Color
}

func NewBoard() *Board {
	board := &Board{
		Turn: Black,
	}
	
	for y:=0; y<3; y++ {
		for x:=0; x<8; x++ {
			if (x+y)%2 != 0 {
				board.Grid[y][x] = &Piece{
					Type: Normal, 
					Color: Black,
				}
			}
		}
	}
	
	for y:=5; y<8; y++ {
		for x:=0; x<8; x++ {
			if (x+y)%2 != 0 {
				board.Grid[y][x] = &Piece{
					Type: Normal, 
					Color: Red,
				}
			}
		}
	}
	
	return board
}

func (b *Board) VisualizeBoard() {
	fmt.Println("\n    0   1   2   3   4   5   6   7")
	fmt.Println("  +---+---+---+---+---+---+---+---+")

	for y:=0; y<8; y++ {
		
		fmt.Printf("%d |", y) // print the y number
		for x:=0; x<8; x++ {
			if b.Grid[y][x] == nil { // this space is empty
				if (x+y)%2 == 0 {
					fmt.Print(" . |")
				} else {
					fmt.Print("   |")
				}
			} else { // this space is NOT empty
				piece := b.Grid[y][x]
				
				if piece.Color == Black {
					if piece.Type == Queen {
						fmt.Print(" B |") // B for Black Queen
					} else {
						fmt.Print(" b |") // b for Black Normal
					}
				} else {
					if piece.Type == Queen {
						fmt.Print(" R |") // R for Red Queen
					} else {
						fmt.Print(" r |") // r for Red Normal
					}
				}
			}
		}
		
		fmt.Println("\n  +---+---+---+---+---+---+---+---+")
		
	}
}

func (b *Board) IsValidMove(move Move) (bool, string) {
	if !isInBounds(move.Start) || !isInBounds(move.End) {
		return false, "Move is out of bounds"
	}
	
	piece := b.Grid[move.Start.Y][move.Start.X]
	if piece == nil {
		return false, "There is no piece at the start position"
	}
	
	if b.Grid[move.End.Y][move.End.X] != nil {
		return false, "There is already a piece at the end position"
	}
	
	if piece.Color != b.Turn {
		return false, fmt.Sprintf("It's %s's turn", colorToString(b.Turn))
	}
	
	dx := move.End.X - move.Start.X
	dy := move.End.Y - move.Start.Y
	
	if abs(dx) != abs(dy) {
		return false, "Moves must be diagonal"
	}
	
	if piece.Type == Normal {
		if piece.Color == Black && dy < 0 {
			return false, "Normal black pieces can only move down"
		}
		if piece.Color == Red && dy > 0 {
			return false, "Normal red pieces can only move up"
		}
	}
	
	if abs(dx) == 1 && abs(dy) == 1 {
		return true, "" // this is a valid regular move
	} else if abs(dx) == 2 && abs(dy) == 2 {
		captureX := move.Start.X + dx/2
		captureY := move.Start.Y + dy/2
		capturedPiece := b.Grid[captureY][captureX]
		
		if capturedPiece == nil {
			return false, "There is no piece to capture"
		}
		
		if capturedPiece.Color == b.Turn {
			return false, "You can't capture your own piece"
		}
		
		return true, "" // this is a valid capture move
	} 
	
	return false, "Invalid move distance"
}

func (b *Board) MakeMove(move Move) bool {
	// Check if the move is valid
	
	valid, reason := b.IsValidMove(move)
	
	if !valid {
		fmt.Printf("Invalid move: %s\n", reason)
		return false
	}
	
	// Execute the move
	
	piece := b.Grid[move.Start.Y][move.Start.X]
	b.Grid[move.End.Y][move.End.X] = piece
	b.Grid[move.Start.Y][move.Start.X] = nil
	
	// Handle the capture
	
	if abs(move.End.X - move.Start.X) == 2 {
		captureX := (move.Start.X + move.End.X) / 2
		captureY := (move.Start.Y + move.End.Y) / 2
		capturedPiece := b.Grid[captureY][captureX]
		
		b.Grid[captureY][captureX] = nil // remove the captured piece
		fmt.Printf("Captured %s piece at %d, %d\n", colorToString(capturedPiece.Color), captureX, captureY)
	}
	
	// Handle the piece promotion
	
	if piece.Type == Normal {
		if (piece.Color == Black && move.End.Y == 7) || (piece.Color == Red && move.End.Y == 0) {
			piece.Type = Queen
			fmt.Printf("%s piece promoted to Queen at %d, %d\n", colorToString(piece.Color), move.End.X, move.End.Y)
		}
	}
	
	// Switch the turn
	
	b.Turn = opponent(b.Turn)
		
	return true
}

func GameLoop(board *Board) {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("\nWelcome to Checkers!")
	fmt.Println("Instructions:")
	fmt.Println("- Black pieces (b/B) move down the board (Increasing Y)")
	fmt.Println("- Red pieces (r/R) move up the board (Decreasing Y)")
	fmt.Println("- Uppercase letters (B/R) represent Queens")
	fmt.Println("- Enter moves in the format 'startX startY endX endY'")
	fmt.Println("- Example: '2 2 3 3' moves the piece at (2, 2) to (3, 3)")
	
	for {
		board.VisualizeBoard() // update the board for every move
		
		fmt.Printf("%s's turn\n", colorToString(board.Turn))
		fmt.Printf("Enter move (startX startY endX endY): ")
		
		if !scanner.Scan() {
			return
		}
		
		input := scanner.Text()
		coords := strings.Split(input, " ")
		
		if len(coords) != 4 {
			fmt.Println("Invalid move format. Please enter the move in the format 'startX startY endX endY'")
			continue
		}
		
		numbers := make([]int, 4)
		valid := true
		
		for i, coord := range coords {
			num, err := strconv.Atoi(coord)
			
			if err != nil || num < 0 || num > 7 {
				fmt.Println("Invalid input. Please enter numbers between 0 and 7")
				valid = false
				break
			}
			
			numbers[i] = num
		}
		
		if !valid {
			continue
		}
		
		move := Move{
			Start: Position{X: numbers[0], Y: numbers[1]},
			End: Position{X: numbers[2], Y: numbers[3]},
		}
		
		board.MakeMove(move)
	}
}

func isInBounds(position Position) bool {
	return position.X >= 0 && position.X < 8 && position.Y >= 0 && position.Y < 8 // this is a move that is in the boundaries of the board
}

func colorToString(color Color) string {
	if color == Black {
		return "Black"
	} else {
		return "Red"
	}
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func opponent(color Color) Color {
	if color == Black {
		return Red
	}
	return Black
}

func main() {
	board := NewBoard()
	GameLoop(board)
}
