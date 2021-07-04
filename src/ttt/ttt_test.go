package main

import (
	"testing"
)

func TestTTT(t *testing.T) {
	expected := "1win"
	board:=NewGame()
	board.InitialBoard()
	var (result=""
		i=0)
	result,i=board.Play(0,0,0)
	result,i=board.Play(0,2,i)
	result,i=board.Play(1,0,i)
	result,i=board.Play(1,2,i)
	result,i=board.Play(2,0,i)
	//o.x
	//o.x
	//o..  is 1 win
	result,i= board.Play(1,3,i)
	if expected != result {
		t.Errorf("Test fail expected: %s, result: %s\n", expected, result)
	}
}


func TestTTT2(t *testing.T) {
	
}


func TestTTTbadinput(t *testing.T) {
	
}


func TestTTTbadinputANDreset(t *testing.T) {
	
}
