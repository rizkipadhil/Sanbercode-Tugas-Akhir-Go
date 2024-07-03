package models

import (
    "encoding/json"
    "time"
    "database/sql/driver"
)

type Description struct {
    Set2 string `json:"2set"`
    Set4 string `json:"4set"`
}

type Artifact struct {
    ID          uint       `gorm:"primary_key"`
    Name        string     `json:"name"`
    Description Description `json:"description" gorm:"type:json"`
    Rarity      int        `json:"rarity"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
    CreatedBy   string     `json:"created_by"`
    UpdatedBy   string     `json:"updated_by"`
}

func (d Description) Value() (driver.Value, error) {
    return json.Marshal(d)
}

func (d *Description) Scan(value interface{}) error {
    return json.Unmarshal(value.([]byte), &d)
}