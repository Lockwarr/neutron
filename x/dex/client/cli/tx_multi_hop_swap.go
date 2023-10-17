package cli

import (
	"strconv"
	"strings"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "cosmossdk.io/errors"
	math_utils "github.com/neutron-org/neutron/utils/math"
	"github.com/neutron-org/neutron/x/dex/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMultiHopSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multi-hop-swap [receiver] [routes] [amount-in] [exit-limit-price] [pick-best-route]",
		Short: "Broadcast message multiHopSwap",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiever := args[0]
			argRoutes := strings.Split(args[1], ";")
			argAmountIn := args[2]
			argExitLimitPrice := args[3]
			argPickBest := args[4]

			routesArr := make([][]string, len(argRoutes))
			for i, route := range argRoutes {
				routesArr[i] = strings.Split(route, ",")
			}

			amountInInt, ok := math.NewIntFromString(argAmountIn)
			if !ok {
				return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Invalid value for amount-in")
			}

			exitLimitPriceDec, err := math_utils.NewPrecDecFromStr(argExitLimitPrice)
			if err != nil {
				return sdkerrors.Wrapf(types.ErrIntOverflowTx, "Invalid value for exit-limit-price")
			}

			pickBest, err := strconv.ParseBool(argPickBest)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMultiHopSwap(
				clientCtx.GetFromAddress().String(),
				argReceiever,
				routesArr,
				amountInInt,
				exitLimitPriceDec,
				pickBest,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
