package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v3 "github.com/regen-network/regen-ledger/x/ecocredit/migrations/v3"
)

func (s serverImpl) RunMigrations(ctx sdk.Context, cdc codec.Codec) error {
	return v3.MigrateState(ctx, s.storeKey, cdc, s.stateStore)
}
