package member

import (
	"context"

	appInit "github.com/fauzanmh/olp-admin/init"
	userAdapter "github.com/fauzanmh/olp-admin/repository/adapter/user"
	mysqlRepo "github.com/fauzanmh/olp-admin/repository/mysql"
	"github.com/fauzanmh/olp-admin/schema/member"
)

type usecase struct {
	config      *appInit.Config
	mysqlRepo   mysqlRepo.Repository
	userAdapter userAdapter.UserAdapter
}

func NewMemberUseCase(config *appInit.Config, mysqlRepo mysqlRepo.Repository, userAdapter userAdapter.UserAdapter) Usecase {
	return &usecase{
		config:      config,
		mysqlRepo:   mysqlRepo,
		userAdapter: userAdapter,
	}
}

func (u *usecase) DeleteMember(ctx context.Context, req *member.DeleteMemberRequest) (err error) {
	// delete user to ms user
	err = u.userAdapter.DeleteMember(ctx, req.ID)
	if err != nil {
		return
	}

	return
}
