# tictactoe - console version playable
- ttt built using golang
- only on a 3 x 3 board

# set up
- if installing go for the first time, 
  - follow these [mac installation instructions](https://jjude.com/golang-on-macosx/)

- to check your go version
  - run `$ go version`
  - my version is: `go1.11.1 darwin/amd64`
  
- clone this repo into your go directory you set up from the installation instructions
  - run `$ git clone git@github.com:rrugamba/go-ttt.git`

# running tests
- **core** 
   - to run all tests
     - inside the `core/test` directory
     - run `$ go test ./...`
   
   - to run unit tests for packages
      - run `go test` inside the specific package directory
      - forexample (within cloned repo):
        - to run board tests 
          - `$ cd core/test/board`
          - `$ go test` - to run all tests for the board package
        - to run moveStrategy tests
          - `$ cd core/test/moveStrategy`
          - `$ go test` - to run all tests for the moveStrategy package

- **console**
  - to run all tests
    - inside the `console/test` directory
    - run `$ go test ./...`
    
  - to run unit tests for packages
    - `$ cd console/test/setup`
    - `$ go test` - to run all tests for the setup package
  

# play game
  - **console version**
    -  _within the cloned repo_
        - `$ cd console/src/main`
        - `$ go run main.go` - should start game options to play TTT
  
