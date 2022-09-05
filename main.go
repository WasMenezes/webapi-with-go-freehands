package main

import (
	"github.com/WasMenezes/webapi-with-go-freehands/database"
	"github.com/WasMenezes/webapi-with-go-freehands/server"
)

func main() {
	database.StartDb()
	server := server.NewServer()

	server.Run()
}
