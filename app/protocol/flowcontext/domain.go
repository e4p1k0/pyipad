package flowcontext

import (
	"github.com/e4p1k0/pyipad/domain"
)

// Domain returns the Domain object associated to the flow context.
func (f *FlowContext) Domain() domain.Domain {
	return f.domain
}
