package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/utils"
)

type (
	AccountsServiceInterface interface {
		CreateUser(ctx context.Context,
			input AddAccountInput,
			passwordGen utils.StrGenFunc,
			accountsRepo *repository.Repository[models.Account],
		) (*models.Account, error)

		EditAccount(ctx context.Context,
			input EditAccountInput,
			accountsRepo *repository.Repository[models.Account],
		) (*models.Account, error)

		LoginUser(ctx context.Context,
			input LoginUserInput,
			accountsRepo *repository.Repository[models.Account],
		) (*models.Account, error)

		ForgotPassword(ctx context.Context,
			input ForgotPasswordInput,
			accountsRepo *repository.Repository[models.Account],
		) (*models.Account, error)

		ResetPassword(ctx context.Context,
			input ResetPasswordInput,
			accountsRepo *repository.Repository[models.Account],
		) (*models.Account, error)

		ListAccounts(ctx context.Context,
			input ListAccountReportsInput,
			accountsRepo *repository.Repository[models.Account],
		) ([]models.Account, *repository.Paginator, error)
	}

	FractionationServiceInterface interface {
		CreateFractionationRecord(
			ctx context.Context,
			input CreateFractionationRecordInput,
			fractionationRepo *repository.Repository[models.FractionationReport],
		) (*models.FractionationReport, error)

		UpdateFractionationRecord(
			ctx context.Context,
			input UpdateFractionationRecordInput,
			fractionationRepo *repository.Repository[models.FractionationReport],
		) (*models.FractionationReport, error)

		GetFractionationRecord(
			ctx context.Context,
			input GetFractionationRecordInput,
			fractionationRepo *repository.Repository[models.FractionationReport],
		) (*models.FractionationReport, error)

		ListFractionationRecords(
			ctx context.Context,
			input ListFractionationReportsInput,
			fractionationRepo *repository.Repository[models.FractionationReport],
		) ([]models.FractionationReport, *repository.Paginator, error)
	}

	QualityControlServiceInterface interface {
		CreateQualityRecord(
			ctx context.Context,
			input CreateQualityRecordInput,
			dailyQualityRepo *repository.Repository[models.DailyQualityReadings],
		) (*models.DailyQualityReadings, error)

		GetDailyQualityRecord(
			ctx context.Context,
			input GetQualityRecordInput,
			dailyQualityRepo *repository.Repository[models.DailyQualityReadings],
		) (*models.DailyQualityReadings, error)

		ListQualityRecords(
			ctx context.Context,
			input ListQualityReportsInput,
			dailyQualityRepo *repository.Repository[models.DailyQualityReadings],
		) ([]models.DailyQualityReadings, *repository.Paginator, error)
	}

	RefineryServiceInterface interface {
		CreateRefineryRecord(
			ctx context.Context,
			input CreateRefineryInput,
			refineryRepo *repository.Repository[models.RefineryReport],
		) (*models.RefineryReport, error)

		UpdateRefineryRecord(
			ctx context.Context,
			input UpdateRefineryRecordInput,
			refineryRepo *repository.Repository[models.RefineryReport],
		) (*models.RefineryReport, error)

		GetRefineryRecord(
			ctx context.Context,
			input GetRefineryRecordInput,
			refineryRepo *repository.Repository[models.RefineryReport],
		) (*models.RefineryReport, error)

		ListRefineryRecords(
			ctx context.Context,
			input ListRefineryReportsInput,
			refineryRepo *repository.Repository[models.RefineryReport],
		) ([]models.RefineryReport, *repository.Paginator, error)
	}

	ActivityServiceInterface interface {
		ListActivities(
			ctx context.Context,
			input ListActivityInput,
			activityRepo *repository.Repository[models.Activity],
		) ([]models.Activity, *repository.Paginator, error)
	}
)
