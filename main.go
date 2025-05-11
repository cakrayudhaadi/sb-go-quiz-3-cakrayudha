package main

import (
	"quiz3/configs"
	"quiz3/databases/connection"
	"quiz3/databases/migration"
	"quiz3/routers"

	_ "github.com/lib/pq"
)

func main() {
	configs.Initiator()
	connection.Initiator()
	defer connection.SqlDBConnections.Close()
	migration.Initiator(connection.SqlDBConnections)
	routers.StartServer()
}
