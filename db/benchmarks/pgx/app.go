package main

import (
	"context"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	start := time.Now()

	pgxPool, err := pgxpool.Connect(context.Background(), "user=postgres password=postgres host=127.0.0.1 port=55432 dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	pgxConn, err := pgxPool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer pgxConn.Release()

	gofakeit.Seed(0)

	query := `
		SELECT
		    t.id         AS ticket_id,
		    t.created_at AS ticket_created_at,
		    t.creator_id AS ticket_creator_id,
		    tu.name      AS ticket_creator_name,
		    t.subject    AS ticket_subject,
		    c.created_at AS comment_created_at,
		    c.user_id    AS comment_user_id,
		    cu.name      AS comment_user_name,
		    c.comment    AS comment
		FROM tickets AS t
			JOIN users AS tu ON t.creator_id = tu.id
			LEFT JOIN comments AS c ON c.ticket_id = t.id
			JOIN users AS cu ON c.user_id = cu.id
		WHERE
  			t.id = $1
	`

	type CommentsPgx struct {
		TicketID          uint64    `db:"ticket_id"`
		TicketCreatedAt   time.Time `db:"ticket_created_at"`
		TicketCreatorID   uint64    `db:"ticket_creator_id"`
		TicketCreatorName string    `db:"ticket_creator_name"`
		TicketSubject     string    `db:"ticket_subject"`
		CommentCreatedAt  time.Time `db:"comment_created_at"`
		CommentUserID     uint64    `db:"comment_user_id"`
		CommentUserName   string    `db:"comment_user_name"`
		Comment           string    `db:"comment"`
	}
	j := 0
	for i := 0; i < 100; i++ {
		rows, err := pgxConn.Query(context.Background(), query, gofakeit.Number(1, 100))
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var row CommentsPgx
			if err := rows.Scan(
				&row.TicketID,
				&row.TicketCreatedAt,
				&row.TicketCreatorID,
				&row.TicketCreatorName,
				&row.TicketSubject,
				&row.CommentCreatedAt,
				&row.CommentUserID,
				&row.CommentUserName,
				&row.Comment,
			); err != nil {
				log.Fatal(err)
			}
			//log.Printf("%v", row)
			j++
		}
	}

	log.Println("\n", time.Since(start).Seconds(), j)
}
