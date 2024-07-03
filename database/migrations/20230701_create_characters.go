package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func CreateCharactersTable() {
    database.DB.AutoMigrate(&models.Character{})
}