package models

type Item struct {
	Model
	Name         string
	HealthPoints int `gorm:"not null"`
	ManaPoints   int `gorm:"not null"`
	Strength     int `gorm:"not null"`
	Agility      int `gorm:"not null"`
	Intelligence int `gorm:"not null"`
	ManaCost     int `gorm:"not null"`
	Description  string
}
