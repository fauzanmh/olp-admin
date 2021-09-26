package entity

import (
	"database/sql"
)

type Course struct {
	ID               int64         `json:"id"`
	CourseCategoryID int32         `json:"course_category_id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Price            string        `json:"price"`
	CreatedAt        int64         `json:"created_at"`
	UpdatedAt        sql.NullInt64 `json:"updated_at"`
	DeletedAt        sql.NullInt64 `json:"deleted_at"`
}

type CreateCourseParams struct {
	CourseCategoryID int32         `json:"course_category_id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Price            string        `json:"price"`
	CreatedAt        int64         `json:"created_at"`
	UpdatedAt        sql.NullInt64 `json:"updated_at"`
}
