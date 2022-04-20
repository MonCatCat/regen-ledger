//go:build norace
// +build norace

package testsuite_test

import (
	"testing"

	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/MonCatCat/regen-ledger/v3/app/testsuite"
	"github.com/stretchr/testify/suite"
)

func TestNetwork(t *testing.T) {
	cfg := testsuite.DefaultConfig()

	suite.Run(t, network.NewIntegrationTestSuite(cfg))
}
