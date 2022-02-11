package basket

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Create(ctx context.Context, msg *basket.MsgCreate) (*basket.MsgCreateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	fee := k.ecocreditKeeper.GetCreateBasketFee(ctx)
	if err := msg.Validate(fee); err != nil {
		return nil, err
	}
	// TODO: need to decide about the denom creation
	denom := msg.Name
	sender, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}
	k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, sender, feeModuleAccName, fee)

	err = sdkCtx.EventManager().EmitTypedEvent(&basket.EventCreate{
		BasketDenom: denom,
		Curator:     msg.Curator,
	})
	return &basket.MsgCreateResponse{BasketDenom: denom}, err
}