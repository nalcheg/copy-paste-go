package pgx

import (
	"github.com/brianvoe/gofakeit"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	gormModel "github.com/nalcheg/copy-paste-go/db/benchmarks/gorm-model"
)

func selectTenCommentsGorm(db *gorm.DB) []*gormModel.Comment {
	gofakeit.Seed(0)
	n := gofakeit.Number(1, 100)

	var comments []*gormModel.Comment
	db.Preload("Ticket").
		Preload("User").
		Preload("Ticket.Creator").
		Where("ticket_id = ?", n).
		Find(&comments)

	return comments
}
