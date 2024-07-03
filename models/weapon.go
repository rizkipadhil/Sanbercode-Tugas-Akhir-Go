package models

import "time"

type Weapon struct {
    ID            uint      `gorm:"primaryKey"`
    Name          string
    CreatedAt     time.Time
    UpdatedAt     time.Time
    CreatedBy     string
    UpdatedBy     string
}