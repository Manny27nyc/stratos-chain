package register

import (
	"encoding/hex"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stratosnet/stratos-chain/x/register/keeper"
	"github.com/stratosnet/stratos-chain/x/register/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreateResourceNode:
			return handleMsgCreateResourceNode(ctx, msg, k)
		case types.MsgRemoveResourceNode:
			return handleMsgRemoveResourceNodeNew(ctx, msg, k)
		case types.MsgUpdateResourceNode:
			return handleMsgUpdateResourceNode(ctx, msg, k)

		case types.MsgCreateIndexingNode:
			return handleMsgCreateIndexingNode(ctx, msg, k)
		case types.MsgRemoveIndexingNode:
			return handleMsgRemoveIndexingNodeNew(ctx, msg, k)
		case types.MsgUpdateIndexingNode:
			return handleMsgUpdateIndexingNode(ctx, msg, k)
		case types.MsgIndexingNodeRegistrationVote:
			return handleMsgIndexingNodeRegistrationVote(ctx, msg, k)

		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateResourceNode(ctx sdk.Context, msg types.MsgCreateResourceNode, k keeper.Keeper) (*sdk.Result, error) {
	// check to see if the pubkey or sender has been registered before
	if _, found := k.GetResourceNode(ctx, sdk.AccAddress(msg.PubKey.Address())); found {
		ctx.Logger().Error("Resource node already exist")
		return nil, ErrResourceNodePubKeyExists
	}
	if msg.Value.Denom != k.BondDenom(ctx) {
		return nil, ErrBadDenom
	}

	err := k.RegisterResourceNode(ctx, msg.NetworkID, msg.PubKey, msg.OwnerAddress, msg.Description, msg.NodeType, msg.Value)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateResourceNode,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.OwnerAddress.String()),
			sdk.NewAttribute(types.AttributeKeyNetworkAddress, sdk.AccAddress(msg.PubKey.Address()).String()),
			sdk.NewAttribute(types.AttributeKeyPubKey, hex.EncodeToString(msg.PubKey.Bytes())),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgCreateIndexingNode(ctx sdk.Context, msg types.MsgCreateIndexingNode, k keeper.Keeper) (*sdk.Result, error) {
	// check to see if the pubkey or sender has been registered before
	if _, found := k.GetIndexingNode(ctx, sdk.AccAddress(msg.PubKey.Address())); found {
		ctx.Logger().Error("Indexing node already exist")
		return nil, ErrIndexingNodePubKeyExists
	}
	if msg.Value.Denom != k.BondDenom(ctx) {
		return nil, ErrBadDenom
	}

	err := k.RegisterIndexingNode(ctx, msg.NetworkID, msg.PubKey, msg.OwnerAddress, msg.Description, msg.Value)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateIndexingNode,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.OwnerAddress.String()),
			sdk.NewAttribute(types.AttributeKeyNetworkAddress, sdk.AccAddress(msg.PubKey.Address()).String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgRemoveResourceNode(ctx sdk.Context, msg types.MsgRemoveResourceNode, k keeper.Keeper) (*sdk.Result, error) {
	resourceNode, found := k.GetResourceNode(ctx, msg.ResourceNodeAddress)
	if !found {
		return nil, ErrNoResourceNodeFound
	}

	err := k.SubtractResourceNodeStake(ctx, resourceNode, sdk.NewCoin(k.BondDenom(ctx), resourceNode.GetTokens()))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemoveResourceNode,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.OwnerAddress.String()),
			sdk.NewAttribute(types.AttributeKeyResourceNode, msg.ResourceNodeAddress.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgRemoveResourceNodeNew(ctx sdk.Context, msg types.MsgRemoveResourceNode, k keeper.Keeper) (*sdk.Result, error) {
	resourceNode, found := k.GetResourceNode(ctx, msg.ResourceNodeAddress)
	if !found {
		return nil, ErrNoResourceNodeFound
	}

	ctx.Logger().Info("11111111111")

	completionTime, err := k.DoRemoveResourceNode(ctx, resourceNode, resourceNode.Tokens)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Info("22222222222")
	completionTimeBz := types.ModuleCdc.MustMarshalBinaryLengthPrefixed(completionTime)
	//ctx.EventManager().EmitEvents(sdk.Events{
	//	sdk.NewEvent(
	//		types.EventTypeUnbond,
	//		sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress.String()),
	//		sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),
	//		sdk.NewAttribute(types.AttributeKeyCompletionTime, completionTime.Format(time.RFC3339)),
	//	),
	//	sdk.NewEvent(
	//		sdk.EventTypeMessage,
	//		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
	//		sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress.String()),
	//	),
	//})

	return &sdk.Result{Data: completionTimeBz, Events: ctx.EventManager().Events()}, nil
}

func handleMsgRemoveIndexingNode(ctx sdk.Context, msg types.MsgRemoveIndexingNode, k keeper.Keeper) (*sdk.Result, error) {
	indexingNode, found := k.GetIndexingNode(ctx, msg.IndexingNodeAddress)
	if !found {
		return nil, ErrNoIndexingNodeFound
	}
	err := k.SubtractIndexingNodeStake(ctx, indexingNode, sdk.NewCoin(k.BondDenom(ctx), indexingNode.GetTokens()))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemoveIndexingNode,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.OwnerAddress.String()),
			sdk.NewAttribute(types.AttributeKeyIndexingNode, msg.IndexingNodeAddress.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgRemoveIndexingNodeNew(ctx sdk.Context, msg types.MsgRemoveIndexingNode, k keeper.Keeper) (*sdk.Result, error) {
	indexingNode, found := k.GetIndexingNode(ctx, msg.IndexingNodeAddress)
	if !found {
		return nil, ErrNoIndexingNodeFound
	}

	completionTime, err := k.DoRemoveIndexingNode(ctx, indexingNode, indexingNode.Tokens)
	if err != nil {
		return nil, err
	}

	completionTimeBz := types.ModuleCdc.MustMarshalBinaryLengthPrefixed(completionTime)
	//ctx.EventManager().EmitEvents(sdk.Events{
	//	sdk.NewEvent(
	//		types.EventTypeRemoveIndexingNode,
	//		sdk.NewAttribute(sdk.AttributeKeySender, msg.OwnerAddress.String()),
	//		sdk.NewAttribute(types.AttributeKeyIndexingNode, msg.IndexingNodeAddress.String()),
	//	),
	//	sdk.NewEvent(
	//		sdk.EventTypeMessage,
	//		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
	//	),
	//})
	//return &sdk.Result{Events: ctx.EventManager().Events()}, nil

	return &sdk.Result{Data: completionTimeBz, Events: ctx.EventManager().Events()}, nil
}

func handleMsgIndexingNodeRegistrationVote(ctx sdk.Context, msg types.MsgIndexingNodeRegistrationVote, k keeper.Keeper) (*sdk.Result, error) {
	nodeToApprove, found := k.GetIndexingNode(ctx, msg.CandidateNetworkAddress)
	if !found {
		return nil, ErrNoIndexingNodeFound
	}
	if !nodeToApprove.GetOwnerAddr().Equals(msg.CandidateOwnerAddress) {
		return nil, ErrInvalidOwnerAddr
	}

	voter, found := k.GetIndexingNode(ctx, msg.VoterNetworkAddress)
	if !found {
		return nil, ErrInvalidApproverAddr
	}
	if !voter.Status.Equal(sdk.Bonded) || voter.IsSuspended() {
		return nil, ErrInvalidApproverStatus
	}

	nodeStatus, err := k.HandleVoteForIndexingNodeRegistration(ctx, msg.CandidateNetworkAddress, msg.CandidateOwnerAddress, msg.Opinion, msg.VoterNetworkAddress)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeIndexingNodeRegistrationVote,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.VoterNetworkAddress.String()),
			sdk.NewAttribute(types.AttributeKeyCandidateNetworkAddress, msg.CandidateNetworkAddress.String()),
			sdk.NewAttribute(types.AttributeKeyCandidateStatus, nodeStatus.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgUpdateResourceNode(ctx sdk.Context, msg types.MsgUpdateResourceNode, k keeper.Keeper) (*sdk.Result, error) {
	err := k.UpdateResourceNode(ctx, msg.NetworkID, msg.Description, msg.NodeType, msg.NetworkAddress, msg.OwnerAddress)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateResourceNode,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.OwnerAddress.String()),
			sdk.NewAttribute(types.AttributeKeyNetworkAddress, msg.NetworkAddress.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgUpdateIndexingNode(ctx sdk.Context, msg types.MsgUpdateIndexingNode, k keeper.Keeper) (*sdk.Result, error) {
	err := k.UpdateIndexingNode(ctx, msg.NetworkID, msg.Description, msg.NetworkAddress, msg.OwnerAddress)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateIndexingNode,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.OwnerAddress.String()),
			sdk.NewAttribute(types.AttributeKeyNetworkAddress, msg.NetworkAddress.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
