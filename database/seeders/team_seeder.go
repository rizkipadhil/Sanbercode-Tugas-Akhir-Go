package seeders

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func SeedTeams() {
		var count int64
    database.DB.Model(&models.Team{}).Count(&count)
    if count > 0 {
        log.Println("Seeder sudah dijalankan sebelumnya di team.")
        return
    }

    file, err := ioutil.ReadFile("data/teams.json")
    if err != nil {
        log.Fatalf("Failed to read teams.json file: %v", err)
    }

    var teams []struct {
        Name           string                  `json:"name"`
        Description    string                  `json:"description"`
        VerifyBy       string                  `json:"verify_by"`
        VerifyStatus   string                  `json:"verify_status"`
        TeamCharacters []models.TeamCharacter  `json:"team_characters"`
    }
    if err := json.Unmarshal(file, &teams); err != nil {
        log.Fatalf("Failed to parse teams.json file: %v", err)
    }

    for _, teamData := range teams {
        team := models.Team{
            Name:          teamData.Name,
            Description:   teamData.Description,
            VerifyBy:      teamData.VerifyBy,
            VerifyStatus:  teamData.VerifyStatus,
        }

        tx := database.DB.Begin()
        if err := tx.Create(&team).Error; err != nil {
            tx.Rollback()
            log.Printf("Failed to create team %s: %v", team.Name, err)
            continue
        }

        for _, teamCharacter := range teamData.TeamCharacters {
            teamCharacter.TeamID = team.ID

            if err := tx.Create(&teamCharacter).Error; err != nil {
                tx.Rollback()
                log.Printf("Failed to create team character for team %s: %v", team.Name, err)
                break
            }
        }

        if err := tx.Commit().Error; err != nil {
            log.Printf("Failed to commit transaction for team %s: %v", team.Name, err)
        }
    }
}