package course

type CourseCreateRequest struct {
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description" validate:"required"`
	Price            string `json:"price" validate:"required"`
	CourseCategoryID int32  `json:"course_category_id" validate:"required"`
}

type CourseUpdateRequest struct {
	ID               int64  `param:"id" json:"-" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description" validate:"required"`
	Price            string `json:"price" validate:"required"`
	CourseCategoryID int32  `json:"course_category_id" validate:"required"`
}
