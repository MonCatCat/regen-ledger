package server

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/types/ormstore"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	coretypes "github.com/regen-network/regen-ledger/x/ecocredit/core"
	marketplacetypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/marketplace"
)

const (
	ProjectInfoTablePrefix    byte = 0x10
	ProjectInfoTableSeqPrefix byte = 0x11
	ProjectsByClassIDIndex    byte = 0x12
	BatchesByProjectIndex     byte = 0x13

	// sell order table
	SellOrderTablePrefix             byte = 0x20
	SellOrderTableSeqPrefix          byte = 0x21
	SellOrderByExpirationIndexPrefix byte = 0x22
	SellOrderByAddressIndexPrefix    byte = 0x23
	SellOrderByBatchDenomIndexPrefix byte = 0x24

	// buy order table
	BuyOrderTablePrefix             byte = 0x25
	BuyOrderTableSeqPrefix          byte = 0x26
	BuyOrderByExpirationIndexPrefix byte = 0x27
	BuyOrderByAddressIndexPrefix    byte = 0x28

	AskDenomTablePrefix byte = 0x30
)

type serverImpl struct {
	storeKey sdk.StoreKey

	paramSpace    paramtypes.Subspace
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper

	basketKeeper      basket.Keeper
	coreKeeper        core.Keeper
	marketplaceKeeper marketplace.Keeper

	db         ormdb.ModuleDB
	stateStore api.StateStore
}

func newServer(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper, distKeeper ecocredit.DistributionKeeper) serverImpl {
	s := serverImpl{
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}

	var err error
	s.db, err = ormstore.NewStoreKeyDB(&ecocredit.ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	coreStore, basketStore, marketStore := getStateStores(s.db)
	s.basketKeeper = basket.NewKeeper(basketStore, coreStore, bankKeeper, distKeeper, s.paramSpace)
	s.coreKeeper = core.NewKeeper(coreStore, bankKeeper, s.paramSpace)
	s.marketplaceKeeper = marketplace.NewKeeper(marketStore, coreStore, bankKeeper, s.paramSpace)

	return s
}

func getStateStores(db ormdb.ModuleDB) (api.StateStore, basketapi.StateStore, marketApi.StateStore) {
	coreStore, err := api.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	basketStore, err := basketapi.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	marketStore, err := marketApi.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	return coreStore, basketStore, marketStore
}

func RegisterServices(
	configurator server.Configurator,
	paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper,
	distKeeper ecocredit.DistributionKeeper,
) ecocredit.Keeper {
	impl := newServer(configurator.ModuleKey(), paramSpace, accountKeeper, bankKeeper, distKeeper)

	coretypes.RegisterMsgServer(configurator.MsgServer(), impl.coreKeeper)
	coretypes.RegisterQueryServer(configurator.QueryServer(), impl.coreKeeper)

	baskettypes.RegisterMsgServer(configurator.MsgServer(), impl.basketKeeper)
	baskettypes.RegisterQueryServer(configurator.QueryServer(), impl.basketKeeper)

	marketplacetypes.RegisterMsgServer(configurator.MsgServer(), impl.marketplaceKeeper)
	marketplacetypes.RegisterQueryServer(configurator.QueryServer(), impl.marketplaceKeeper)

	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	configurator.RegisterMigrationHandler(impl.RunMigrations)

	// TODO: uncomment when sims are refactored https://github.com/regen-network/regen-ledger/issues/920
	// configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
	configurator.RegisterInvariantsHandler(impl.RegisterInvariants)
	return impl
}
