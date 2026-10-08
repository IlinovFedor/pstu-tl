package usecases

import (
	"context"
	"uuid"

	"github.com/IlinovFedor/pstu-tl/internal/models"
)

type SpecializationService struct {
	repo SpecializationRepository
}

func (t *SpecializationService) GetSpecializations(ctx context.Context, showDeleted bool) ([]models.Specialization, error) {
	return t.repo.GetSpecializations(ctx, showDeleted)
}

func (t *SpecializationService) CreateSpecialization(ctx context.Context,
	name string,
) (*models.Specialization, error) {
	specialization := models.NewSpecialization(
		uuid.NewV7(),
		name,
		nil,
	)

	return t.repo.InsertSpecialization(ctx, specialization)
}

func (t *SpecializationService) GetSpecialization(ctx context.Context, id uuid.UUID) (*models.Specialization, error) {
	return t.repo.GetSpecialization(ctx, id)
}

func (t *SpecializationService) UpdateSpecialization(ctx context.Context,
	id uuid.UUID,
	name string,
) (*models.Specialization, error) {

	return t.repo.EditSpecialization(ctx, id, name)
}

func (t *SpecializationService) DeleteSpecialization(ctx context.Context, id uuid.UUID) error {
	return t.repo.SoftDeleteSpecialization(ctx, id)
}

func (t *SpecializationService) RestoreSpecialization(ctx context.Context, id uuid.UUID) error {
	return t.repo.RestoreSpecialization(ctx, id)
}

func (t *SpecializationService) GetHistory(ctx context.Context) ([]models.SpecializationAction, error) {
	return t.repo.GetSpecializationsActions(ctx)
}

func (t *SpecializationService) GetHistoryByID(ctx context.Context, id uuid.UUID) ([]models.SpecializationAction, error) {
	return t.repo.GetSpecializationActions(ctx, id)
}

func (t *SpecializationService) RevertAction(ctx context.Context, id uuid.UUID) error {
	return t.repo.RevertSpecializationAction(ctx, id)
}
