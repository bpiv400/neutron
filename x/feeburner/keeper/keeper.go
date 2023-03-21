package keeper

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	consumertypes "github.com/cosmos/interchain-security/x/ccv/consumer/types"
	"github.com/neutron-org/neutron/x/feeburner/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

var KeyBurnedFees = []byte("BurnedFees")

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// RecordBurnedFees adds `amount` to the total amount of burned NTRN tokens
func (k Keeper) RecordBurnedFees(ctx sdk.Context, amount sdk.Coin) {
	store := ctx.KVStore(k.storeKey)

	totalBurnedNeutronsAmount := k.GetTotalBurnedNeutronsAmount(ctx)
	totalBurnedNeutronsAmount.Coin = totalBurnedNeutronsAmount.Coin.Add(amount)

	store.Set(KeyBurnedFees, k.cdc.MustMarshal(&totalBurnedNeutronsAmount))
}

// GetTotalBurnedNeutronsAmount gets the total burned amount of NTRN tokens
func (k Keeper) GetTotalBurnedNeutronsAmount(ctx sdk.Context) types.TotalBurnedNeutronsAmount {
	store := ctx.KVStore(k.storeKey)

	var totalBurnedNeutronsAmount types.TotalBurnedNeutronsAmount
	bzTotalBurnedNeutronsAmount := store.Get(KeyBurnedFees)
	if bzTotalBurnedNeutronsAmount != nil {
		k.cdc.MustUnmarshal(bzTotalBurnedNeutronsAmount, &totalBurnedNeutronsAmount)
	}

	if totalBurnedNeutronsAmount.Coin.Denom == "" {
		totalBurnedNeutronsAmount.Coin = sdk.NewCoin(k.GetParams(ctx).NeutronDenom, sdk.NewInt(0))
	}

	return totalBurnedNeutronsAmount
}

// BurnAndDistribute is an important part of tokenomics. It does few things:
// 1. Burns NTRN fee coins distributed to consumertypes.ConsumerRedistributeName in ICS (https://github.com/cosmos/interchain-security/blob/v0.2.0/x/ccv/consumer/keeper/distribution.go#L17)
// 2. Updates total amount of burned NTRN coins
// 3. Sends non-NTRN fee tokens to treasury contract address
// Panics if no `consumertypes.ConsumerRedistributeName` module found OR could not burn NTRN tokens
func (k Keeper) BurnAndDistribute(ctx sdk.Context) {
	moduleAddr := k.accountKeeper.GetModuleAddress(consumertypes.ConsumerRedistributeName)
	if moduleAddr == nil {
		panic("ConsumerRedistributeName must have module address")
	}

	params := k.GetParams(ctx)
	balances := k.bankKeeper.GetAllBalances(ctx, moduleAddr)
	fundsForTreasury := make(sdk.Coins, 0, len(balances))

	for _, balance := range balances {
		if !balance.IsZero() {
			if balance.Denom == params.NeutronDenom {
				err := k.bankKeeper.BurnCoins(ctx, consumertypes.ConsumerRedistributeName, sdk.Coins{balance})
				if err != nil {
					panic(sdkerrors.Wrapf(err, "failed to burn NTRN tokens during fee processing"))
				}

				k.RecordBurnedFees(ctx, balance)
			} else {
				fundsForTreasury = append(fundsForTreasury, balance)
			}
		}
	}

	if len(fundsForTreasury) > 0 {
		err := k.bankKeeper.SendCoins(
			ctx,
			moduleAddr, sdk.MustAccAddressFromBech32(params.TreasuryAddress),
			fundsForTreasury,
		)
		if err != nil {
			panic(sdkerrors.Wrapf(err, "failed sending funds to treasury"))
		}
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
