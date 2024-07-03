package models

import (
    "time"
)

type Team struct {
    ID             uint            `gorm:"primary_key"`
    Name           string          `json:"name"`
    Description    string          `json:"description"`
    VerifyBy       string          `json:"verify_by"`
    VerifyStatus   string          `json:"verify_status"` // diterima, ditolak, pending
    CreatedAt      time.Time       `json:"created_at"`
    UpdatedAt      time.Time       `json:"updated_at"`
    CreatedBy      string          `json:"created_by"`
    UpdatedBy      string          `json:"updated_by"`

    TeamCharacters []TeamCharacter `json:"team_characters" gorm:"foreignkey:TeamID;constraint:OnDelete:CASCADE;"`
}