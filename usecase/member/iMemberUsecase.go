package member

import (
	"context"

	"github.com/fauzanmh/olp-admin/schema/member"
)

type Usecase interface {
	DeleteMember(ctx context.Context, req *member.DeleteMemberRequest) (err error)
}
