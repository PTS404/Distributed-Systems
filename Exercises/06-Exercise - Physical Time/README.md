**For Normal Client**

Terminal 1:  `go run ./server/server.go -port 5454`

Terminal 2: `go run ./client/client.go -cPort 8080 -sPort 5454`

**For TimeServiceClient**

Terminal 1: `go run ./server/server.go -port 5454`

Terminal 2: `go run ./server/server.go -port 5455`

Terminal 3: `go run ./client/client.go -cPort 8080 -sPort 5454`