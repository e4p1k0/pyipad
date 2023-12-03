package model

import "github.com/e4p1k0/pyipad/domain/consensus/model/externalapi"

// AcceptanceDataStore represents a store of AcceptanceData
type AcceptanceDataStore interface {
	Store
	Stage(stagingArea *StagingArea, blockHash *externalapi.DomainHash, acceptanceData externalapi.AcceptanceData)
	IsStaged(stagingArea *StagingArea) bool
	Get(dbContext DBReader, stagingArea *StagingArea, blockHash *externalapi.DomainHash) (externalapi.AcceptanceData, error)
	Delete(stagingArea *StagingArea, blockHash *externalapi.DomainHash)
}
