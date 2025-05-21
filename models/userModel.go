package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique; not null"`
	Password string `json:"password" gorm:"not null"`
	Tasks    []Task // One (user) to many (tasks)
}
