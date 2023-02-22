package rest

import "context"

//go:generate mockery --name=Interface
type Interface interface {
	BeefSummary(ctx context.Context) (*[]string, error)
}
