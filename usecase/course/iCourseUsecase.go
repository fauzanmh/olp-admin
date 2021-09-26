package course

import (
	"context"

	"github.com/fauzanmh/olp-admin/schema/course"
)

type Usecase interface {
	Create(ctx context.Context, req *course.CourseCreateRequest) (err error)
	Update(ctx context.Context, req *course.CourseUpdateRequest) (err error)
}
