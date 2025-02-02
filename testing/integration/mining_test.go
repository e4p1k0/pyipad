package integration

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/e4p1k0/pyipad/app/appmessage"
	"github.com/e4p1k0/pyipad/domain/consensus/model/externalapi"
	"github.com/e4p1k0/pyipad/domain/consensus/utils/mining"
)

func mineNextBlock(t *testing.T, harness *appHarness) *externalapi.DomainBlock {
	if os.Getenv("SKIP_ADDRESSES_RELATED_TESTS") != "" {
		t.Skip()
	}

	blockTemplate, err := harness.rpcClient.GetBlockTemplate(harness.miningAddress, "integration")
	if err != nil {
		t.Fatalf("Error getting block template: %+v", err)
	}

	block, err := appmessage.RPCBlockToDomainBlock(blockTemplate.Block)
	if err != nil {
		t.Fatalf("Error converting block: %s", err)
	}

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	mining.SolveBlock(block, rd)

	_, err = harness.rpcClient.SubmitBlockAlsoIfNonDAA(block)
	if err != nil {
		t.Fatalf("Error submitting block: %s", err)
	}

	return block
}
