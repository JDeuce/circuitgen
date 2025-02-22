// Code generated by circuitgen tool. DO NOT EDIT

package circuittestmulti

import (
	"context"

	"github.com/cep21/circuit"
	"github.com/twitchtv/circuitgen/internal/circuitgentest"
	"github.com/twitchtv/circuitgen/internal/circuitgentest/model"
	"github.com/twitchtv/circuitgen/internal/circuitgentest/rep"
)

// CircuitWrapperPublisherConfig contains configuration for CircuitWrapperPublisher. All fields are optional
type CircuitWrapperPublisherConfig struct {
	// ShouldSkipError determines whether an error should be skipped and have the circuit
	// track the call as successful. This takes precedence over IsBadRequest
	ShouldSkipError func(error) bool

	// IsBadRequest is an optional bad request checker. It is useful to not count user errors as faults
	IsBadRequest func(error) bool

	// Prefix is prepended to all circuit names
	Prefix string

	// Defaults are used for all created circuits. Per-circuit configs override this
	Defaults circuit.Config

	// CircuitPublish is the configuration used for the Publish circuit. This overrides values set by Defaults
	CircuitPublish circuit.Config
	// CircuitPublishWithResult is the configuration used for the PublishWithResult circuit. This overrides values set by Defaults
	CircuitPublishWithResult circuit.Config
}

// CircuitWrapperPublisher is a circuit wrapper for circuitgentest.Publisher
type CircuitWrapperPublisher struct {
	circuitgentest.Publisher

	// ShouldSkipError determines whether an error should be skipped and have the circuit
	// track the call as successful. This takes precedence over IsBadRequest
	ShouldSkipError func(error) bool

	// IsBadRequest checks whether to count a user error against the circuit. It is recommended to set this
	IsBadRequest func(error) bool

	// CircuitPublish is the circuit for method Publish
	CircuitPublish *circuit.Circuit
	// CircuitPublishWithResult is the circuit for method PublishWithResult
	CircuitPublishWithResult *circuit.Circuit
}

// NewCircuitWrapperPublisher creates a new circuit wrapper and initializes circuits
func NewCircuitWrapperPublisher(
	manager *circuit.Manager,
	embedded circuitgentest.Publisher,
	conf CircuitWrapperPublisherConfig,
) (*CircuitWrapperPublisher, error) {
	if conf.ShouldSkipError == nil {
		conf.ShouldSkipError = func(err error) bool {
			return false
		}
	}

	if conf.IsBadRequest == nil {
		conf.IsBadRequest = func(err error) bool {
			return false
		}
	}

	w := &CircuitWrapperPublisher{
		Publisher:       embedded,
		ShouldSkipError: conf.ShouldSkipError,
		IsBadRequest:    conf.IsBadRequest,
	}

	var err error

	w.CircuitPublish, err = manager.CreateCircuit(conf.Prefix+"Publisher.Publish", conf.CircuitPublish, conf.Defaults)
	if err != nil {
		return nil, err
	}

	w.CircuitPublishWithResult, err = manager.CreateCircuit(conf.Prefix+"Publisher.PublishWithResult", conf.CircuitPublishWithResult, conf.Defaults)
	if err != nil {
		return nil, err
	}

	return w, nil
}

// Publish calls the embedded circuitgentest.Publisher's method Publish with CircuitPublish
func (w *CircuitWrapperPublisher) Publish(ctx context.Context, p1 map[circuitgentest.Seed][][]circuitgentest.Grant, p2 circuitgentest.TopicsList, p3 ...rep.PublishOption) (map[string]struct{}, error) {
	var r0 map[string]struct{}
	var skippedErr error

	err := w.CircuitPublish.Run(ctx, func(ctx context.Context) error {
		var err error
		r0, err = w.Publisher.Publish(ctx, p1, p2, p3...)

		if w.ShouldSkipError(err) {
			skippedErr = err
			return nil
		}

		if w.IsBadRequest(err) {
			return &circuit.SimpleBadRequest{Err: err}
		}
		return err
	})

	if skippedErr != nil {
		err = skippedErr
	}

	if berr, ok := err.(*circuit.SimpleBadRequest); ok {
		err = berr.Err
	}

	return r0, err
}

// PublishWithResult calls the embedded circuitgentest.Publisher's method PublishWithResult with CircuitPublishWithResult
func (w *CircuitWrapperPublisher) PublishWithResult(ctx context.Context, p1 rep.PublishInput) (*model.Result, error) {
	var r0 *model.Result
	var skippedErr error

	err := w.CircuitPublishWithResult.Run(ctx, func(ctx context.Context) error {
		var err error
		r0, err = w.Publisher.PublishWithResult(ctx, p1)

		if w.ShouldSkipError(err) {
			skippedErr = err
			return nil
		}

		if w.IsBadRequest(err) {
			return &circuit.SimpleBadRequest{Err: err}
		}
		return err
	})

	if skippedErr != nil {
		err = skippedErr
	}

	if berr, ok := err.(*circuit.SimpleBadRequest); ok {
		err = berr.Err
	}

	return r0, err
}

var _ circuitgentest.Publisher = (*CircuitWrapperPublisher)(nil)
