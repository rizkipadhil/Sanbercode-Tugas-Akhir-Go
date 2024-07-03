package models

import "time"

type Weapon struct {
    ID            uint      `gorm:"primaryKey"`
    Name          string
    Type          string
    Rarity        int
    BaseAttack    int
    SecondaryStat string
    CreatedAt     time.Time
    UpdatedAt     time.Time
    CreatedBy     string
    UpdatedBy     string
}