package core

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgCreateProject(t *testing.T) {
	t.Parallel()
	issuer := testutil.GenAddress()

	testCases := []struct {
		name   string
		src    MsgCreateProject
		expErr bool
	}{
		{
			"valid msg with project id",
			MsgCreateProject{
				Issuer:              issuer,
				ClassId:             "A00",
				Metadata:            "hello",
				ProjectJurisdiction: "AB-CDE FG1 345",
				ProjectId:           "A0",
			},
			false,
		},
		{
			"valid msg without project id",
			MsgCreateProject{
				Issuer:              issuer,
				ClassId:             "A00",
				Metadata:            "hello",
				ProjectJurisdiction: "AB-CDE FG1 345",
			},
			false,
		},
		{
			"invalid issuer",
			MsgCreateProject{
				Issuer:              "invalid address",
				ClassId:             "A00",
				Metadata:            "hello",
				ProjectJurisdiction: "AB-CDE FG1 345",
				ProjectId:           "A0",
			},
			true,
		},
		{
			"invalid project id",
			MsgCreateProject{
				Issuer:              issuer,
				ClassId:             "A00",
				Metadata:            "hello",
				ProjectJurisdiction: "AB-CDE FG1 345",
				ProjectId:           "A",
			},
			true,
		},
		{
			"invalid class id",
			MsgCreateProject{
				Issuer:              issuer,
				ClassId:             "ABCD",
				Metadata:            "hello",
				ProjectJurisdiction: "AB-CDE FG1 345",
				ProjectId:           "AB",
			},
			true,
		},
		{
			"invalid project jurisdiction",
			MsgCreateProject{
				Issuer:              issuer,
				ClassId:             "A01",
				Metadata:            "hello",
				ProjectJurisdiction: "abcd",
				ProjectId:           "AB",
			},
			true,
		},
		{
			"invalid: metadata is too large",
			MsgCreateProject{
				Issuer:              issuer,
				ClassId:             "A01",
				Metadata:            strings.Repeat("x", 288),
				ProjectJurisdiction: "AB-CDE FG1 345",
				ProjectId:           "AB",
			},
			true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.src.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
