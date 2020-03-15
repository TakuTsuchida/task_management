package model

import "time"

type Circle struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"username"`
	Use       string `json:"use"`
	TypeId    int
	Users     []User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Circles []Circle

func CreateCircle(circle *Circle) bool {
	err := db.Create(circle).Error
	if err != nil {
		return false
	}
	return true
}

func FindCircle(c *Circle) Circle {
	var circle Circle
	db.Where(c).First(&circle)
	return circle
}

func AllCircles() Circles {
	var circles Circles
	db.Find(&circles)
	return circles
}
