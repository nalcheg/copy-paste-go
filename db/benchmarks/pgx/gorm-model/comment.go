package gorm_model

import "time"

func (Comment) TableName() string {
	return "comments"
}

type Comment struct {
	ID        string `gorm:"PRIMARY_KEY"`
	TicketID  uint64
	Ticket    Ticket `gorm:"foreignkey:id;association_foreignkey:ticket_id"`
	UserID    uint64
	User      User `gorm:"foreignkey:id;association_foreignkey:user_id"`
	Comment   string
	CreatedAt time.Time
}
