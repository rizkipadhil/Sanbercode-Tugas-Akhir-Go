package models

import (
    "time"
)

type Character struct {
    ID          uint      `gorm:"primary_key"`
    Name        string    `json:"name"`
    ElementID   uint      `json:"element_id"`
    WeaponID    uint      `json:"weapon_id"`
    Rarity      int       `json:"rarity"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    CreatedBy   string    `json:"created_by"`
    UpdatedBy   string    `json:"updated_by"`

    Element     Element   `json:"element" gorm:"foreignkey:ElementID"`
    Weapon      Weapon    `json:"weapon" gorm:"foreignkey:WeaponID"`
}