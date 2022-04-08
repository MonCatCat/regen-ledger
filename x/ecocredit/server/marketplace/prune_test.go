package marketplace

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestSell_Prune(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	blockTime, err := time.Parse("2006-01-02", "2020-01-01")
	assert.NilError(t, err)
	expired, err := time.Parse("2006-01-02", "2019-12-30")
	assert.NilError(t, err)
	notExpired, err := time.Parse("2006-01-02", "2022-01-01")
	assert.NilError(t, err)

	sellOrderIds := s.insertSellOrder(&marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &expired},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &notExpired},
		},
	})

	// setup block time so the orders expire
	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// get the balance before pruning
	balBefore, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supBefore, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	// prune the orders
	err = s.k.PruneSellOrders(s.ctx)
	assert.NilError(t, err)

	balAfter, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supAfter, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	// we can reuse this function and pass the negated amount to get our desired behavior.
	assertCreditsEscrowed(t, balBefore, balAfter, supBefore, supAfter, math.NewDecFromInt64(-10))

	assert.Equal(t, 2, len(sellOrderIds))
	shouldBeExpired := sellOrderIds[0]
	shouldBeValid := sellOrderIds[1]

	_, err = s.marketStore.SellOrderTable().Get(s.ctx, shouldBeExpired)
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	_, err = s.marketStore.SellOrderTable().Get(s.ctx, shouldBeValid)
	assert.NilError(t, err)
}
