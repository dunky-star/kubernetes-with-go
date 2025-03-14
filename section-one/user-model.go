package main

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"username" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Age   int    `json:"age"`
}

type UserParam struct {
	Name  string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   int    `json:"age"`
}
