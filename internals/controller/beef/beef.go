package beef

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	apiV1 "github.com/mamochiro/beef/pkg/api/v1"
)

func (c *Controller) Beef(ctx context.Context, empty *empty.Empty) (*apiV1.BeefResponse, error) {
	beef, err := c.service.BeefSummary(ctx)

	if err != nil {
		return nil, err
	}

	return &apiV1.BeefResponse{
		Beef: beef,
	}, nil
}
