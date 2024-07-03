package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func CreateElementsTable() {
    database.DB.AutoMigrate(&models.Element{})
}