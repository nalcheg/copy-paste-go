package gorm_model

func (User) TableName() string {
	return "users"
}

type User struct {
	ID   string `gorm:"PRIMARY_KEY"`
	Name string
}
