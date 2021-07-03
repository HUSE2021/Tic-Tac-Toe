package main

import (
	"fmt"
)

type Board struct {
	tokens [9]int
}

func (b *Board) InitialBoard() {
	var i int
	for i = 0; i < 9; i++ {
		b.tokens[i] = 0
	}
}

func (b *Board) put(x, y, u int)  int {
	//out of board
	if x < 0 || x > 2 || y < 0 || y > 2 {
		fmt.Print("Out of board\n")
		return  -1
	}
	//Already have pieces
	if b.tokens[3*x+y] != 0 {
		fmt.Print("Already have pieces\n")
		return -1
	}
	if u == 1 || u == 2 {
		b.tokens[3*x+y] = u
		return 0
	} else {
		fmt.Print("bad input\n")
		return -1
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[3*x+y] == 1 {
		return "o"
	} else if b.tokens[3*x+y] == 2 {
		return "x"
	} else {
		return "."
	}
}

func (b *Board) print() {
	
}

func (b *Board) CheckWin() int {

}

func NewGame() *Board{
	b:=Board{}
	b.InitialBoard()
	return &b
}


func (b *Board)Play(sx,sy,i int) (string,int) {
	var temp int
	temp = b.CheckWin()
	if temp == 0 {
		fmt.Println("Player ", i%2+1, ": Input (x,y) ")
		if b.put(sx, sy, i%2+1) == -1 {
			i--
			b.print()
			}
		b.print()
		i=i+1
	} else if temp == 1 {
		return "1win",i+1
	} else if temp == 2 {
		fmt.Print("Player 2 \"x\" won\n")
		return "2win",i+1
	}
	return "nowin",i+1
}
