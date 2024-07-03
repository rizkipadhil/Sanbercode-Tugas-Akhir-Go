package seeders

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
    "log"
)

func SeedElements() {
	var count int64
    database.DB.Model(&models.Element{}).Count(&count)
    if count > 0 {
        log.Println("Seeder sudah dijalankan sebelumnya.")
        return
    }
    elements := []models.Element{
        {Name: "Pyro", CreatedBy: "seeder"},
        {Name: "Hydro", CreatedBy: "seeder"},
        {Name: "Cryo", CreatedBy: "seeder"},
        {Name: "Geo", CreatedBy: "seeder"},
        {Name: "Anemo", CreatedBy: "seeder"},
        {Name: "Electro", CreatedBy: "seeder"},
        {Name: "Dendro", CreatedBy: "seeder"},
    }

    for _, element := range elements {
        database.DB.Create(&element)
    }
}