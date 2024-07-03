package main

import (
    "github.com/joho/godotenv"
    "tugas-akhir/database"
    "tugas-akhir/database/migrations"
    "tugas-akhir/database/seeders"
    "tugas-akhir/router"
)

func main() {
    godotenv.Load()
    database.Connect()
    migrations.MigrateAll()
    seeders.SeedAll()

    r := router.SetupRouter()
    r.Run()
}