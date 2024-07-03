package models

import "time"

type Artifact struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string
    Description string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    CreatedBy   string
    UpdatedBy   string
}