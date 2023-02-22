package beef

import "context"

//go:generate mockery --name=Interface
type Interface interface {
	BeefSummary(ctx context.Context) (map[string]int32, error)
}
