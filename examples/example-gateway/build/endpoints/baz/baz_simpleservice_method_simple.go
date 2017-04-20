// Code generated by zanzibar
// @generated

package baz

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
)

// SimpleHandler is the handler for "/baz/simple-path"
type SimpleHandler struct {
	Clients *clients.Clients
}

// NewSimpleEndpoint creates a handler
func NewSimpleEndpoint(
	gateway *zanzibar.Gateway,
) *SimpleHandler {
	return &SimpleHandler{
		Clients: gateway.Clients.(*clients.Clients),
	}
}

// HandleRequest handles "/baz/simple-path".
func (handler *SimpleHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {

	workflow := SimpleEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	cliRespHeaders, err := workflow.Handle(ctx, req.Header)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}

	res.WriteJSONBytes(204, cliRespHeaders, nil)
}

// SimpleEndpoint calls thrift client Baz.Simple
type SimpleEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w SimpleEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.ServerHeaderInterface,
) (zanzibar.ServerHeaderInterface, error) {

	clientHeaders := map[string]string{}

	_, err := w.Clients.Baz.Simple(ctx, clientHeaders)

	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		// TODO(sindelar): Consider returning partial headers in error case.
		return nil, err
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	return resHeaders, nil
}
