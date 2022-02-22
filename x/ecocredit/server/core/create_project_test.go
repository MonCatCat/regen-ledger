package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestValidProjectState(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	makeClass(t, s.ctx, s.stateStore, s.addr)
	res, err := s.k.CreateProject(s.ctx, &v1beta1.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "C01",
		Metadata:        nil,
		ProjectLocation: "US-NY",
		ProjectId:       "FOO",
	})
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, "FOO")

	project, err := s.stateStore.ProjectInfoStore().GetByName(s.ctx, "FOO")
	assert.NilError(t, err)
	assert.Equal(t, project.ProjectLocation, "US-NY")
}



func TestSequence(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	makeClass(t, s.ctx, s.stateStore, s.addr)
	res, err := s.k.CreateProject(s.ctx, &v1beta1.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "C01",
		Metadata:        nil,
		ProjectLocation: "US-NY",
		ProjectId:       "FOO",
	})
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, "FOO")

	res, err = s.k.CreateProject(s.ctx, &v1beta1.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "C01",
		Metadata:        nil,
		ProjectLocation: "US-NY",
		ProjectId:       "FOO",
	})
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, "FOO1")
}

func TestGeneratedProjectID(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	makeClass(t, s.ctx, s.stateStore, s.addr)
	res, err := s.k.CreateProject(s.ctx, &v1beta1.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "C01",
		Metadata:        nil,
		ProjectLocation: "US-NY",
		ProjectId:       "",
	})
	assert.NilError(t, err)
	assert.Equal(t, res.ProjectId, "C0101", "got project id: %s", res.ProjectId)
}

func TestBadClassID(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, err := s.k.CreateProject(s.ctx, &v1beta1.MsgCreateProject{
		Issuer:          s.addr.String(),
		ClassId:         "NOPE",
		ProjectLocation: "US-NY",
		ProjectId:       "",
	})
	assert.ErrorContains(t, err, "not found")
}

func makeClass(t *testing.T, ctx context.Context, ss ecocreditv1beta1.StateStore, addr types.AccAddress) {
	assert.NilError(t, ss.ClassInfoStore().Insert(ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C01",
		Admin:      addr,
		Metadata:   nil,
		CreditType: "C",
	}))
	assert.NilError(t, ss.ClassIssuerStore().Insert(ctx, &ecocreditv1beta1.ClassIssuer{
		ClassId: 1,
		Issuer:  addr,
	}))
}