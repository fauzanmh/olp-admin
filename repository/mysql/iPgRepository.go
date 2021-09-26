package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-admin/entity"
)

type Repository interface {
	// Courses
	CreateCourse(ctx context.Context, args *entity.CreateCourseParams) (err error)
	GetAllCourses(ctx context.Context) ([]entity.GetAllCoursesRow, error)
	UpdateCourse(ctx context.Context, arg *entity.UpdateCourseParams) error

	//Tx
	BeginTx(ctx context.Context) (*sql.Tx, error)
	WithTx(tx *sql.Tx) *Queries
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}
