package models

type User struct {
	Name  string `json:"name" gorm:"column:name;unique"`
	Email string `json:"email" gorm:"column:email;unique"`
}
