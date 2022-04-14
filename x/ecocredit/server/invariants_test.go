package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	coreserver "github.com/regen-network/regen-ledger/x/ecocredit/server/core"
)

type baseSuite struct {
	t            *testing.T
	db           ormdb.ModuleDB
	stateStore   api.StateStore
	ctx          context.Context
	k            coreserver.Keeper
	ctrl         *gomock.Controller
	bankKeeper   *mocks.MockBankKeeper
	paramsKeeper *mocks.MockParamKeeper
	storeKey     *sdk.KVStoreKey
	sdkCtx       sdk.Context
}

func setupBase(t *testing.T) *baseSuite {
	// prepare database
	s := &baseSuite{t: t}
	var err error
	s.db, err = ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	s.stateStore, err = api.NewStateStore(s.db)
	assert.NilError(t, err)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	s.storeKey = sdk.NewKVStoreKey("test")
	cms.MountStoreWithDB(s.storeKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// setup test keeper
	s.ctrl = gomock.NewController(t)
	assert.NilError(t, err)
	s.bankKeeper = mocks.NewMockBankKeeper(s.ctrl)
	s.paramsKeeper = mocks.NewMockParamKeeper(s.ctrl)
	s.k = coreserver.NewKeeper(s.stateStore, s.bankKeeper, s.paramsKeeper)

	return s
}

func TestBatchSupplyInvariant(t *testing.T) {
	acc1 := sdk.AccAddress([]byte("account1"))
	acc2 := sdk.AccAddress([]byte("account2"))

	testCases := []struct {
		msg           string
		balances      []*core.BatchBalance
		supply        []*core.BatchSupply
		basketBalance map[uint64]math.Dec
		expBroken     bool
	}{
		{
			"valid test case",
			[]*core.BatchBalance{
				{
					Address:  acc1,
					BatchId:  1,
					Tradable: "210",
					Retired:  "110",
				},
			},
			[]*core.BatchSupply{
				{
					BatchId:        1,
					TradableAmount: "220",
					RetiredAmount:  "110",
				},
			},
			map[uint64]math.Dec{1: math.NewDecFromInt64(10)},
			false,
		},
		{
			"valid test case multiple denom",
			[]*core.BatchBalance{
				{
					Address:  acc1,
					BatchId:  1,
					Tradable: "310.579",
					Retired:  "0",
				},
				{
					Address:  acc2,
					BatchId:  2,
					Tradable: "210.456",
					Retired:  "100.1234",
				},
			},
			[]*core.BatchSupply{
				{
					BatchId:        1,
					TradableAmount: "320.579",
					RetiredAmount:  "0",
				},
				{
					BatchId:        2,
					TradableAmount: "220.456",
					RetiredAmount:  "100.1234",
				},
			},
			map[uint64]math.Dec{1: math.NewDecFromInt64(10), 2: math.NewDecFromInt64(10)},
			false,
		},
		{
			"fail with error tradable balance not found",
			[]*core.BatchBalance{
				{
					Address:  acc1,
					BatchId:  1,
					Tradable: "100.123",
				},
				{
					Address:  acc2,
					BatchId:  1,
					Tradable: "210.456",
				},
			},
			[]*core.BatchSupply{
				{
					BatchId:        1,
					TradableAmount: "310.579",
					RetiredAmount:  "0",
				},
				{
					BatchId:        3,
					TradableAmount: "1234",
					RetiredAmount:  "0",
				},
			},
			map[uint64]math.Dec{},
			true,
		},
		{
			"fail with error supply does not match",
			[]*core.BatchBalance{
				{
					Address:  acc1,
					BatchId:  1,
					Tradable: "310.579",
				},
				{
					BatchId:  2,
					Address:  acc2,
					Tradable: "1234",
				},
			},
			[]*core.BatchSupply{
				{
					BatchId:        1,
					TradableAmount: "310.579",
					RetiredAmount:  "123",
				},
				{
					BatchId:        2,
					TradableAmount: "12345",
					RetiredAmount:  "0",
				},
			},
			map[uint64]math.Dec{},
			true,
		},
		{
			"valid clase escrowed balance",
			[]*core.BatchBalance{
				{
					Address:  acc1,
					BatchId:  1,
					Tradable: "100",
					Escrowed: "10",
					Retired:  "1",
				},
				{
					BatchId:  2,
					Address:  acc2,
					Tradable: "1234",
					Retired:  "123",
					Escrowed: "766",
				},
			},
			[]*core.BatchSupply{
				{
					BatchId:        1,
					TradableAmount: "110",
					RetiredAmount:  "10",
				},
				{
					BatchId:        2,
					TradableAmount: "2000",
					RetiredAmount:  "123",
				},
			},
			map[uint64]math.Dec{},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite := setupBase(t)
		t.Run(tc.msg, func(t *testing.T) {
			initBalances(t, suite.ctx, suite.stateStore, tc.balances)
			initSupply(t, suite.ctx, suite.stateStore, tc.supply)

			msg, broken := coreserver.BatchSupplyInvariant(suite.ctx, suite.k, tc.basketBalance)
			if tc.expBroken {
				require.True(t, broken, msg)
			} else {
				require.False(t, broken, msg)
			}
		})
	}

}

func initBalances(t *testing.T, ctx context.Context, ss api.StateStore, balances []*core.BatchBalance) {
	for _, b := range balances {
		_, err := math.NewNonNegativeDecFromString(b.Tradable)
		require.NoError(t, err)

		require.NoError(t, ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
			Address:  b.Address,
			BatchId:  b.BatchId,
			Tradable: b.Tradable,
			Retired:  b.Retired,
			Escrowed: b.Escrowed,
		}))
	}
}

func initSupply(t *testing.T, ctx context.Context, ss api.StateStore, supply []*core.BatchSupply) {
	for _, s := range supply {
		err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
			BatchId:         s.BatchId,
			TradableAmount:  s.TradableAmount,
			RetiredAmount:   s.RetiredAmount,
			CancelledAmount: s.CancelledAmount,
		})
		require.NoError(t, err)
	}
}
