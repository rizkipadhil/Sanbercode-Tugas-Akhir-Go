package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func UpdateTeamsTable() {
    database.DB.Model(&models.Team{}).Migrator().AddColumn(&models.Team{}, "VerifyBy")
    database.DB.Model(&models.Team{}).Migrator().AddColumn(&models.Team{}, "VerifyStatus")
}