package seeders

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func SeedCharacters() {
    // Baca file JSON
    file, err := ioutil.ReadFile("data/characters.json")
    if err != nil {
        log.Fatalf("Failed to read characters.json file: %v", err)
    }

    // Parse JSON
    var characters []models.Character
    if err := json.Unmarshal(file, &characters); err != nil {
        log.Fatalf("Failed to parse characters.json file: %v", err)
    }

    // Insert ke database
    for _, character := range characters {
        if result := database.DB.Create(&character); result.Error != nil {
            log.Printf("Failed to insert character %s: %v", character.Name, result.Error)
        }
    }
}