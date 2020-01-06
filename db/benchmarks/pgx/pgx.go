package pgx

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
)

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

func selectTenCommentsPgx(pgxConn *pgxpool.Conn) (comments []*CommentsPgx, err error) {
	gofakeit.Seed(0)
	n := gofakeit.Number(1, 100)

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

	rows, err := pgxConn.Query(context.Background(), query, n)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		comments = append(comments, &row)
	}

	return comments, nil
}
