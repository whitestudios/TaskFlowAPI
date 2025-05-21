package models

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	SubTitle  string `json:"subtitle"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id"` // foreign key
}
