package bonusSystemRepository

import (
	"github.com/jackc/pgx"
	"time"
)

type RepositoryBonusSystem struct {
	DB *pgx.ConnPool
}

func NewRepositoryBonusSystem(db *pgx.ConnPool) *RepositoryBonusSystem {
	return &RepositoryBonusSystem{DB: db}
}

func (r *RepositoryBonusSystem) GetCountTaskOfDate(id int, day time.Time) (int, error) {
	var countTask int
	sql := `select count(*) from send_task where "user_id" = $1 and $2 < date::date + interval '3 hours' and date::date + interval '3 hours' < $3 and "tests_passed" = "tests_total";`
	err := r.DB.QueryRow(sql, id, day.Format("2006-01-02"), day.Add(+24*time.Hour).Format("2006-01-02")).Scan(
		&countTask,
	)
	if err != nil {
		return 0, err
	}
	return countTask, nil
}
