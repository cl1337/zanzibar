// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package barEndpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarArgNotStructHandler is the handler for "/bar/arg-not-struct-path"
type BarArgNotStructHandler struct {
	Clients *module.ClientDependencies
}

// NewBarArgNotStructHandler creates a handler
func NewBarArgNotStructHandler(
	gateway *zanzibar.Gateway,
	deps *module.Dependencies,
) *BarArgNotStructHandler {
	return &BarArgNotStructHandler{
		Clients: deps.Client,
	}
}

// Register adds the http handler to the gateway's http router
func (handler *BarArgNotStructHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"POST", "/bar/arg-not-struct-path",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"argNotStruct",
			handler.HandleRequest,
		),
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/bar/arg-not-struct-path".
func (handler *BarArgNotStructHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsBarBar.Bar_ArgNotStruct_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := ArgNotStructEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBarBar.BarException:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			req.Logger.Warn("Workflow for endpoint returned error",
				zap.String("error", errValue.Error()),
			)
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}

	res.WriteJSONBytes(200, cliRespHeaders, nil)
}

// ArgNotStructEndpoint calls thrift client Bar.ArgNotStruct
type ArgNotStructEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w ArgNotStructEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_ArgNotStruct_Args,
) (zanzibar.Header, error) {
	clientRequest := convertToArgNotStructClientRequest(r)

	clientHeaders := map[string]string{}

	_, err := w.Clients.Bar.ArgNotStruct(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBarBar.BarException:
			serverErr := convertArgNotStructBarException(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.String("error", errValue.Error()),
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	return resHeaders, nil
}

func convertToArgNotStructClientRequest(in *endpointsBarBar.Bar_ArgNotStruct_Args) *clientsBarBar.Bar_ArgNotStruct_Args {
	out := &clientsBarBar.Bar_ArgNotStruct_Args{}

	out.Request = string(in.Request)

	return out
}

func convertArgNotStructBarException(
	clientError *clientsBarBar.BarException,
) *endpointsBarBar.BarException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBarBar.BarException{}
	return serverError
}
