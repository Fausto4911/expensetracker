package dto

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ExpenseResponse struct {
	Id          uint16             `json:"id"`
	Amount      float32            `json:"amount"`
	Category_id uint16             `json:"categoryId"`
	Created     pgtype.Timestamptz `json:"created"`
	Modified    pgtype.Timestamptz `json:"modified"`
}

type Expense struct {
	id          uint16
	amount      float32
	category_id uint16
	created     pgtype.Timestamptz
	modified    pgtype.Timestamptz
}

func (e *Expense) Id() uint16 {
	return e.id
}

func (e *Expense) SetId(id uint16) {
	e.id = id
}

func (e *Expense) Amount() float32 {
	return e.amount
}

func (e *Expense) SetAmount(amount float32) {
	e.amount = amount
}

func (e *Expense) CategoryId() uint16 {
	return e.category_id
}

func (e *Expense) SetCategoryId(category_id uint16) {
	e.category_id = category_id
}

func (e *Expense) Created() pgtype.Timestamptz {
	return e.created
}

func (e *Expense) SetCreated(created pgtype.Timestamptz) {
	e.created = created
}

func (e *Expense) Modified() pgtype.Timestamptz {
	return e.modified
}

func (e *Expense) SetModified(modified pgtype.Timestamptz) {
	e.modified = modified
}

func BuildExpenseResponse(e Expense) ExpenseResponse {
	return ExpenseResponse{Id: e.Id(), Amount: e.Amount(), Category_id: e.CategoryId(), Created: e.Created(), Modified: e.Modified()}
}
