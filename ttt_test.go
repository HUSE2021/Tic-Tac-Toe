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
	result,i= board.Play(1,3,5)
	if expected != result {
		t.Errorf("Test fail expected: %s, result: %s\n", expected, result)
	}
}


func TestTTT2(t *testing.T) {
	expected := "2win"
	board:=NewGame()
	board.InitialBoard()
	var (result=""
		i=0)
	result,i=board.Play(0,0,0)
	result,i=board.Play(0,2,i)
	result,i=board.Play(1,0,i)
	result,i=board.Play(1,2,i)
	result,i=board.Play(2,1,i)
	result,i=board.Play(2,2,i)
	//o.x
	//o.x
	//.ox  is 2 win
	result,i= board.Play(1,3,i)
	if expected != result {
		t.Errorf("Test fail expected: %s, result: %s\n", expected, result)
	}
}


func TestTTTbadinput(t *testing.T) {
	expected := "nowin"
	board:=NewGame()
	board.InitialBoard()
	var (result=""
		i=0)
	result,i=board.Play(0,0,0)
	result,i=board.Play(0,0,1)
	result,i=board.Play(0,0,i)
	result,i=board.Play(0,0,i)
	result,i=board.Play(0,0,i)
	result,i=board.Play(0,0,i)
	//o..
	//...
	//...  on one win
	result,i= board.Play(0,0,6)
	if expected != result {
		t.Errorf("Test fail expected: %s, result: %s\n", expected, result)
	}
}


func TestTTTbadinputANDreset(t *testing.T) {
	expected := "2win"
	board:=NewGame()
	board.InitialBoard()
	var (result=""
		i=0)
	result,i=board.Play(0,0,i)
	result,i=board.Play(0,0,i)  //Player2 need Re-set !!
	result,i=board.Play(0,2,i)  //Player2  Re-set !!
	result,i=board.Play(1,0,i)
	result,i=board.Play(1,2,i)
	result,i=board.Play(2,1,i)
	result,i=board.Play(2,2,i)
	//o.x
	//o.x
	//.ox  on one win
	result,i= board.Play(0,0,i)
	if expected != result {
		t.Errorf("Test fail expected: %s, result: %s\n", expected, result)
	}
}