package repository

import (
	"girls/internal/model"

	"github.com/jmoiron/sqlx"
)

type Holiday interface {
	Congratulate(congratulations model.Congratulations) (bool, error)
	ReadCongratulations() ([]model.Congratulations, error)
}

type Repository struct {
	Holiday
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Holiday: NewHolidayRepository(db),
	}
}
