package main

import (
	"embed"
	"tx_service/internal/app"

	_ "github.com/lib/pq"
)

const serviceName = "tx_service"

//go:embed dbschema/migrations
var dbMigrationFS embed.FS


func main() {
	a := app.New(serviceName)
	a.Run(dbMigrationFS)
}
