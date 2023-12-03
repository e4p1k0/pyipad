package rpchandlers

import (
	"github.com/e4p1k0/pyipad/app/appmessage"
	"github.com/e4p1k0/pyipad/app/rpc/rpccontext"
	"github.com/e4p1k0/pyipad/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
