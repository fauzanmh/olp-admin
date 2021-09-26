package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-admin/constant"
	"github.com/fauzanmh/olp-admin/entity"
)

const getOneCourseCategory = `-- name: GetOneCourseCategory :one
SELECT id, name, total_used FROM course_categories
WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetOneCourseCategory(ctx context.Context, id int32) (entity.GetOneCourseCategoryRow, error) {
	row := q.queryRow(ctx, q.getOneCourseCategoryStmt, getOneCourseCategory, id)
	var i entity.GetOneCourseCategoryRow
	err := row.Scan(&i.ID, &i.Name, &i.TotalUsed)
	if err == sql.ErrNoRows {
		err = constant.ErrorMessageCourseCategoryNotFound
	}
	return i, err
}

const updateTotalUsed = `-- name: UpdateTotalUsed :exec
UPDATE course_categories SET total_used = total_used+1
WHERE id = ?
`

func (q *Queries) UpdateTotalUsed(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.updateTotalUsedStmt, updateTotalUsed, id)
	return err
}
