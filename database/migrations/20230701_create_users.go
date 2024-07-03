package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func CreateUsersTable() {
    database.DB.AutoMigrate(&models.User{})
}