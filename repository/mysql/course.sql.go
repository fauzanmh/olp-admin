package mysql

import (
	"context"

	"github.com/fauzanmh/olp-admin/entity"
)

const createCourse = `-- name: CreateCourse :exec
INSERT INTO courses (course_category_id, name, description, price, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
`

func (q *Queries) CreateCourse(ctx context.Context, arg *entity.CreateCourseParams) error {
	_, err := q.exec(ctx, q.createCourseStmt, createCourse,
		arg.CourseCategoryID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const updateCourse = `-- name: UpdateCourse :exec
UPDATE courses SET course_category_id = ?, name = ?, description = ?, price = ?, updated_at = ?
WHERE id = ?
`

func (q *Queries) UpdateCourse(ctx context.Context, arg *entity.UpdateCourseParams) error {
	_, err := q.exec(ctx, q.updateCourseStmt, updateCourse,
		arg.CourseCategoryID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
