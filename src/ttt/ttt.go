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

func (b *Board) put(x, y, u int) int {

}

func (b *Board) get(x, y int) string {

}

func (b *Board) print() {

}

func (b *Board) CheckWin() int {

}

func NewGame() *Board {
	b := Board{}
	b.InitialBoard()
	return &b
}

func (b *Board) Play(sx, sy, i int) (string, int) {
	var temp int
	temp = b.CheckWin()
	if temp == 0 {
		fmt.Println("Player ", i%2+1, ": Input (x,y) ")
		if b.put(sx, sy, i%2+1) == -1 {
			i--
			b.print()
		}
		b.print()
		i = i + 1
	} else if temp == 1 {
		return "1win", i + 1
	} else if temp == 2 {
		fmt.Print("Player 2 \"x\" won\n")
		return "2win", i + 1
	}
	return "nowin", i + 1
}
