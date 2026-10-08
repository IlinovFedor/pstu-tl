package usecases

import (
	"context"
	"uuid"

	"github.com/IlinovFedor/pstu-tl/internal/models"
)

type GroupService struct {
	repo GroupRepository
}

func (t *GroupService) GetGroups(ctx context.Context, showDeleted bool) ([]models.Group, error) {
	return t.repo.GetGroups(ctx, showDeleted)
}

func (t *GroupService) CreateGroup(ctx context.Context,
	facultyID uuid.UUID, specializationID uuid.UUID, yearOfEnrollment int, sequenceNumber int, groupType models.GroupType,
) (*models.Group, error) {
	discipline := models.NewGroup(
		uuid.NewV7(),
		facultyID, specializationID, yearOfEnrollment, sequenceNumber, groupType,
		nil,
	)

	return t.repo.InsertGroup(ctx, discipline)
}

func (t *GroupService) GetGroup(ctx context.Context, id uuid.UUID) (*models.Group, error) {
	return t.repo.GetGroup(ctx, id)
}

func (t *GroupService) UpdateGroup(ctx context.Context,
	id uuid.UUID,
	facultyID uuid.UUID, specializationID uuid.UUID, yearOfEnrollment int, sequenceNumber int, groupType models.GroupType,
) (*models.Group, error) {

	return t.repo.EditGroup(ctx, id, facultyID, specializationID, yearOfEnrollment, sequenceNumber, groupType)
}

func (t *GroupService) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	return t.repo.SoftDeleteGroup(ctx, id)
}

func (t *GroupService) RestoreGroup(ctx context.Context, id uuid.UUID) error {
	return t.repo.RestoreGroup(ctx, id)
}

func (t *GroupService) GetHistory(ctx context.Context) ([]models.GroupAction, error) {
	return t.repo.GetGroupsActions(ctx)
}

func (t *GroupService) GetHistoryByID(ctx context.Context, id uuid.UUID) ([]models.GroupAction, error) {
	return t.repo.GetGroupActions(ctx, id)
}

func (t *GroupService) RevertAction(ctx context.Context, id uuid.UUID) error {
	return t.repo.RevertGroupAction(ctx, id)
}
