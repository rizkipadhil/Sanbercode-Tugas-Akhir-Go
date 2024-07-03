package models

import "time"

type Character struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string
    ElementID   uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    WeaponID    uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Rarity      int
    BaseStats   string    `gorm:"type:json"`
    Talents     string    `gorm:"type:json"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    CreatedBy   string
    UpdatedBy   string
}