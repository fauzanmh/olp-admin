package http

import (
	"github.com/fauzanmh/olp-admin/pkg/util"
	"github.com/fauzanmh/olp-admin/schema/member"
	usecase "github.com/fauzanmh/olp-admin/usecase/member"
	"github.com/labstack/echo/v4"
)

type MemberHandler struct {
	usecase usecase.Usecase
}

func NewMemberHandler(e *echo.Group, uc usecase.Usecase) {
	handler := &MemberHandler{
		usecase: uc,
	}

	routerV1 := e.Group("/v1")
	routerV1.DELETE("/member/:id", handler.Delete)
}

// Delete godoc
// @Summary Delete Member
// @Description Delete Member
// @Tags Member
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} schema.Base
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/member/{id} [delete]
func (h *MemberHandler) Delete(c echo.Context) error {
	req := member.DeleteMemberRequest{}
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

	err = h.usecase.DeleteMember(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success delete member", nil)
}
