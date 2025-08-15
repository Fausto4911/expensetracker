package dto

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Expense struct {
	Id          uint16             `json:"id"`
	Amount      float32            `json:"amount"`
	Category_id uint16             `json:"categoryId"`
	Created     pgtype.Timestamptz `json:"created"`
	Modified    pgtype.Timestamptz `json:"modified"`
}
