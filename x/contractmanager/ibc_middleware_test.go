package contractmanager_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	test_keeper "github.com/neutron-org/neutron/testutil/interchaintxs/keeper"
	mock_types "github.com/neutron-org/neutron/testutil/mocks/contractmanager/types"
	"github.com/neutron-org/neutron/x/contractmanager/types"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	ShouldNotBeWrittenKey = []byte("shouldnotkey")
	ShouldNotBeWritten    = []byte("should not be written")
	ShouldBeWritten       = []byte("should be written")
	TestOwnerAddress      = "neutron17dtl0mjt3t77kpuhg2edqzjpszulwhgzcdvagh"
)

func ShouldBeWrittenKey(suffix string) []byte {
	return append([]byte("shouldkey"), []byte(suffix)...)
}

func TestSudo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmKeeper := mock_types.NewMockContractManagerKeeper(ctrl)
	wmKeeper := mock_types.NewMockWasmKeeper(ctrl)
	middleware, infCtx, storeKey := test_keeper.NewSudoLimitWrapper(t, cmKeeper, wmKeeper)
	st := infCtx.KVStore(storeKey)

	// at this point the payload struct does not matter
	msg := []byte("sudo_payload")
	contractAddress := sdk.AccAddress{}

	//  success during Sudo
	ctx := infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().GetParams(ctx).Return(types.Params{SudoCallGasLimit: 10000})
	wmKeeper.EXPECT().Sudo(gomock.AssignableToTypeOf(ctx), contractAddress, msg).Do(func(cachedCtx sdk.Context, senderAddress sdk.AccAddress, msg []byte) {
		st := cachedCtx.KVStore(storeKey)
		st.Set(ShouldBeWrittenKey("sudo"), ShouldBeWritten)
	}).Return(nil, nil)
	_, err := middleware.Sudo(ctx, contractAddress, msg)
	require.NoError(t, err)
	require.Equal(t, ShouldBeWritten, st.Get(ShouldBeWrittenKey("sudo")))

	//  error during Sudo
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().GetParams(ctx).Return(types.Params{SudoCallGasLimit: 10000})
	cmKeeper.EXPECT().AddContractFailure(ctx, contractAddress.String(), msg)
	wmKeeper.EXPECT().Sudo(gomock.AssignableToTypeOf(ctx), contractAddress, msg).Do(func(cachedCtx sdk.Context, senderAddress sdk.AccAddress, msg []byte) {
		st := cachedCtx.KVStore(storeKey)
		st.Set(ShouldNotBeWrittenKey, ShouldNotBeWritten)
	}).Return(nil, fmt.Errorf("sudo error"))
	_, err = middleware.Sudo(ctx, contractAddress, msg)
	require.Error(t, err)
	require.Nil(t, st.Get(ShouldNotBeWrittenKey))

	// ou of gas during Sudo
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().GetParams(ctx).Return(types.Params{SudoCallGasLimit: 10000})
	cmKeeper.EXPECT().AddContractFailure(ctx, contractAddress.String(), msg)
	wmKeeper.EXPECT().Sudo(gomock.AssignableToTypeOf(ctx), contractAddress, msg).Do(func(cachedCtx sdk.Context, senderAddress sdk.AccAddress, msg []byte) {
		st := cachedCtx.KVStore(storeKey)
		st.Set(ShouldNotBeWrittenKey, ShouldNotBeWritten)
		cachedCtx.GasMeter().ConsumeGas(10001, "heavy calculations")
	})
	_, err = middleware.Sudo(ctx, contractAddress, msg)
	require.ErrorContains(t, err, "{heavy calculations}: panic")
	require.Nil(t, st.Get(ShouldNotBeWrittenKey))
}
