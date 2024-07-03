package migrations

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func CreateArtifactsTable() {
    database.DB.AutoMigrate(&models.Artifact{})
}