package seeders

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func SeedCharacters() {
		var count int64
    database.DB.Model(&models.Character{}).Count(&count)
    if count > 0 {
        log.Println("Seeder sudah dijalankan sebelumnya.")
        return
    }
    file, err := ioutil.ReadFile("data/characters.json")
    if err != nil {
        log.Fatalf("Failed to read characters.json file: %v", err)
    }

    var characters []models.Character
    if err := json.Unmarshal(file, &characters); err != nil {
        log.Fatalf("Failed to parse characters.json file: %v", err)
    }

    for _, character := range characters {
        if result := database.DB.Create(&character); result.Error != nil {
            log.Printf("Failed to insert character %s: %v", character.Name, result.Error)
        }
    }
}