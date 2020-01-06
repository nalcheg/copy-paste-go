package gorm_model

import "time"

func (Ticket) TableName() string {
	return "tickets"
}

type Ticket struct {
	ID        string `gorm:"PRIMARY_KEY"`
	CreatorID uint64
	Creator   User `gorm:"foreignkey:id;association_foreignkey:creator_id"`
	Subject   string
	CreatedAt time.Time
	Comments  []Comment `gorm:"foreignkey:TicketID"`
}
