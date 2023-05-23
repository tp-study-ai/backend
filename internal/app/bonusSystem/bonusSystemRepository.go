package bonusSystem

import "time"

type Repository interface {
	GetCountTaskOfDate(id int, day time.Time) (int, error)
}
