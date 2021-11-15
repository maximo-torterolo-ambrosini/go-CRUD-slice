package models

type Course struct {
	Id    int     `json:"id"`
	Name  string  `json:"course"`
	Price float32 `json:"price"` // in usd
}

type Courses []Course
