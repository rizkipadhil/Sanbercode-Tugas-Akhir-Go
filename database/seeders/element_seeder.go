package seeders

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
)

func SeedElements() {
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