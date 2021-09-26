package http

import (
	"github.com/fauzanmh/olp-admin/pkg/util"
	"github.com/fauzanmh/olp-admin/schema/course"
	usecaseCourse "github.com/fauzanmh/olp-admin/usecase/course"
	"github.com/labstack/echo/v4"
)

type CourseHandler struct {
	usecase usecaseCourse.Usecase
}

func NewCourseHandler(e *echo.Group, uc usecaseCourse.Usecase) {
	handler := &CourseHandler{
		usecase: uc,
	}

	routerV1 := e.Group("/v1")
	routerV1.GET("/course", handler.Get)
	routerV1.POST("/course", handler.Create)
	routerV1.PUT("/course/:id", handler.Update)
}

// Get godoc
// @Summary Get All Courses
// @Description Get All Courses
// @Tags Course
// @Accept json
// @Produce json
// @Success 200 {object} schema.SwaggerGetAllCoursesResponse
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/course [get]
func (h *CourseHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.usecase.Get(ctx)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success get courses", data)
}

// Create godoc
// @Summary Create Course
// @Description Create a new course
// @Tags Course
// @Accept json
// @Produce json
// @Param request body course.CourseCreateRequest{} true "Request Body"
// @Success 200 {object} schema.Base
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/course [post]
func (h *CourseHandler) Create(c echo.Context) error {
	req := course.CourseCreateRequest{}
	ctx := c.Request().Context()

	// parsing
	err := util.ParsingParameter(c, &req)
	if err != nil {
		return util.ErrorParsing(c, err, nil)
	}

	// validate
	err = util.ValidateParameter(c, &req)
	if err != nil {
		return util.ErrorValidate(c, err, nil)
	}

	err = h.usecase.Create(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success create course", nil)
}

// Update godoc
// @Summary Update Course
// @Description Update course
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int true "ID of Course"
// @Param request body course.CourseUpdateRequest{} true "Request Body"
// @Success 200 {object} schema.Base
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/course/{id} [put]
func (h *CourseHandler) Update(c echo.Context) error {
	req := course.CourseUpdateRequest{}
	ctx := c.Request().Context()

	// parsing
	err := util.ParsingParameter(c, &req)
	if err != nil {
		return util.ErrorParsing(c, err, nil)
	}

	// validate
	err = util.ValidateParameter(c, &req)
	if err != nil {
		return util.ErrorValidate(c, err, nil)
	}

	err = h.usecase.Update(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success update course", nil)
}
