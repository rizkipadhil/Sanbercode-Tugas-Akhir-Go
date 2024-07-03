package models

import "time"

type Team struct {
    ID           uint      `gorm:"primaryKey"`
    Name         string
    Description  string
    VerifyBy     uint      `gorm:"default:null"` // ID of the user who verified the team
    VerifyStatus string    // Values: "diterima", "ditolak"
    CreatedAt    time.Time
    UpdatedAt    time.Time
    CreatedBy    string
    UpdatedBy    string
}