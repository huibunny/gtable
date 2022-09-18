// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"gtable/config"
	"gtable/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Load -.
	Table interface {
		Load(context.Context, []config.Table) (entity.TableValue, int, error)
	}
)
