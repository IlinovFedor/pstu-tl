package usecases

import (
	"context"
	"uuid"

	"github.com/IlinovFedor/pstu-tl/internal/models"
)

type DisciplineService struct {
	repo DisciplineRepository
}

func (t *DisciplineService) GetDisciplines(ctx context.Context, showDeleted bool) ([]models.Discipline, error) {
	return t.repo.GetDisciplines(ctx, showDeleted)
}

func (t *DisciplineService) CreateDiscipline(ctx context.Context,
	name string,
) (*models.Discipline, error) {
	discipline := models.NewDiscipline(
		uuid.NewV7(),
		name,
		nil,
	)

	return t.repo.InsertDiscipline(ctx, discipline)
}

func (t *DisciplineService) GetDiscipline(ctx context.Context, id uuid.UUID) (*models.Discipline, error) {
	return t.repo.GetDiscipline(ctx, id)
}

func (t *DisciplineService) UpdateDiscipline(ctx context.Context,
	id uuid.UUID,
	name string,
) (*models.Discipline, error) {

	return t.repo.EditDiscipline(ctx, id, name)
}

func (t *DisciplineService) DeleteDiscipline(ctx context.Context, id uuid.UUID) error {
	return t.repo.SoftDeleteDiscipline(ctx, id)
}

func (t *DisciplineService) RestoreDiscipline(ctx context.Context, id uuid.UUID) error {
	return t.repo.RestoreDiscipline(ctx, id)
}

func (t *DisciplineService) GetHistory(ctx context.Context) ([]models.DisciplineAction, error) {
	return t.repo.GetDisciplinesActions(ctx)
}

func (t *DisciplineService) GetHistoryByID(ctx context.Context, id uuid.UUID) ([]models.DisciplineAction, error) {
	return t.repo.GetDisciplineActions(ctx, id)
}

func (t *DisciplineService) RevertAction(ctx context.Context, id uuid.UUID) error {
	return t.repo.RevertDisciplineAction(ctx, id)
}
