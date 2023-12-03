package rpchandlers

import (
	"github.com/e4p1k0/pyipad/app/appmessage"
	"github.com/e4p1k0/pyipad/app/rpc/rpccontext"
	"github.com/e4p1k0/pyipad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
