// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transfer.sql

package db

import (
	"context"
	"time"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
from_account_id,
to_account_id,
amount,
created_at
) VALUES (
    $1,$2,$3,$4
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
		arg.CreatedAt,
	)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers WHERE id=$1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
