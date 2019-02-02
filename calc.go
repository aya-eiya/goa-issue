package calc

import (
	"context"
	"log"

	calcsvc "github.com/aya-eiya/goa-issue/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calcsvc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calcsvc.AddPayload) (res int, err error) {
	return p.A + p.B, nil
}
