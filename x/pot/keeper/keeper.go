package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/stratosnet/stratos-chain/x/register"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stratosnet/stratos-chain/x/pot/types"
)

// Keeper of the pot store
type Keeper struct {
	storeKey         sdk.StoreKey
	Cdc              *codec.Codec
	paramSpace       params.Subspace
	feeCollectorName string // name of the FeeCollector ModuleAccount
	BankKeeper       bank.Keeper
	SupplyKeeper     supply.Keeper
	AccountKeeper    auth.AccountKeeper
	StakingKeeper    staking.Keeper
	RegisterKeeper   register.Keeper
}

// NewKeeper creates a pot keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace, feeCollectorName string,
	bankKeeper bank.Keeper, supplyKeeper supply.Keeper, accountKeeper auth.AccountKeeper, stakingKeeper staking.Keeper,
	registerKeeper register.Keeper,
) Keeper {
	keeper := Keeper{
		Cdc:              cdc,
		storeKey:         key,
		paramSpace:       paramSpace.WithKeyTable(types.ParamKeyTable()),
		feeCollectorName: feeCollectorName,
		BankKeeper:       bankKeeper,
		SupplyKeeper:     supplyKeeper,
		AccountKeeper:    accountKeeper,
		StakingKeeper:    stakingKeeper,
		RegisterKeeper:   registerKeeper,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetVolumeReport returns the hash of volume report
//func (k Keeper) GetVolumeReport(ctx sdk.Context, reporter sdk.AccAddress) ([]byte, error) {
//	store := ctx.KVStore(k.storeKey)
//	bz := store.Get(types.VolumeReportStoreKey(reporter))
//	if bz == nil {
//		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress,
//			"key %s does not exist", hex.EncodeToString(types.VolumeReportStoreKey(reporter)))
//	}
//	return bz, nil
//}

func (k Keeper) GetVolumeReport(ctx sdk.Context, epoch sdk.Int) (res types.ReportRecord, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.VolumeReportStoreKey(epoch))
	if bz == nil {
		return types.ReportRecord{}, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"key %s does not exist", epoch)
	}
	k.Cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return res, nil
}

func (k Keeper) SetVolumeReport(ctx sdk.Context, epoch sdk.Int, reportRecord types.ReportRecord) {
	store := ctx.KVStore(k.storeKey)
	//storeKey := types.VolumeReportStoreKey(reporter)
	storeKey := types.VolumeReportStoreKey(epoch)
	bz := k.Cdc.MustMarshalBinaryLengthPrefixed(reportRecord)
	store.Set(storeKey, bz)
}

func (k Keeper) DeleteVolumeReport(ctx sdk.Context, key []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(key)
}

func (k Keeper) IsSPNode(ctx sdk.Context, addr sdk.AccAddress) (found bool) {
	_, found = k.RegisterKeeper.GetIndexingNode(ctx, addr)
	return found
}
