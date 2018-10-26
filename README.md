# TicTacToe - console version playable
ttt built using golang

# Set up
- clone this repo
  - run `$ git clone git@github.com:rrugamba/go-ttt.git`

- to check your go version
  - run `$ go version`
  - my version is: `go1.11.1 darwin/amd64`

# Running tests
- **core** 
   - cd into a directory for any of the packages ie board, player etc
   - run `go test` inside the specific package directory
   - forexample (within cloned repo):
     - to run board tests 
        - `$ cd core/src/board`
        - `$ go test` - to run all tests for the board package
     - to run moveStrategy tests
        - `$ cd core/src/moveStrategy`
        - `$ go test` - to run all tests for the moveStrategy package

- **console**
  - _within cloned repo
    - `$ cd console/test/setup`
    - `$ go test` - to run all tests for the setup package
  

# Play Game
  - **Console version**
    -  _within the cloned repo_
        - `$ cd console/src/main`
        - `$ go run main.go` - should start game options to play TTT
  
