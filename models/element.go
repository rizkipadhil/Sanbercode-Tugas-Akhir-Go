package models

import "time"

type Element struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    CreatedBy   string
    UpdatedBy   string
}