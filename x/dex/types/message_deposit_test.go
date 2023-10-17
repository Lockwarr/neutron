package types_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/neutron-org/neutron/testutil/common/sample"
	. "github.com/neutron-org/neutron/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestMsgDeposit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeposit
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgDeposit{
				Creator:         "invalid_address",
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				AmountsA:        []math.Int{math.OneInt()},
				AmountsB:        []math.Int{math.OneInt()},
			},
			err: ErrInvalidAddress,
		},
		{
			name: "invalid receiver",
			msg: MsgDeposit{
				Creator:         sample.AccAddress(),
				Receiver:        "invalid address",
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				AmountsA:        []math.Int{math.OneInt()},
				AmountsB:        []math.Int{math.OneInt()},
			},
			err: ErrInvalidAddress,
		},
		{
			name: "invalid fee indexes length",
			msg: MsgDeposit{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{},
				AmountsA:        []math.Int{},
				AmountsB:        []math.Int{},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid tick indexes length",
			msg: MsgDeposit{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{0},
				AmountsA:        []math.Int{},
				AmountsB:        []math.Int{},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid amounts A length",
			msg: MsgDeposit{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{},
				AmountsA:        []math.Int{math.OneInt()},
				AmountsB:        []math.Int{},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid amounts B length",
			msg: MsgDeposit{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{},
				AmountsA:        []math.Int{},
				AmountsB:        []math.Int{math.OneInt()},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid no deposit",
			msg: MsgDeposit{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{},
				AmountsA:        []math.Int{},
				AmountsB:        []math.Int{},
			},
			err: ErrZeroDeposit,
		},
		{
			name: "invalid no deposit",
			msg: MsgDeposit{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				AmountsA:        []math.Int{math.ZeroInt()},
				AmountsB:        []math.Int{math.ZeroInt()},
			},
			err: ErrZeroDeposit,
		},
		{
			name: "valid msg",
			msg: MsgDeposit{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				AmountsA:        []math.Int{math.OneInt()},
				AmountsB:        []math.Int{math.OneInt()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
