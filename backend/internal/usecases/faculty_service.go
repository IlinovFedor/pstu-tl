package usecases

import (
	"context"
	"uuid"

	"github.com/IlinovFedor/pstu-tl/internal/models"
)

type FacultyService struct {
	repo FacultyRepository
}

func (t *FacultyService) GetFaculties(ctx context.Context, showDeleted bool) ([]models.Faculty, error) {
	return t.repo.GetFaculties(ctx, showDeleted)
}

func (t *FacultyService) CreateFaculty(ctx context.Context,
	name string,
) (*models.Faculty, error) {
	faculty := models.NewFaculty(
		uuid.NewV7(),
		name,
		nil,
	)

	return t.repo.InsertFaculty(ctx, faculty)
}

func (t *FacultyService) GetFaculty(ctx context.Context, id uuid.UUID) (*models.Faculty, error) {
	return t.repo.GetFaculty(ctx, id)
}

func (t *FacultyService) UpdateFaculty(ctx context.Context,
	id uuid.UUID,
	name string,
) (*models.Faculty, error) {

	return t.repo.EditFaculty(ctx, id, name)
}

func (t *FacultyService) DeleteFaculty(ctx context.Context, id uuid.UUID) error {
	return t.repo.SoftDeleteFaculty(ctx, id)
}

func (t *FacultyService) RestoreFaculty(ctx context.Context, id uuid.UUID) error {
	return t.repo.RestoreFaculty(ctx, id)
}

func (t *FacultyService) GetHistory(ctx context.Context) ([]models.FacultyAction, error) {
	return t.repo.GetFacultiesActions(ctx)
}

func (t *FacultyService) GetHistoryByID(ctx context.Context, id uuid.UUID) ([]models.FacultyAction, error) {
	return t.repo.GetFacultyActions(ctx, id)
}

func (t *FacultyService) RevertAction(ctx context.Context, id uuid.UUID) error {
	return t.repo.RevertFacultyAction(ctx, id)
}
