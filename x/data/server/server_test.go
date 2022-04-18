package server

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/data/mocks"
)

type baseSuite struct {
	t      gocuke.TestingT
	ctx    context.Context
	sdkCtx sdk.Context
	server serverImpl
	addrs  []sdk.AccAddress
}

func setupBase(t gocuke.TestingT) *baseSuite {
	s := &baseSuite{t: t}

	// set up store
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	sk := sdk.NewKVStoreKey("test")
	cms.MountStoreWithDB(sk, sdk.StoreTypeIAVL, db)
	require.NoError(t, cms.LoadLatestVersion())

	// set up context
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// set up server
	ctrl := gomock.NewController(t)
	ak := mocks.NewMockAccountKeeper(ctrl)
	bk := mocks.NewMockBankKeeper(ctrl)
	s.server = newServer(sk, ak, bk)

	// set up addresses
	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()
	s.addrs = append(s.addrs, addr1)
	s.addrs = append(s.addrs, addr2)

	return s
}
