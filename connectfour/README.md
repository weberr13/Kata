# Code Exam Connect Four Player

## installation instructions

this project is written in golang.  You can get the latest version of golang 
from https://golang.org/dl/

## build instructions

Once golang is installed, there is no "main" function but you can test with

`go test` 

from the project directory.  

## Program interface

The public function `connectfour.NewGame()` will produce a game "object"

The only public function on a Game object is `Turn(int, int)` which takes a column
to drop in (0 indexed) and a color (connectfour.Red or connectfour.Yellow)

Only variables and methods that begin with a Capital letter are public in golang.  
You may notice the tests are manipulating private variables in Game but this is only 
possible for code contained in the connectfour package. 

Tests are contained in connectfour_test.go and fall in a few categories
1) move validation
   Golang convention is that functions should return an "error" object in their last
   argument position.  This is in place of exceptions in other languages. 
2) game play tests
   These tests are not 100% exhaustive but many of them iterate over many possible morphisms of a given solution 
3) unit tests for one of the helper methods which could cause seg faults if it wasn't    right

## Design considerations

The meat of the solver is a recursive function that is given a starting point, a 
direction and a count of already matched tokens.  The function has 5 major sections

1) invalid token position.  
   Because rows are stacks implemented with slices (similar to a java.util.ArrayList) you can have an position with no token. 
2) success criteria 
   Got a connect 4
3) connect 4 search 
   Search in the given valid direction for a 4 of a kind
4) node search
   Search a reasonable set of directions that could possibly produce a 4. 
   Reasonable means there is board to work with in that direction and only searching 
   in a direction if we haven't already searched that way (and failed)
5) increment start position and start again
