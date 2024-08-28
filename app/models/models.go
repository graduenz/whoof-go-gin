package models

import (
	"time"
)

type PetType int

const (
	Dog PetType = iota
	Cat
	Capybara
	Other
)

// BaseModel TODO: Maybe replace this model by gorm.Model and switch to DTOs instead of using this model in JSON
type BaseModel struct {
	ID        int       `gorm:"column:id; primary_key; not null;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at; not null;" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at; not null;" json:"-"`
	DeletedAt time.Time `gorm:"column:deleted_at; not null;" json:"-"`
}

type Pet struct {
	BaseModel
	Name    string  `gorm:"column:name; not null;" json:"name"`
	PetType PetType `gorm:"column:pet_type; not null;" json:"pet_type"`
}

//type Vaccine struct {
//	gorm.Model
//	Name        string
//	Description string
//	PetType     PetType
//	Duration    int
//}
//
//type PetVaccination struct {
//	gorm.Model
//	PetId     uint
//	VaccineId uint
//	AppliedAt time.Time
//}
