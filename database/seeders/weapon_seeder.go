package seeders

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
    "log"
)

func SeedWeapons() {
		var count int64
    database.DB.Model(&models.Weapon{}).Count(&count)
    if count > 0 {
        log.Println("Seeder sudah dijalankan sebelumnya.")
        return
    }	
    weapons := []models.Weapon{
        {Name: "Bow", CreatedBy: "seeder"},
        {Name: "Catalyst", CreatedBy: "seeder"},
        {Name: "Claymore", CreatedBy: "seeder"},
        {Name: "Polearm", CreatedBy: "seeder"},
        {Name: "Sword", CreatedBy: "seeder"},
    }

    for _, weapon := range weapons {
        database.DB.Create(&weapon)
    }
}