package tictactoe

import "fmt"

func ExampleNewBoard() {
	fmt.Println(NewBoard())

	// Output:
	// [[. . .] [. . .] [. . .]]
}

func ExamplePlace() {
	cells := [][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}
	fmt.Println(Place(cells, 1, "X"))
	fmt.Println(Place(cells, 6, "O"))
	fmt.Println(Place(cells, 1, "O"))
	fmt.Println(Place(cells, 0, "O"))
	fmt.Println(Place(cells, 10, "O"))
	fmt.Println(cells)

	// Output:
	// true
	// true
	// false
	// false
	// false
	// [[X . .] [. . .] [O . .]]
}

func ExampleNextPlayer() {
	fmt.Println(NextPlayer("X"))
	fmt.Println(NextPlayer("O"))

	// Output:
	// O
	// X
}

func ExampleWins() {
	row := [][]string{{"X", "X", "X"}, {"O", "O", "."}, {".", ".", "."}}
	col := [][]string{{"X", "O", "."}, {"X", "O", "."}, {".", "O", "X"}}
	diagonal := [][]string{{"X", "O", "."}, {"O", "X", "."}, {".", ".", "X"}}
	antiDiagonal := [][]string{{"X", "X", "O"}, {"X", "O", "."}, {"O", ".", "."}}
	empty := [][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}
	fmt.Println(Wins(row, "X"), Wins(row, "O"))
	fmt.Println(Wins(col, "X"), Wins(col, "O"))
	fmt.Println(Wins(diagonal, "X"), Wins(diagonal, "O"))
	fmt.Println(Wins(antiDiagonal, "X"), Wins(antiDiagonal, "O"))
	fmt.Println(Wins(empty, "X"))

	// Output:
	// true false
	// false true
	// true false
	// false true
	// false
}

func ExampleWinner() {
	fmt.Printf("%q\n", Winner([][]string{{"X", "X", "X"}, {"O", "O", "."}, {".", ".", "."}}))
	fmt.Printf("%q\n", Winner([][]string{{"X", "O", "."}, {"X", "O", "."}, {".", "O", "X"}}))
	fmt.Printf("%q\n", Winner([][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}))

	// Output:
	// "X"
	// "O"
	// ""
}

func ExampleIsFull() {
	fmt.Println(IsFull([][]string{{"X", "O", "X"}, {"X", "O", "O"}, {"O", "X", "X"}}))
	fmt.Println(IsFull([][]string{{"X", "O", "X"}, {"X", "O", "O"}, {"O", "X", "."}}))

	// Output:
	// true
	// false
}

func ExampleGameOver() {
	fmt.Println(GameOver([][]string{{"X", "X", "X"}, {"O", "O", "."}, {".", ".", "."}}))
	fmt.Println(GameOver([][]string{{"X", "O", "X"}, {"X", "O", "O"}, {"O", "X", "X"}}))
	fmt.Println(GameOver([][]string{{"X", "O", "X"}, {"X", "O", "O"}, {"O", "X", "."}}))

	// Output:
	// true
	// true
	// false
}
