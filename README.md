# tictactoe - console version playable
- ttt built using golang
- only on a 3 x 3 board

# set up
- to check your go version
  - run `$ go version`
  - my version is: `go1.11.1 darwin/amd64`
  
- clone this repo 
  - run `$ git clone git@github.com:rrugamba/go-ttt.git`

- add these lines in the `.bash_profile`
  ```
  export GOPATH=$HOME/path_to_your_cloned_repo
  export GOROOT=/usr/local/opt/go/libexec
  export PATH=$PATH:$GOPATH/bin
  export PATH=$PATH:$GOROOT/bin
  ```
# get project dependencies
- navigate to your root directory
  - run `$ cd /`
  - then `$ cd usr/local/opt/go/libexec/src/`
  - then `$ go get github.com/franela/goblin`
  
# running tests
- **core** 
   - to run all tests
     - inside the `core/test` directory
     - run `$ go test ./...`
   
   - to run unit tests for packages
      - run `go test` inside the specific package directory
      - forexample to run board tests 
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
  
