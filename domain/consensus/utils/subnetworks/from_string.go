package subnetworks

import (
	"encoding/hex"

	"github.com/e4p1k0/pyipad/domain/consensus/model/externalapi"
)

// FromString creates a DomainSubnetworkID from the given byte slice
func FromString(str string) (*externalapi.DomainSubnetworkID, error) {
	subnetworkIDBytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return FromBytes(subnetworkIDBytes)
}
