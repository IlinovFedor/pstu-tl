package usecases

import (
	"context"
	"uuid"

	"github.com/IlinovFedor/pstu-tl/internal/models"
)

type TeacherService struct {
	repo TeacherRepository
}

func (t *TeacherService) GetTeachers(ctx context.Context, showDeleted bool) ([]models.Teacher, error) {
	return t.repo.GetTeachers(ctx, showDeleted)
}

func (t *TeacherService) CreateTeacher(ctx context.Context,
	name string,
	job models.TeacherJob,
	wage models.TeacherWage) (*models.Teacher, error) {
	teacher := models.NewTeacher(
		uuid.NewV7(),
		name,
		job,
		wage,
		nil,
	)

	return t.repo.InsertTeacher(ctx, teacher)
}

func (t *TeacherService) GetTeacher(ctx context.Context, id uuid.UUID) (*models.Teacher, error) {
	return t.repo.GetTeacher(ctx, id)
}

func (t *TeacherService) UpdateTeacher(ctx context.Context,
	id uuid.UUID,
	name string,
	job models.TeacherJob,
	wage models.TeacherWage) (*models.Teacher, error) {

	return t.repo.EditTeacher(ctx, id, name, job, wage)
}

func (t *TeacherService) DeleteTeacher(ctx context.Context, id uuid.UUID) error {
	return t.repo.SoftDeleteTeacher(ctx, id)
}

func (t *TeacherService) RestoreTeacher(ctx context.Context, id uuid.UUID) error {
	return t.repo.RestoreTeacher(ctx, id)
}

func (t *TeacherService) GetHistory(ctx context.Context) ([]models.TeacherAction, error) {
	return t.repo.GetTeachersActions(ctx)
}

func (t *TeacherService) GetHistoryByID(ctx context.Context, id uuid.UUID) ([]models.TeacherAction, error) {
	return t.repo.GetTeacherActions(ctx, id)
}

func (t *TeacherService) RevertAction(ctx context.Context, id uuid.UUID) error {
	return t.repo.RevertTeacherAction(ctx, id)
}
