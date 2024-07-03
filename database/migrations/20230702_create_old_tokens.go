package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func CreateOldTokensTable() {
    database.DB.AutoMigrate(&models.OldToken{})
}