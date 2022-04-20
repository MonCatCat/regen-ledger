package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// CreateProject creates a new project for a specific credit class.
func (k Keeper) CreateProject(ctx context.Context, req *core.MsgCreateProject) (*core.MsgCreateProjectResponse, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	classID := req.ClassId
	classInfo, err := k.stateStore.ClassInfoTable().GetById(ctx, classID)
	if err != nil {
		return nil, err
	}

	adminAddress, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err
	}

	err = k.assertClassIssuer(ctx, classInfo.Key, adminAddress)
	if err != nil {
		return nil, err
	}

	projectID := req.ProjectId
	if projectID == "" {
		exists := true
		for exists {
			projectID, err = k.genProjectID(ctx, classInfo.Key, classInfo.Id)
			if err != nil {
				return nil, err
			}

			exists, err = k.stateStore.ProjectInfoTable().HasByClassKeyId(ctx, classInfo.Key, projectID)
			if err != nil {
				return nil, err
			}

			sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/MsgCreateProject id iteration")
		}
	}

	if err = k.stateStore.ProjectInfoTable().Insert(ctx, &api.ProjectInfo{
		Id:                  projectID,
		Admin:               adminAddress,
		ClassKey:            classInfo.Key,
		ProjectJurisdiction: req.ProjectJurisdiction,
		Metadata:            req.Metadata,
	}); err != nil {
		return nil, err
	}

	if err := sdkCtx.EventManager().EmitTypedEvent(&core.EventCreateProject{
		ProjectId:           projectID,
		Admin:               adminAddress.String(),
		ClassId:             classID,
		ProjectJurisdiction: req.ProjectJurisdiction,
	}); err != nil {
		return nil, err
	}

	return &core.MsgCreateProjectResponse{
		ProjectId: projectID,
	}, nil
}

// genProjectID generates a projectID when no projectID was given for CreateProject.
// The ID is generated by concatenating the classID and a sequence number.
func (k Keeper) genProjectID(ctx context.Context, classKey uint64, classID string) (string, error) {
	var nextID uint64
	projectSeqNo, err := k.stateStore.ProjectSequenceTable().Get(ctx, classKey)
	switch err {
	case ormerrors.NotFound:
		nextID = 1
	case nil:
		nextID = projectSeqNo.NextSequence
	default:
		return "", err
	}

	if err = k.stateStore.ProjectSequenceTable().Save(ctx, &api.ProjectSequence{
		ClassKey:     classKey,
		NextSequence: nextID + 1,
	}); err != nil {
		return "", err
	}

	return ecocredit.FormatProjectID(classID, nextID), nil
}
