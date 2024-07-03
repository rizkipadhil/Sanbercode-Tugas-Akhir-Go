package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func CreateWeaponsTable() {
    database.DB.AutoMigrate(&models.Weapon{})
}