package models

import (
    "time"
)

type TeamCharacter struct {
    ID          uint       `gorm:"primary_key"`
    TeamID      uint       `json:"team_id"`
    CharacterID uint       `json:"character_id"`
    ArtifactID  uint       `json:"artifact_id"`
    TypeSet     string     `json:"type_set"`
    Mechanism   string     `json:"mechanism"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
    CreatedBy   string     `json:"created_by"`
    UpdatedBy   string     `json:"updated_by"`

    Character   Character  `json:"character" gorm:"foreignkey:CharacterID"`
    Artifact    Artifact   `json:"artifact" gorm:"foreignkey:ArtifactID"`
}