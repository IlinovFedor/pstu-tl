package usecases

import (
	"context"
	"uuid"

	"github.com/IlinovFedor/pstu-tl/internal/models"
)

// CalculationRepository интерфейс для работы с калькуляциями по преподавателям
type CalculationRepository interface {
	GetCalculationsTier1(ctx context.Context, yearName int) (models.Calculation, error)
	GetCalculationsTier2(ctx context.Context, yearName int) (models.Calculation, error)
	GetCalculationsTier3(ctx context.Context, yearName int) (models.Calculation, error)

	GetCalculationsOverall(ctx context.Context, yearName int) (models.Calculation, error)
	GetCalculationsIncomplete(ctx context.Context, yearName int) (models.Calculation, error)
	GetCalculationsErrors(ctx context.Context, yearName int) (models.Calculation, error)

	GetCalculationsTeacher(ctx context.Context, yearName int, teacherID uuid.UUID) (models.Calculation, error)
}

// DisciplineRepository интерфейс для работы с дисциплинами
type DisciplineRepository interface {
	CreateDiscipline(ctx context.Context, discipline models.Discipline) error
	GetDiscipline(ctx context.Context, disciplineID uuid.UUID) (*models.Discipline, error)
	EditDiscipline(ctx context.Context, name string) error

	SoftDeleteDiscipline(ctx context.Context, disciplineID uuid.UUID) error
	RestoreDiscipline(ctx context.Context, disciplineID uuid.UUID) error

	GetDisciplines(ctx context.Context, showDeleted bool) ([]models.Discipline, error)
	GetDisciplinesByIDs(ctx context.Context, recordsIds []uuid.UUID, showDeleted bool) ([]models.Discipline, error)
	GetDisciplinesByYear(ctx context.Context, yearName int) ([]models.Discipline, error)

	GetDisciplinesActions(ctx context.Context) ([]models.DisciplineAction, error)
	GetDisciplineActions(ctx context.Context, disciplineID uuid.UUID) ([]models.DisciplineAction, error)
	RevertDisciplineAction(ctx context.Context, actionID uuid.UUID) error
}

// FileExportRepository интерфейс для работы с файлами экспорта
type FileExportRepository interface {
	CreateFileExport(ctx context.Context, fileExport models.FileExport) error
	GetFileExport(ctx context.Context, fileID uuid.UUID) (*models.FileExport, error)
	GetFileExportData(ctx context.Context, fileID uuid.UUID) ([]byte, error)

	SoftDeleteFileExport(ctx context.Context, fileID uuid.UUID) error
	RestoreFileExport(ctx context.Context, fileID uuid.UUID) error

	GetFileExportsByYearName(ctx context.Context, showDeleted bool, yearName int) ([]models.FileExport, error)

	GetFilesExportsActions(ctx context.Context) ([]models.FileExportAction, error)
	GetFileExportActions(ctx context.Context, fileID uuid.UUID) ([]models.FileExportAction, error)
	RevertFileExportAction(ctx context.Context, actionID uuid.UUID) error
}

// FileImportRepository интерфейс для работы с файлами импорта
type FileImportRepository interface {
	CreateFileImport(ctx context.Context, fileImport models.FileImport) error
	GetFileImport(ctx context.Context, fileID uuid.UUID) (*models.FileImport, error)
	GetFileImportData(ctx context.Context, fileID uuid.UUID) ([]byte, error)

	SoftDeleteFileImport(ctx context.Context, fileID uuid.UUID) error
	RestoreFileImport(ctx context.Context, fileID uuid.UUID) error

	GetFilesImportsByYearName(ctx context.Context, showDeleted bool, yearName int) ([]models.FileImport, error)

	GetFilesImportsActionsByYearName(ctx context.Context, yearName int) ([]models.FileImportAction, error)
	GetFileImportActions(ctx context.Context, fileID uuid.UUID) ([]models.FileImportAction, error)
	RevertFileImportAction(ctx context.Context, actionID uuid.UUID) error
}

// FileReportRepository интерфейс для работы с файлами индивидуальной нагрузки
type FileReportRepository interface {
	CreateFileReport(ctx context.Context, fileReport models.FileReport) error
	GetFileReport(ctx context.Context, fileID uuid.UUID) (*models.FileReport, error)
	GetFileReportData(ctx context.Context, fileID uuid.UUID) ([]byte, error)

	SoftDeleteFileReport(ctx context.Context, fileID uuid.UUID) error
	RestoreFileReport(ctx context.Context, fileID uuid.UUID) error

	GetFileReportsByYearName(ctx context.Context, showDeleted bool, yearName int) ([]models.FileReport, error)
	GetFileReportsDataByYearName(ctx context.Context, showDeleted bool, yearName int) ([][]byte, error)

	GetFilesReportsActionsByYearName(ctx context.Context, yearName int) ([]models.FileReportAction, error)
	GetFileReportActions(ctx context.Context, fileID uuid.UUID) ([]models.FileReportAction, error)
	RevertFileReportAction(ctx context.Context, actionID uuid.UUID) error
}

// GroupRepository интерфейс работы с репозиторием групп
type GroupRepository interface {
	CreateGroup(ctx context.Context, group models.Group) error
	GetGroup(ctx context.Context, groupID uuid.UUID) (*models.Group, error)
	EditGroup(
		ctx context.Context,
		specialization string,
		yearOfEnrollment int,
		sequenceNumber int,
		groupType models.GroupType,
		studentsAmount int,
	) error

	SoftDeleteGroup(ctx context.Context, groupID uuid.UUID) error
	RestoreGroup(ctx context.Context, groupID uuid.UUID) error

	GetGroups(ctx context.Context, showDeleted bool) ([]models.Group, error)
	GetGroupsByIDs(ctx context.Context, recordsIds []uuid.UUID, showDeleted bool) ([]models.Group, error)

	GetGroupsActions(ctx context.Context) ([]models.GroupAction, error)
	GetGroupActions(ctx context.Context, groupID uuid.UUID) ([]models.GroupAction, error)
	RevertGroupAction(ctx context.Context, actionID uuid.UUID) error
}

// RecordsRepository интерфейс для работы с файлами учебных записей
type RecordsRepository interface {
	CreateRecords(ctx context.Context, records []models.Record) error
	GetRecord(ctx context.Context, recordID uuid.UUID) (*models.Record, error)
	EditRecord(ctx context.Context, teachersHours []models.TeacherHours, isError bool) error

	SoftDeleteRecordsByFileID(ctx context.Context, fileID uuid.UUID) error
	RestoreRecordsByFileID(ctx context.Context, fileID uuid.UUID) error

	GetRecordsByYearAndTeacherID(ctx context.Context, yearName int, teacherID uuid.UUID) ([]models.Record, error)
	GetRecordsByYearAndSubjectID(ctx context.Context, yearName int, subjectID uuid.UUID) ([]models.Record, error)

	GetRecordsActionsByYearName(ctx context.Context, yearName int) ([]models.RecordAction, error)
	GetRecordActions(ctx context.Context, recordID uuid.UUID) ([]models.RecordAction, error)
	RevertRecordsByFileID(ctx context.Context, fileID uuid.UUID) error
}

// TeacherRepository интерфейс для работы с преподавателями
type TeacherRepository interface {
	CreateTeacher(ctx context.Context, teacher models.Teacher) error
	GetTeacher(ctx context.Context, teacherID uuid.UUID) (*models.Teacher, error)
	EditTeacher(ctx context.Context,
		name string,
		job models.TeacherJob,
		wage models.TeacherWage,
	) error

	SoftDeleteTeacher(ctx context.Context, teacherID uuid.UUID) error
	RestoreTeacher(ctx context.Context, teacherID uuid.UUID) error

	GetTeachers(ctx context.Context, showDeleted bool) ([]models.Teacher, error)
	GetTeachersByIDs(ctx context.Context, recordsIds []uuid.UUID, showDeleted bool) ([]models.Teacher, error)
	GetTeachersByYear(ctx context.Context, yearName int) ([]models.Teacher, error)

	GetTeachersActions(ctx context.Context) ([]models.TeacherAction, error)
	GetTeacherActions(ctx context.Context, teacherID uuid.UUID) ([]models.TeacherAction, error)
	RevertTeacherAction(ctx context.Context, actionID uuid.UUID) error
}

// YearRepository интерфейс для работы с учебными годами
type YearRepository interface {
	CreateYear(ctx context.Context, year models.Year) error
	GetYear(ctx context.Context, yearID uuid.UUID) (*models.Year, error)
	EditYear(
		ctx context.Context,
		name int,
		wageNormalHours int,
		extHourlyWageTier1 int,
		extHourlyWageTier2 int,
		extHourlyWageTier3 int,
		wageTolerance int,
	) error

	SoftDeleteYear(ctx context.Context, yearID uuid.UUID) error
	RestoreYear(ctx context.Context, yearID uuid.UUID) error

	GetYears(ctx context.Context, showDeleted bool) ([]models.Year, error)
	GetYearsByIDs(ctx context.Context, recordsIds []uuid.UUID, showDeleted bool) ([]models.Year, error)

	GetYearsActions(ctx context.Context) ([]models.YearAction, error)
	GetYearActions(ctx context.Context, yearID uuid.UUID) ([]models.YearAction, error)
	RevertYearAction(ctx context.Context, actionID uuid.UUID) error
}

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	GetUsers(ctx context.Context, showDeleted bool) ([]models.User, error)
	GetUserByName(ctx context.Context, name string) (*models.User, error)

	CreateUser(ctx context.Context, user models.User) error
	CreateUsers(ctx context.Context, users []models.User) error
	UpdateUser(ctx context.Context, userName string, role models.UserRole, associatedTeacherID uuid.UUID) error
	DeleteUser(ctx context.Context, userName string) error

	ChangePassword(ctx context.Context, userName string, passwordHash string) error
}
