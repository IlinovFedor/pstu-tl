package usecases

import (
	"context"
	"uuid"

	"github.com/IlinovFedor/pstu-tl/internal/models"
)

type YearService struct {
	repo YearRepository
}

func (t *YearService) GetYears(ctx context.Context, showDeleted bool) ([]models.Year, error) {
	return t.repo.GetYears(ctx, showDeleted)
}

func (t *YearService) CreateYear(ctx context.Context,
	name int, wageNormalHours int, extHourlyWageTier1 int, extHourlyWageTier2 int, extHourlyWageTier3 int, wageTolerance int,
) (*models.Year, error) {
	year := models.NewYear(
		name,
		wageNormalHours,
		extHourlyWageTier1,
		extHourlyWageTier2,
		extHourlyWageTier3,
		wageTolerance,
		nil,
	)

	return t.repo.InsertYear(ctx, year)
}

func (t *YearService) GetYear(ctx context.Context, name int) (*models.Year, error) {
	return t.repo.GetYear(ctx, name)
}

func (t *YearService) UpdateYear(ctx context.Context,
	name int, wageNormalHours int, extHourlyWageTier1 int, extHourlyWageTier2 int, extHourlyWageTier3 int, wageTolerance int,
) (*models.Year, error) {

	return t.repo.EditYear(ctx, name, wageNormalHours, extHourlyWageTier1, extHourlyWageTier2, extHourlyWageTier3, wageTolerance)
}

func (t *YearService) DeleteYear(ctx context.Context, name int) error {
	return t.repo.SoftDeleteYear(ctx, name)
}

func (t *YearService) RestoreYear(ctx context.Context, name int) error {
	return t.repo.RestoreYear(ctx, name)
}

func (t *YearService) GetHistory(ctx context.Context) ([]models.YearAction, error) {
	return t.repo.GetYearsActions(ctx)
}

func (t *YearService) GetHistoryByName(ctx context.Context, name int) ([]models.YearAction, error) {
	return t.repo.GetYearActions(ctx, name)
}

func (t *YearService) RevertAction(ctx context.Context, id uuid.UUID) error {
	return t.repo.RevertYearAction(ctx, id)
}
