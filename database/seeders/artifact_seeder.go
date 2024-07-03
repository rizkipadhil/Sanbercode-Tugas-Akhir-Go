package seeders

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func SeedArtifacts() {
    // Baca file JSON
    file, err := ioutil.ReadFile("data/artifacts.json")
    if err != nil {
        log.Fatalf("Failed to read artifacts.json file: %v", err)
    }

    // Parse JSON
    var artifacts []models.Artifact
    if err := json.Unmarshal(file, &artifacts); err != nil {
        log.Fatalf("Failed to parse artifacts.json file: %v", err)
    }

    // Insert ke database
    for _, artifact := range artifacts {
        if result := database.DB.Create(&artifact); result.Error != nil {
            log.Printf("Failed to insert artifact %s: %v", artifact.Name, result.Error)
        }
    }
}