package dal

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProjectDal struct {
	DB *gorm.DB
}

func NewProjectDal(db *gorm.DB) *ProjectDal {
	return &ProjectDal{
		DB: db,
	}
}

func (dal *ProjectDal) Create(ctx context.Context, s string) error {
	result := dal.DB.Create(s)
	return errors.WithStack(result.Error)
}
