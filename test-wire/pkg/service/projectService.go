package service

import (
	"context"

	"clsld.com/test-wire/pkg/dal"
)

type ProjectService struct {
	ProjectDal *dal.ProjectDal
}

func NewProjectService(projectDal *dal.ProjectDal) *ProjectService {
	return &ProjectService{
		ProjectDal: projectDal,
	}
}
func (s *ProjectService) Create(ctx context.Context) (int64, error) {
	return 0, nil
}
