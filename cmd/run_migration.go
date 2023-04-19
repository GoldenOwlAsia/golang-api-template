package main

import (
	"api/api"
	"api/migrations"
)

func main() {
	api.Api = api.NewServer()
	api.Api.InitEnv().InitDb()
	migrations.Migrate()
	migrations.Seed()
}
