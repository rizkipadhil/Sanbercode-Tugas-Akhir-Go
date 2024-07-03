package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func CreateTeamCharactersTable() {
    database.DB.AutoMigrate(&models.TeamCharacter{})
}