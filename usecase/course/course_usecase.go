package course

import (
	"context"
	"database/sql"
	"time"

	"github.com/fauzanmh/olp-admin/entity"
	appInit "github.com/fauzanmh/olp-admin/init"
	mysqlRepo "github.com/fauzanmh/olp-admin/repository/mysql"
	"github.com/fauzanmh/olp-admin/schema/course"
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

// --- get all course --- ///
func (u *usecase) Get(ctx context.Context) (res []course.GetAllCoursesResponse, err error) {
	// get from database
	data, err := u.mysqlRepo.GetAllCourses(ctx)
	if err != nil {
		return
	}

	// convert from entity to schema
	for idx := range data {
		res = append(res, course.GetAllCoursesResponse{
			ID:               data[idx].ID,
			CourseCategoryID: data[idx].CourseCategoryID,
			Name:             data[idx].Name,
			Description:      data[idx].Description,
			Price:            data[idx].Price,
		})
	}

	return
}

// --- create course --- ///
func (u *usecase) Create(ctx context.Context, req *course.CourseCreateRequest) (err error) {
	// arguments
	createCourseParams := &entity.CreateCourseParams{
		CourseCategoryID: req.CourseCategoryID,
		Name:             req.Name,
		Description:      req.Description,
		Price:            req.Price,
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}

	// store to database
	err = u.mysqlRepo.CreateCourse(ctx, createCourseParams)
	if err != nil {
		return
	}

	return
}

// --- update course --- ///
func (u *usecase) Update(ctx context.Context, req *course.CourseUpdateRequest) (err error) {
	// arguments
	updateCourseParams := &entity.UpdateCourseParams{
		ID:               req.ID,
		CourseCategoryID: req.CourseCategoryID,
		Name:             req.Name,
		Description:      req.Description,
		Price:            req.Price,
		UpdatedAt:        sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}

	err = u.mysqlRepo.UpdateCourse(ctx, updateCourseParams)
	if err != nil {
		return
	}

	return
}
