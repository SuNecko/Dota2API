package models

type Hero struct {
	Model
	Name         string `gorm:"not null; unique_index"`
	Attribute    string `gorm:"not null"`
	HealthPoints int    `gorm:"not null"`
	ManaPoints   int    `gorm:"not null"`
	Strength     int    `gorm:"not null"`
	Agility      int    `gorm:"not null"`
	Intelligence int    `gorm:"not null"`
	Skills       []Skill
}
