package model

type Student struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`

	Salary float64 `json:"salary"`
}
