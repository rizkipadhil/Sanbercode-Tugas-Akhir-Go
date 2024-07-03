package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func MigrateAll() {
    database.DB.AutoMigrate(&models.Element{})
    database.DB.AutoMigrate(&models.Weapon{})
    database.DB.AutoMigrate(&models.Character{})
    database.DB.AutoMigrate(&models.Artifact{})
    database.DB.AutoMigrate(&models.Team{})
    database.DB.AutoMigrate(&models.TeamCharacter{})
    database.DB.AutoMigrate(&models.User{})
    database.DB.AutoMigrate(&models.OldToken{})
}