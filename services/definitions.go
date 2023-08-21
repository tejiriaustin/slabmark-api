package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/utils"
)

type AccountsServiceInterface interface {
	SignInUser(ctx context.Context,
		input AddAccountInput,
		accountsRepo *repository.Repository[models.Account],
	) (*models.Account, error)

	AddAccount(ctx context.Context,
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
}

type (
	StoreServiceInterface interface {
		AddItem(ctx context.Context, input AddItemInput) (*models.StoreItem, error)
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
	}
)
