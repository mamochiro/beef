package wrapper

import (
	"context"
)

func (wrp *Wrapper) BeefSummary(ctx context.Context) (map[string]int32, error) {
	res, err := wrp.Service.BeefSummary(ctx)
	return res, err
}
