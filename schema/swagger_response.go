package schema

import "github.com/fauzanmh/olp-admin/schema/course"

type SwaggerGetAllCoursesResponse struct {
	Base
	course.GetAllCoursesResponse `json:"data"`
}
