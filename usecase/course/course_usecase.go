package course

import (
	"context"
	"database/sql"
	"time"

	"github.com/fauzanmh/olp-admin/entity"
	appInit "github.com/fauzanmh/olp-admin/init"
	mysqlRepo "github.com/fauzanmh/olp-admin/repository/mysql"
	"github.com/fauzanmh/olp-admin/schema/course"
	"go.uber.org/zap"
)

type usecase struct {
	config    *appInit.Config
	mysqlRepo mysqlRepo.Repository
}

func NewCourseUseCase(config *appInit.Config, mysqlRepo mysqlRepo.Repository) Usecase {
	return &usecase{
		config:    config,
		mysqlRepo: mysqlRepo,
	}
}

func (u *usecase) Create(ctx context.Context, req *course.CourseCreateRequest) (err error) {
	// init arguments
	createCourseParams := &entity.CreateCourseParams{
		CourseCategoryID: req.CourseCategoryID,
		Name:             req.Name,
		Description:      req.Description,
		Price:            req.Price,
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}

	err = u.mysqlRepo.CreateCourse(ctx, createCourseParams)
	zap.S().Error(err)
	if err != nil {
		return
	}

	return
}
