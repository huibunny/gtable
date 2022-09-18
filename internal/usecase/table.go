package usecase

import (
	"context"

	"gtable/config"
	"gtable/internal/entity"
	"gtable/internal/usecase/repo"
)

// TableUserCase -.
type TableUserCase struct {
	repo *repo.TableRepo
}

// New -.
func New(r *repo.TableRepo) *TableUserCase {
	return &TableUserCase{
		repo: r,
	}
}

// Login -.
func (uc *TableUserCase) Load(ctx context.Context, tables []config.Table) (entity.TableValue, int, error) {
	return uc.repo.Load(ctx, tables)
}
