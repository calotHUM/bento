package ratelimit_test

import (
	"testing"

	"github.com/Jeffail/benthos/v3/internal/log"
	"github.com/Jeffail/benthos/v3/internal/manager/mock"
	"github.com/Jeffail/benthos/v3/internal/old/metrics"
	"github.com/Jeffail/benthos/v3/internal/old/ratelimit"
)

func TestConstructorBadType(t *testing.T) {
	conf := ratelimit.NewConfig()
	conf.Type = "not_exist"

	if _, err := ratelimit.New(conf, mock.NewManager(), log.Noop(), metrics.Noop()); err == nil {
		t.Error("Expected error, received nil for invalid type")
	}
}