package dto

import (
	"time"
)

type Expense struct {
	id          uint16
	amount      float32
	category_id uint16
	created     time.Time
	modified    time.Time
}

func (e *Expense) Id() uint16 {
	return e.id
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

func (e *Expense) Created() time.Time {
	return e.created
}

func (e *Expense) SetCreated(created time.Time) {
	e.created = created
}

func (e *Expense) Modified() time.Time {
	return e.modified
}

func (e *Expense) SetModified(modified time.Time) {
	e.modified = modified
}
