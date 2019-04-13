package connectfour

import (
	"errors"
)
///To summarize, we are looking for:
// 1.  Data model representing the current state of the game
// 2.  Implementation of the function that will be called after each game turn in order to update the model and check if we have a winner
// a.    Accepts the column into which the piece has been placed
// b.    Accepts the type of piece placed
// c.    Accepts the current board object (if necessary)
// d.    Returns true if last move is the winner, false otherwise
// 3.  Unit tests that validate the logic of the function

// Things to focus on:
// 1.       Elegant and fast implementation (hint: can use recursion)
// 2.       Clean and easy to read code (use comments and meaningful variable names)
// 3.       Good tests
const (
	//Red player token
	Red = iota
	//Yellow player token
	Yellow 
	//No token
	na
)
const (
	up = iota
	down 
	left 
	right 
	upLeft
	downLeft 
	upRight 
	downRight 
	noDirection 
)
const (
	//MaxRows is the depth of a column
	MaxRows = 6
	//MaxColumns is the number of columns
	MaxColumns = 7
)
//Game is a running game
type Game struct {
	cols []Column
	last int // add sequence checking never repeat red or yellow twice
}
//NewGame constructor
func NewGame() (g *Game) {
	return &Game{cols: make([]Column, MaxColumns), last: na}
}

//Column of tokens
type Column struct {
	row []int
}

func (g Game) getToken(col int, row int) int {
	if ((col < 0) || (col > MaxColumns - 1)) {
		return na // off the board
	}
	if ((len(g.cols[col].row) - 1 < row) || row < 0) {
		return na // No token here
	}
	return g.cols[col].row[row]
}

func getNextPos(direction int) (col int, row int) {
	switch direction {
		case up: 
			return 0,1
		case down:
			return 0, -1
		case left:
			return -1, 0
		case right:
			return 1, 0
		case upLeft: 
			return -1, 1
		case downLeft:
			return -1, -1
		case upRight:
			return 1,1
		case downRight:
			return 1, -1	
	}
	return 0,0
}

// look back and see if we've likely been here before
func getOpositePos(direction int) (col int, row int) {
	switch direction {
		case down: 
			return 0,1
		case up:
			return 0, -1
		case right:
			return -1, 0
		case left:
			return 1, 0
		case downRight: 
			return -1, 1
		case upRight:
			return -1, -1
		case downLeft:
			return 1,1
		case upLeft:
			return 1, -1	
	}
	return 0,0
}

// The search starts at the bottom right so right will find left, and up will find down
func getReasonableDirections(col int, row int) []int {
	directions := []int{}
	
	if (row < MaxRows - 3) {
		directions = append(directions, up)
	}
	if (col < MaxColumns - 3) {
		if (row < MaxRows - 3) {
			directions = append(directions, upRight)
		}
		directions = append(directions, right)
	}
	if (col > MaxColumns - 5) {
		if (row < MaxRows - 3) {
			directions = append(directions, upLeft)
		}
	}

	return directions
}
//recursive search from a given position
func (g Game) directionSearch(col int, row int, direction int, count int) (win bool, newcount int) {
	current := g.getToken(col, row)

	// direction search has an invalid postion.  This occurs when columns don't have a
	// token at this row position
	if (current == na) {
		if (col < MaxColumns - 1) {
			return g.directionSearch(col+1, row, noDirection, 0)
		} else if (row < MaxRows - 1) {
			return g.directionSearch(0, row+1, noDirection, 0)
		}
		return false, 0
	}
	// This is the only win criteria
	if (count == 4) {			
		//fmt.Printf("%v won at %v %v in direction %v", current, col, row, direction)
		return true, 0
	}
	newcount = count + 1 // we have a valid token

	//if direction is valid then we are in the process of searching for a connect 4
	if (direction != noDirection) {
		coldif,rowdif := getNextPos(direction)
		if ((g.getToken(col+coldif, row+rowdif) == current)) {
			return g.directionSearch(col+coldif, row+rowdif, direction, newcount)
		} 
		// Either the next location is empty or the wrong color
		return false, 0
	} 
	
	//the above noDirection code will always return so we know that we haven't started a search
	for _, possibleDirection := range getReasonableDirections(col, row) {
		colOld, rowOld := getOpositePos(possibleDirection)
		if (current == g.getToken(col+colOld, row+rowOld)) {
			continue // we've been here before in this direction, this is more efficent than a "visited" table
		}
		win, _ = g.directionSearch(col, row, possibleDirection, newcount)
		if (win) {
			return true, 0
		} 
	}

	// If we can start the search anew do so
	if (col < MaxColumns - 1) {
		return g.directionSearch(col+1, row, noDirection, 0)
	} else if (row < MaxRows - 1) {
		return g.directionSearch(0, row+1, noDirection, 0)
	}
	// we've run out of options
	return false, 0
}

func (g Game) isWon() bool {
	win, _ := g.directionSearch(0, 0, noDirection, 0)

	return win;
}

//Turn Take a turn, specify the column and the color
func (g *Game) Turn(colNum int, color int) (winner bool, err error) {
	if ((colNum >= MaxColumns) || (colNum < 0)) {
		return false, errors.New("illegal move! No place to drop")
	}
	if (len(g.cols[colNum].row) >= MaxRows) {
		return false, errors.New("illegal move! Column is full")
	}
	if (color == g.last) {
		return false, errors.New("illeagl move! Not your turn, you cheater")
	}
	g.last = color
	g.cols[colNum].row = append(g.cols[colNum].row, color)
	win, _ := g.directionSearch(0, 0, noDirection, 0)

	return win, nil
}