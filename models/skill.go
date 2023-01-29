package models

type Skill struct {
	Model
	Name        string `gorm:"not null;unique_index"`
	ManaCost    int    `gorm:"not null"`
	Description string
	HeroId      uint `gorm:"not null"`
}
