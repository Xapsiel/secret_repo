package service

import (
	"girls/internal/model"
	"girls/internal/repository"
)

type Service struct {
	Holiday
}

type Holiday interface {
	WriteCongratulations(congratulations model.Congratulations) (bool, error)
	ReadCongratulations() ([]model.Congratulations, error)
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Holiday: NewHolidayService(repo.Holiday),
	}
}
