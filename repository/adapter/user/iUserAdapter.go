package user

import (
	"context"
)

type UserAdapter interface {
	DeleteMember(ctx context.Context, id int64) (err error)
}
