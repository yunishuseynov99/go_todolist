package statistics_service

import (
	"context"
	"github.com/yunishuseynov99/go_todolist/internal/core/domain"
	"time"
)

type StatisticsService struct {
	StatisticsRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetTasks(
		ctx context.Context,
		userId *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

func NewStatisticsService(StatisticsRepository StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		StatisticsRepository: StatisticsRepository,
	}
}
