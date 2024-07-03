package seeders

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func SeedWeapons() {
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