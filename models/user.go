package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique"`
    Email     string    `gorm:"unique"`
    Password  string
    Role      string    // Values: "superadmin", "member"
    CreatedAt time.Time
    UpdatedAt time.Time
}