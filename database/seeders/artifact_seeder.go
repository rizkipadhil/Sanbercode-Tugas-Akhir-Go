package seeders

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func SeedArtifacts() {
    var count int64
    database.DB.Model(&models.Artifact{}).Count(&count)
    if count > 0 {
        log.Println("Seeder sudah dijalankan sebelumnya.")
        return
    }

    file, err := ioutil.ReadFile("data/artifacts.json")
    if err != nil {
        log.Fatalf("Failed to read artifacts.json file: %v", err)
    }

    var artifacts []models.Artifact
    if err := json.Unmarshal(file, &artifacts); err != nil {
        log.Fatalf("Failed to parse artifacts.json file: %v", err)
    }

    for _, artifact := range artifacts {
        if result := database.DB.Create(&artifact); result.Error != nil {
            log.Printf("Failed to insert artifact %s: %v", artifact.Name, result.Error)
        }
    }
}