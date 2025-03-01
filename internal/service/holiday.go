package service

import (
	"girls/internal/model"
	"girls/internal/repository"
)

type HolidayService struct {
	repo repository.Holiday
}

func NewHolidayService(repo repository.Holiday) *HolidayService {
	return &HolidayService{repo: repo}
}
func (h *HolidayService) WriteCongratulations(congratulations model.Congratulations) (bool, error) {
	return h.repo.Congratulate(congratulations)

}

func (h *HolidayService) ReadCongratulations() ([]model.Congratulations, error) {
	return h.repo.ReadCongratulations()

}
