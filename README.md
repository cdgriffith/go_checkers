# Checkers

A simple console based checkers game written in golang. 


```
Welcome to GO Checkers!

You play by specifying which piece to move, and the position to move it too
You can string captures, the game will let you know if you must take the next capture
You can also 'quit' at anytime

Have fun!

8 | = | O | = | O | = | O | = | O |
7 | O | = | O | = | O | = | O | = |
6 | = | O | = | O | = | O | = | O |
5 |   | = |   | = |   | = |   | = |
4 | = |   | = |   | = |   | = |   |
3 | X | = | X | = | X | = | X | = |
2 | = | X | = | X | = | X | = | X |
1 | X | = | X | = | X | = | X | = |
    A   B   C   D   E   F   G   H 
    
Player 1 (X) move (ex: a3 b4): 


```

## Build

Check out the releases page for already compiled versions https://github.com/cdgriffith/go_checkers/releases

If you don't trust a random person on the internet's binaries, feel free to compile them yourself:

GOOS=windows GOARCH=amd64 go build -o go_checkers_windows_x64.exe
GOOS=linux GOARCH=amd64 go build -o go_checkers_linux_x64

If you have a really old 32-bit system

GOOS=windows GOARCH=386 go build -o go_checkers_windows_32.exe
GOOS=linux GOARCH=386 go build -o go_checkers_linux_32

## Feedback

This is my first program with golang so there will be bugs and major room for improvement.
Feel free to open issues to point out any errors or just coding best practices and suggestions!
