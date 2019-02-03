package calc

import (
	"context"
	"fmt"
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
func (s *calcsrvc) Add(ctx context.Context, p *calcsvc.AddPayload) (*calcsvc.Result, error) {
	str := fmt.Sprintf("%d + %d = %d", p.A, p.B, p.A+p.B)
	res := &calcsvc.Result{Result: &str}
	return res, nil
}
