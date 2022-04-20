package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgRetire(t *testing.T) {
	t.Parallel()

	addr1 := testutil.GenAddress()

	tests := map[string]struct {
		src    MsgRetire
		expErr bool
	}{
		"valid msg": {
			src: MsgRetire{
				Holder: addr1,
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: false,
		},
		"invalid msg without holder": {
			src: MsgRetire{
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg with wrong holder address": {
			src: MsgRetire{
				Holder: "wrongHolder",
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without credits": {
			src: MsgRetire{
				Holder:       addr1,
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without Credits.BatchDenom": {
			src: MsgRetire{
				Holder: addr1,
				Credits: []*MsgRetire_RetireCredits{
					{
						Amount: "10",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without Credits.Amount": {
			src: MsgRetire{
				Holder: addr1,
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: batchDenom,
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg with wrong Credits.Amount": {
			src: MsgRetire{
				Holder: addr1,
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: batchDenom,
						Amount:     "abc",
					},
				},
				Jurisdiction: "AB-CDE FG1 345",
			},
			expErr: true,
		},
		"invalid msg without jurisdiction": {
			src: MsgRetire{
				Holder: addr1,
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong jurisdiction": {
			src: MsgRetire{
				Holder: addr1,
				Credits: []*MsgRetire_RetireCredits{
					{
						BatchDenom: batchDenom,
						Amount:     "10",
					},
				},
				Jurisdiction: "wrongJurisdiction",
			},
			expErr: true,
		},
	}

	for msg, test := range tests {
		t.Run(msg, func(t *testing.T) {
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
