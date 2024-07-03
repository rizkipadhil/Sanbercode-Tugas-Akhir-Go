package seeders

import (
    "tugas-akhir/database"
    "tugas-akhir/models"
    "golang.org/x/crypto/bcrypt"
    "log"
)

func SeedUsers() {
    var count int64
    database.DB.Model(&models.User{}).Count(&count)
    if count > 0 {
        log.Println("Seeder sudah dijalankan sebelumnya.")
        return
    }

    passwordAdmin, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
    if err != nil {
        log.Fatal(err)
    }
    passwordMember, err := bcrypt.GenerateFromPassword([]byte("member123"), bcrypt.DefaultCost)
    if err != nil {
        log.Fatal(err)
    }

    users := []models.User{
        {
            Username: "superadmin",
            Email: "superadmin@example.com",
            Password: string(passwordAdmin),
            Role: "superadmin",
        },
        {
            Username: "member",
            Email: "member@example.com",
            Password: string(passwordMember),
            Role: "member",
        },
    }

    for _, user := range users {
        database.DB.Create(&user)
    }
}