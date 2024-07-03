package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func CreateTeamsTable() {
    database.DB.AutoMigrate(&models.Team{})
}