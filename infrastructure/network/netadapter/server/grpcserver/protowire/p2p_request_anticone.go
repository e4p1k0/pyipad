package protowire

import (
	"github.com/e4p1k0/pyipad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PyipadMessage_RequestAnticone) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PyipadMessage_RequestAnticone is nil")
	}
	return x.RequestAnticone.toAppMessage()
}

func (x *RequestAnticoneMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RequestAnticoneMessage is nil")
	}
	blockHash, err := x.BlockHash.toDomain()
	if err != nil {
		return nil, err
	}

	contextHash, err := x.ContextHash.toDomain()
	if err != nil {
		return nil, err
	}

	return &appmessage.MsgRequestAnticone{
		BlockHash:   blockHash,
		ContextHash: contextHash,
	}, nil

}

func (x *PyipadMessage_RequestAnticone) fromAppMessage(msgRequestPastDiff *appmessage.MsgRequestAnticone) error {
	x.RequestAnticone = &RequestAnticoneMessage{
		BlockHash:   domainHashToProto(msgRequestPastDiff.BlockHash),
		ContextHash: domainHashToProto(msgRequestPastDiff.ContextHash),
	}
	return nil
}
