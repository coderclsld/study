package handle

import (
	"context"

	"clsld.com/test-wire/pkg/service"
	"google.golang.org/grpc/profiling/service"
)

type ProjectHandle struct {
	ProjectService *service.ProjectService
}

func NewProjectHandle(srv *service.ProjectService) *ProjectHandle {
	return &ProjectHandle{
		ProjectService: srv,
	}
}

func (s *ProjectHandle) CreateProject(ctx context.Context) {

}
