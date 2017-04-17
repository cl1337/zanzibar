// Code generated by zanzibar
// @generated

package bar

import (
	"context"
	"net/http"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"
)

// HandleMissingArgRequest handles "/bar/missing-arg-path".
func HandleMissingArgRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {

	workflow := MissingArgEndpoint{
		Clients: clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, respHeaders, err := workflow.Handle(ctx, req.Header)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}

	res.WriteJSON(200, respHeaders, response)
}

// MissingArgEndpoint calls thrift client Bar.MissingArg
type MissingArgEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w MissingArgEndpoint) Handle(
	ctx context.Context,
	// TODO(sindelar): Switch to zanzibar.Headers when tchannel
	// generation is implemented.
	headers http.Header,
) (*endpointsBarBar.BarResponse, map[string]string, error) {

	clientHeaders := map[string]string{}
	for k, v := range map[string]string{} {
		clientHeaders[v] = headers.Get(k)
	}

	clientRespBody, respHeaders, err := w.Clients.Bar.MissingArg(
		ctx, clientHeaders,
	)
	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		return nil, nil, err
	}

	endRespHead := map[string]string{}
	for k, v := range map[string]string{} {
		endRespHead[v] = respHeaders[k]
	}

	response := convertMissingArgClientResponse(clientRespBody)
	return response, endRespHead, nil
}

func convertMissingArgClientResponse(body *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := &endpointsBarBar.BarResponse{}
	return downstreamResponse
}
