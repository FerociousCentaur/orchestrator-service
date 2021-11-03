# orchestrator-service
<strong style="color:red">Notice:</strong>
- Folder 'client','proto' and 'server' at the highest level is for part 2
- Folder 'datamock' is for part 3
- Folder 'logic' is for part 4
- Each part is independent of each other

<b>'Note'</b> Its better to clone the repo in a folder called RPC
Operating System used = 'Windows'
GOPATH = %USERPROFILE%\go
Repository Path = "D:\Books\Go projects\RPC"

## Part 2
<b>Commands to run the server</b>
- cd client
- go run main.go
- cd server
- go run main.go

<b>How to Test</b>
- Go to `localhost:8080/user/<somename>`
- The desired output for part 2 is recieved


## Part 3
<b>Commands to run the server</b>
- cd datamock/client
- go run main.go
- cd datamock/server
- go run main.go

<b>How to Test</b>
- Go to `localhost:8080/user/<somename>`
- The desired output for part 3 is recieved


## Part 4
<b>Commands to run the server</b>
- cd logic/client  (runs on port 8080 and dials on 9000)
- go run main.go
- cd logic/middleware (listens on port 9000 and dials on 9001)
- go run main.go
- cd logic/middleware2  (listens on port 9001 and dials on 10000)
- go run main.go
- cd datamock/server (listens on port 10000)
- go run main.go

<b>How to Test</b>
- Go to `localhost:8080/user/<somename>`
- The desired output for part 4 is recieved





