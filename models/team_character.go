package models

import "time"

type TeamCharacter struct {
    ID          uint      `gorm:"primaryKey"`
    TeamID      uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    CharacterID uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    ArtifactID  uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    TypeSet     string
    Mechanism   string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    CreatedBy   string
    UpdatedBy   string
}