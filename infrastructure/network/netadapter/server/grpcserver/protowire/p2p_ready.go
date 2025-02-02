package protowire

import (
	"github.com/e4p1k0/pyipad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PyipadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *PyipadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
