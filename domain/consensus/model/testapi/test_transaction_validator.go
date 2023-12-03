package testapi

import (
	"github.com/e4p1k0/pyipad/domain/consensus/model"
	"github.com/e4p1k0/pyipad/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
