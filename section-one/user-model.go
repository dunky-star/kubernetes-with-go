package main

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"username" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Age   int    `json:"age"`
}

// UserParam struct serves as a data transfer object (DTO) specifically for input handling, often seen in web APIs or web forms.
type UserParam struct {
	Name  string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   int    `json:"age"`
}
