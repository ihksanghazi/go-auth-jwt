package models

type Users struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"varchar(300)" json:"name"`
	Email    string `gorm:"varchar(300)" json:"email"`
	Password string `gorm:"varchar(300)" json:"password"`
}