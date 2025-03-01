package repository

import (
	"fmt"

	"girls/internal/model"

	"github.com/jmoiron/sqlx"
)

type HolidayRepository struct {
	db *sqlx.DB
}

func NewHolidayRepository(db *sqlx.DB) *HolidayRepository {
	return &HolidayRepository{db: db}
}

func (h *HolidayRepository) Congratulate(congratulations model.Congratulations) (bool, error) {
	query := `INSERT INTO holidays(text, author, telegram_name)
          VALUES($1, $2, $3)
          ON CONFLICT (telegram_name) DO UPDATE 
          SET author = EXCLUDED.author,  text=EXCLUDED.text`
	_, err := h.db.Exec(query, congratulations.Text, congratulations.NickName, congratulations.TelegramName)
	if err != nil {
		return false, fmt.Errorf("Problem with insert data: %v", congratulations)
	}
	return true, nil
}

func (h *HolidayRepository) ReadCongratulations() ([]model.Congratulations, error) {
	query := `SELECT text, author,telegram_name FROM holidays`
	res, err := h.db.Query(query)
	if err != nil {
		return []model.Congratulations{}, fmt.Errorf("Problem with read data: %v", err)
	}
	defer res.Close()
	var congratulations []model.Congratulations
	var congratulation model.Congratulations
	for res.Next() {
		err := res.Scan(&congratulation.Text, &congratulation.NickName, &congratulation.TelegramName)
		if err != nil {
			return []model.Congratulations{}, fmt.Errorf("Problem with read data: %v", err)
		}
		congratulations = append(congratulations, congratulation)
	}
	return congratulations, nil
}
