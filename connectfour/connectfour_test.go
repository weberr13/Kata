package connectfour

import ("testing")
	
func doesWin(t *testing.T, f func() (bool, error)) {
	won, err := f()
	if (!won) {
		t.Fatalf("expected win")
	}
	if (err != nil) {
		t.Fatal("unexpected error")
	}
}
func doesnotWin(t *testing.T, f func() (bool, error)) {
	won, err := f()
	if (won) {
		t.Fatal("expected not to win")
	}
	if (err != nil) {
		t.Fatal("unexpected error")
	}
}
func illegalMove(t *testing.T, f func() (bool, error)) {
	won, err := f()
	if (won) {
		t.Fatal("expected error not win")
	}
	if (err == nil) {
		t.Fatal("expected error")
	}
}
func TestIllegalMovesTooManyTokens(t *testing.T) {
	g := NewGame()
	g.cols[0] = Column{row: []int{Red, Red, Red, Yellow, Yellow, Yellow}}
	g.cols[1] = Column{}
	g.cols[2] = Column{}
	g.cols[3] = Column{}
	g.cols[4] = Column{}
	g.cols[5] = Column{}
	g.cols[6] = Column{}

	illegalMove(t, func() (bool, error) {return g.Turn(0, Red )})
}

func TestIllegalMovesBadColumn(t *testing.T) {
	g := NewGame()

	illegalMove(t, func() (bool, error) {return g.Turn(7, Red )})
	illegalMove(t, func() (bool, error) {return g.Turn(-1, Red )})
}

func TestIllegalBackToBack(t *testing.T) {
	g := NewGame()

	doesnotWin(t, func() (bool, error) {return g.Turn(0, Red )})
	illegalMove(t, func() (bool, error) {return g.Turn(0, Red )})
}

func TestDownWin(t *testing.T) {
	g := NewGame()

	l := Yellow
	for w := range []int{Red,Yellow} {
		g.last = na
		if (w == Yellow) {
			l = Red
		}
		for row := 0 ; row < 3 ; row++ {
			for col := 0 ; col < 7 ; col++ {
				g.last = na
				g.cols[0] = Column{}
				g.cols[1] = Column{}
				g.cols[2] = Column{}
				g.cols[3] = Column{}
				g.cols[4] = Column{}
				g.cols[5] = Column{}
				g.cols[6] = Column{}
				
				for i:=0 ; i < row ; i++ {
					g.cols[col].row = append(g.cols[col].row, l)
				}
				g.cols[col].row = append(g.cols[col].row, w)
				g.cols[col].row = append(g.cols[col].row, w)
				g.cols[col].row = append(g.cols[col].row, w)

				doesWin(t, func() (bool, error) {return g.Turn(col, w )})
			}
		}
	}
}
func TestAccrossWin(t *testing.T) {
	g := NewGame()

	l := Yellow

	for w := range []int{Red,Yellow} {
		if (w == Yellow) {
			l = Red
		}
		offset := 0
		g.last = na
		g.cols[offset] =            Column{row: []int{w, w, w, l, l, l}}
		g.cols[nextCol(offset+1)] = Column{row: []int{w, l, l, l, w, l}}
		g.cols[nextCol(offset+2)] = Column{row: []int{w, w, w, l, w, l}}
		g.cols[nextCol(offset+3)] = Column{row: []int{l, w, l, w, l, w}}
		g.cols[nextCol(offset+4)] = Column{row: []int{w, l, w, w, l, w}}
		g.cols[nextCol(offset+5)] = Column{row: []int{w, l, w, l, w, w}}
		g.cols[nextCol(offset+6)] = Column{row: []int{w, w, w, l, w}}

		doesWin(t, func() (bool, error) {return g.Turn(offset+6, w)})

		g.last = na
		g.cols[0] = Column{row: []int{w}}
		g.cols[1] = Column{row: []int{w}}
		g.cols[2] = Column{row: []int{w}}
		g.cols[3] = Column{}
		g.cols[4] = Column{}
		g.cols[5] = Column{}
		g.cols[6] = Column{}
		doesWin(t, func() (bool, error) {return g.Turn(3, w)})

		g.cols[0] = Column{}
		g.cols[1] = Column{}
		g.cols[2] = Column{}
		g.cols[3] = Column{}
		g.cols[4] = Column{row: []int{w}}
		g.cols[5] = Column{row: []int{w}}
		g.cols[6] = Column{row: []int{w}}
		g.last = na
		doesWin(t, func() (bool, error) {return g.Turn(3, w)})

		g.cols[0] = Column{}
		g.cols[1] = Column{row: []int{w}}
		g.cols[2] = Column{row: []int{w}}
		g.cols[3] = Column{row: []int{w}}
		g.cols[4] = Column{}
		g.cols[5] = Column{}
		g.cols[6] = Column{}
		g.last = na
		doesWin(t, func() (bool, error) {return g.Turn(4, w)})

		g.cols[4] = Column{}
		g.last = na
		doesWin(t, func() (bool, error) {return g.Turn(0, w)})

		g.cols[0] = Column{}
		g.cols[1] = Column{row: []int{w}}
		g.cols[2] = Column{row: []int{w}}
		g.cols[3] = Column{}
		g.cols[4] = Column{row: []int{w}}
		g.cols[5] = Column{}
		g.cols[6] = Column{}
		g.last = na
		doesWin(t, func() (bool, error) {return g.Turn(3, w)})

		g.cols[0] = Column{row: []int{w}}
		g.cols[1] = Column{}
		g.cols[2] = Column{row: []int{w}}
		g.cols[3] = Column{}
		g.cols[4] = Column{row: []int{w}}
		g.cols[5] = Column{}
		g.cols[6] = Column{row: []int{w}}
		g.last = na
		doesnotWin(t, func() (bool, error) {return g.Turn(1, w)})

		g.cols[0] = Column{row: []int{w}}
		g.cols[1] = Column{}
		g.cols[2] = Column{row: []int{w}}
		g.cols[3] = Column{}
		g.cols[4] = Column{row: []int{w}}
		g.cols[5] = Column{}
		g.cols[6] = Column{row: []int{w}}
		g.last = na		
		doesnotWin(t, func() (bool, error) {return g.Turn(3, w)})

		g.cols[0] = Column{row: []int{w}}
		g.cols[1] = Column{}
		g.cols[2] = Column{row: []int{w}}
		g.cols[3] = Column{}
		g.cols[4] = Column{row: []int{w}}
		g.cols[5] = Column{}
		g.cols[6] = Column{row: []int{w}}
		g.last = na		
		doesnotWin(t, func() (bool, error) {return g.Turn(5, w)})

	}
	
}

func nextCol(curr int) int {
	return (curr % 7);
}

func TestCatsGame(t *testing.T) {
	g := NewGame()

	l := Yellow
	offset := 0
	for w := range []int{Red,Yellow} {
		if (w == Yellow) {
			l = Red
		}
		g.cols[offset] =            Column{row: []int{w, w, w, l, l, l}}
		g.cols[nextCol(offset+1)] = Column{row: []int{w, l, l, l, w, l}}
		g.cols[nextCol(offset+2)] = Column{row: []int{w, w, w, l, w, w}}
		g.cols[nextCol(offset+3)] = Column{row: []int{l, w, l, w, l, l}}
		g.cols[nextCol(offset+4)] = Column{row: []int{w, l, w, w, l, w}}
		g.cols[nextCol(offset+5)] = Column{row: []int{w, l, w, l, w, w}}
		g.cols[nextCol(offset+6)] = Column{row: []int{w, w, w, l, w}}

		doesnotWin(t, func() (bool, error) {return g.Turn(nextCol(offset+6), w)})
	}

}
func TestDiagnonalWin(t *testing.T) {
	g := NewGame()

	l := Yellow
	for offset := 0 ; offset < 3 ; offset++ {

		for w := range []int{Red,Yellow} {
			if (w == Yellow) {
				l = Red
			}
	
			g.last = na
			g.cols[offset] = Column{row: []int{w}}
			g.cols[nextCol(offset+1)] = Column{row: []int{l, w}}
			g.cols[nextCol(offset+2)] = Column{row: []int{l, l, w}}
			g.cols[nextCol(offset+3)] = Column{row: []int{l, l, l}}
			g.cols[nextCol(offset+4)] = Column{}
			g.cols[nextCol(offset+5)] = Column{}
			g.cols[nextCol(offset+6)] = Column{}
			doesWin(t, func() (bool, error) {return g.Turn(nextCol(offset+3), w)})
	
			g.last = na
			g.cols[offset] = Column{row: []int{l, l, l}}
			g.cols[nextCol(offset+1)] = Column{row: []int{l, l, w}}
			g.cols[nextCol(offset+2)] = Column{row: []int{l, w}}
			g.cols[nextCol(offset+3)] = Column{row: []int{w}}
			g.cols[nextCol(offset+4)] = Column{}
			g.cols[nextCol(offset+5)] = Column{}
			g.cols[nextCol(offset+6)] = Column{}
			doesWin(t, func() (bool, error) {return g.Turn(nextCol(offset), w)})
	
			g.last = na
			g.cols[offset] = Column{}
			g.cols[nextCol(offset+1)] = Column{row: []int{l, w}}
			g.cols[nextCol(offset+2)] = Column{row: []int{l, l, w}}
			g.cols[nextCol(offset+3)] = Column{row: []int{l, l, l, w}}
			g.cols[nextCol(offset+4)] = Column{}
			g.cols[nextCol(offset+5)] = Column{}
			g.cols[nextCol(offset+6)] = Column{}
			doesWin(t, func() (bool, error) {return g.Turn(offset, w)})
	
			g.last = na
			g.cols[offset] = Column{row: []int{l, l, l, w}}
			g.cols[nextCol(offset+1)] = Column{row: []int{l, l, w}}
			g.cols[nextCol(offset+2)] = Column{row: []int{l, w}}
			g.cols[nextCol(offset+3)] = Column{}
			g.cols[nextCol(offset+4)] = Column{}
			g.cols[nextCol(offset+5)] = Column{}
			g.cols[nextCol(offset+6)] = Column{}
			doesWin(t, func() (bool, error) {return g.Turn(nextCol(offset+3), w)})
	
			g.last = na
			g.cols[offset] = Column{row: []int{w}}
			g.cols[nextCol(offset+1)] = Column{row: []int{l}}
			g.cols[nextCol(offset+2)] = Column{row: []int{l, l, w}}
			g.cols[nextCol(offset+3)] = Column{row: []int{l, l, l, w}}
			g.cols[nextCol(offset+4)] = Column{}
			g.cols[nextCol(offset+5)] = Column{}
			g.cols[nextCol(offset+6)] = Column{}
			doesWin(t, func() (bool, error) {return g.Turn(nextCol(offset+1), w)})
	
			g.last = na
			g.cols[offset] = Column{row: []int{l, l, l, w}}
			g.cols[nextCol(offset+1)] = Column{row: []int{l, l}}
			g.cols[nextCol(offset+2)] = Column{row: []int{l, w}}
			g.cols[nextCol(offset+3)] = Column{row: []int{w}}
			g.cols[nextCol(offset+4)] = Column{}
			g.cols[nextCol(offset+5)] = Column{}
			g.cols[nextCol(offset+6)] = Column{}
			doesWin(t, func() (bool, error) {return g.Turn(nextCol(offset+1), w)})
	
			g.last = na
			g.cols[offset] = Column{row: []int{w}}
			g.cols[nextCol(offset+1)] = Column{row: []int{l, w}}
			g.cols[nextCol(offset+2)] = Column{row: []int{l, l}}
			g.cols[nextCol(offset+3)] = Column{row: []int{l, l, l, w}}
			g.cols[nextCol(offset+4)] = Column{}
			g.cols[nextCol(offset+5)] = Column{}
			g.cols[nextCol(offset+6)] = Column{}
			doesWin(t, func() (bool, error) {return g.Turn(nextCol(offset+2), w)})
	
			g.last = na
			g.cols[offset] = Column{row: []int{l, l, l, w}}
			g.cols[nextCol(offset+1)] = Column{row: []int{l, l, w}}
			g.cols[nextCol(offset+2)] = Column{row: []int{l}}
			g.cols[nextCol(offset+3)] = Column{row: []int{w}}
			g.cols[nextCol(offset+4)] = Column{}
			g.cols[nextCol(offset+5)] = Column{}
			g.cols[nextCol(offset+6)] = Column{}
			doesWin(t, func() (bool, error) {return g.Turn(nextCol(offset+2), w)})
		}
	}

}


// Unit tests for helper method getToken
func TestGetTokenEmpty(t *testing.T) {
	g := NewGame()

	g.cols[0] = Column{}
	g.cols[1] = Column{}
	g.cols[2] = Column{}
	g.cols[3] = Column{}
	g.cols[4] = Column{}
	g.cols[5] = Column{}
	g.cols[6] = Column{}

	for col := 0 ; col < 7 ; col++ {
		for row := 0 ; row < 6 ; row++ {
			if (g.getToken(col,row) != na) {
				t.Fatal("invalid token")
			}
		}
	}
}
func TestGetTokenFull(t *testing.T) {
	g := NewGame()

	g.cols[0] = Column{}
	g.cols[1] = Column{}
	g.cols[2] = Column{}
	g.cols[3] = Column{}
	g.cols[4] = Column{}
	g.cols[5] = Column{}
	g.cols[6] = Column{}
	for col := 0 ; col < 7 ; col ++ {
		for row := 0 ; row < 6 ; row++ {
			g.cols[col].row = append(g.cols[col].row, Red)
		}
	}
	for col := 0 ; col < 7 ; col++ {
		for row := 0 ; row < 6 ; row++ {
			if (g.getToken(col,row) != Red) {
				t.Logf("%v", g.getToken(col, row))
				t.Fatal("invalid read")
			}
			if (g.getToken(-1, row) != na) {
				t.Fatal("fell off the edge")
			}
			if (g.getToken(8, row) != na) {
				t.Fatal("fell off the edge")
			}
		}
		if (g.getToken(col, -1) != na) {
			t.Fatal("fell off the bottom")
		}
		if (g.getToken(col, 7) != na) {
			t.Fatal("fell off the top")
		}
	}
	
}


