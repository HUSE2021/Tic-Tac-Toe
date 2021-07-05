

# Tic-Tac-Toe Week4


Instruction of Program  
Tic-Tac-Toe   
The game is played on a 3-by-3 grid of the board   
Two players take turns putting their marks (o or x) at blank squares  
The first player to get 3 of his/her marks in a row, a column or a diagonal   
When all 9 squares are full, the game is over. If no player has 3 marks, the game ends in a tie.   
    
    
    
CLI Interface  
The marks of Player 1 and Player 2 are o and x   
Input a text to indicate the coordinate of square to put his/her mark 0,0, 0,1, 0,2, 1,0, 1,1, 1,2, 2,0, 2,1, 2,2    
Display the current board where the blank square is .   
When the game ends, display the result; Player 1 won, Player 2 won, Draw    



```
$ go run tictactoe.go 
Player 1: Input (x,y) 0,0
o..
...
...
Player 2: Input (x,y) 1,1
o..
.x.
...
Player 1: Input (x,y) 2,2
o..
.x.
..o
Player 2: Input (x,y) 2,0
o..
.x.
x.o
Player 1: Input (x,y) 0,2
o.o
.x.
x.o
Player 2: Input (x,y) 0,1
oxo
.x.
x.o
Player 1: Input (x,y) 1,2
oxo
.xo
x.o
Player 1 won
