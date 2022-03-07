package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// TODO: Revisit this once we have proper gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasCostPerIteration = uint64(10)

type Keeper struct {
	stateStore ecocreditv1beta1.StateStore
	bankKeeper ecocredit.BankKeeper
	params     ParamKeeper
	ak         ecocredit.AccountKeeper
}
type ParamKeeper interface {
	GetParamSet(ctx sdk.Context, ps types.ParamSet)
}

func NewKeeper(ss ecocreditv1beta1.StateStore, bk ecocredit.BankKeeper, params ParamKeeper) Keeper {
	return Keeper{
		stateStore: ss,
		bankKeeper: bk,
		params:     params,
	}
}

// TODO: uncomment when impl
// var _ v1beta1.MsgServer = &Keeper{}
