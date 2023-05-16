package bonusSystemUseCase

import (
	"fmt"
	"github.com/tp-study-ai/backend/internal/app/bonusSystem"
	"github.com/tp-study-ai/backend/internal/app/models"
	"time"
)

type UseCaseBonusSystem struct {
	Repo bonusSystem.Repository
}

func NewUseCaseBonusSystem(TaskRepo bonusSystem.Repository) *UseCaseBonusSystem {
	return &UseCaseBonusSystem{
		Repo: TaskRepo,
	}
}

func (u *UseCaseBonusSystem) GetCountTaskOfDate(id int) (*models.Days, error) {
	now := time.Now()
	days := &models.Days{}
	for i := 0; i < 365; i++ {
		task, _ := u.Repo.GetCountTaskOfDate(id, now)
		days.Days = append(days.Days, models.Day{Day: now, Count: task})
		now = now.Add(-24 * time.Hour)
	}

	return days, nil
}

func (u *UseCaseBonusSystem) GetShockMode(id int) (*models.ShockMode, error) {
	shockMode := &models.ShockMode{}
	now := time.Now()
	a := 0
	for i := 0; i < 365; i++ {
		count, err := u.Repo.GetCountTaskOfDate(id, now)
		if count == 0 {
			break
		}
		if err != nil {
			return nil, err
		}
		a += 1
		now = now.Add(-24 * time.Hour)
	}

	if a != 0 {
		shockMode.Today = true
		shockMode.ShockMode = a
		return shockMode, nil
	} else {
		fmt.Println()
		now = now.Add(-24 * time.Hour)
		for i := 0; i < 60; i++ {
			count, err := u.Repo.GetCountTaskOfDate(id, now)
			if err != nil {
				return nil, err
			}
			if count == 0 {
				break
			}
			a += 1
			now = now.Add(-24 * time.Hour)
		}
	}

	shockMode.Today = false
	shockMode.ShockMode = a

	return shockMode, nil
}
