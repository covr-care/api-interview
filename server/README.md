# Cross Compilation Commands

`CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -a -o ./bin/server-osx cmd/main.go`
`CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -a -o ./bin/server-windows.exe cmd/main.go`
`CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -a -o ./bin/server-osx-intel cmd/main.go`

# Database creation

1. `sqlite3 interview.db`
2. `sqlite3 interview.db < migrations/init.sql`
3. `sqlite3 interview.db < migrations/seed_time_punches.sql`
