package reporter

import (
	"github.com/openzipkin/zipkin-go/model"
	"github.com/openzipkin/zipkin-go/sender"
)

//AsyncReporter report spans asynchronous to a defined Sender and handles batching
type AsyncReporter interface {
	Flush()
}

type boundedAsyncReporter struct {
	sender sender.Sender
}

func (r *boundedAsyncReporter) Send(s model.SpanModel) {

}

func (r *boundedAsyncReporter) Report(s model.SpanModel) {

}

func (r *boundedAsyncReporter) Flush() {

}

func (r *boundedAsyncReporter) Close() error {
	return nil
}

// NewReporter creates a new async reporter
func NewReporter(s sender.Sender) Reporter {
	return &boundedAsyncReporter{
		sender: s,
	}
}
