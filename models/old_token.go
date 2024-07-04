package models

import "time"

type OldToken struct {
    ID        uint      `gorm:"primaryKey"`
    Token     string    `gorm:"type:text"`
    CreatedAt time.Time
}