package statistic

import (
	"context"

	"github.com/fauzanmh/olp-admin/schema/statistic"
)

type Usecase interface {
	Get(ctx context.Context) (res statistic.GetStatisticResponse, err error)
}
