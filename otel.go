package asynq

import (
	"github.com/hibiken/asynq/internal/base"
	"go.opentelemetry.io/otel/propagation"
)

const tracerName = "github.com/hibiken/asynq"

type metadataSupplier struct {
	taskMessage *base.TaskMessage
}

// assert that metadataSupplier implements the TextMapCarrier interface.
var _ propagation.TextMapCarrier = &metadataSupplier{}

func (s *metadataSupplier) Get(key string) string {
	value, ok := s.taskMessage.Headers[key]
	if !ok {
		return ""
	}
	return value
}

func (s *metadataSupplier) Set(key string, value string) {
	if s.taskMessage.Headers == nil {
		s.taskMessage.Headers = make(map[string]string)
	}
	s.taskMessage.Headers[key] = value
}

func (s *metadataSupplier) Keys() []string {
	out := make([]string, 0, len(s.taskMessage.Headers))
	for key := range s.taskMessage.Headers {
		out = append(out, key)
	}
	return out
}
