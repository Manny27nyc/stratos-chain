package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalid                            = sdkerrors.Register(ModuleName, 1, "error invalid")
	ErrEmptyNetworkAddr                   = sdkerrors.Register(ModuleName, 2, "missing network address")
	ErrEmptyOwnerAddr                     = sdkerrors.Register(ModuleName, 3, "missing owner address")
	ErrValueNegative                      = sdkerrors.Register(ModuleName, 4, "value must be positive")
	ErrEmptyDescription                   = sdkerrors.Register(ModuleName, 5, "description must be not empty")
	ErrEmptyMoniker                       = sdkerrors.Register(ModuleName, 6, "moniker must be not empty")
	ErrEmptyResourceNodeAddr              = sdkerrors.Register(ModuleName, 7, "missing resource node address")
	ErrEmptyIndexingNodeAddr              = sdkerrors.Register(ModuleName, 8, "missing indexing node address")
	ErrBadDenom                           = sdkerrors.Register(ModuleName, 9, "invalid coin denomination")
	ErrResourceNodePubKeyExists           = sdkerrors.Register(ModuleName, 10, "resource node already exist for this pubkey; must use new resource node pubkey")
	ErrIndexingNodePubKeyExists           = sdkerrors.Register(ModuleName, 11, "indexing node already exist for this pubkey; must use new indexing node pubkey")
	ErrNoResourceNodeFound                = sdkerrors.Register(ModuleName, 12, "resource node does not exist")
	ErrNoIndexingNodeFound                = sdkerrors.Register(ModuleName, 13, "indexing node does not exist")
	ErrNoOwnerAccountFound                = sdkerrors.Register(ModuleName, 14, "account of owner does not exist")
	ErrInsufficientBalance                = sdkerrors.Register(ModuleName, 15, "insufficient balance")
	ErrNodeType                           = sdkerrors.Register(ModuleName, 16, "node type(s) not supported")
	ErrEmptyCandidateNetworkAddr          = sdkerrors.Register(ModuleName, 17, "missing candidate network address")
	ErrEmptyCandidateOwnerAddr            = sdkerrors.Register(ModuleName, 18, "missing candidate owner address")
	ErrEmptyVoterNetworkAddr              = sdkerrors.Register(ModuleName, 19, "missing voter network address")
	ErrEmptyVoterOwnerAddr                = sdkerrors.Register(ModuleName, 20, "missing voter owner address")
	ErrSameAddr                           = sdkerrors.Register(ModuleName, 21, "node address should not same as the voter address")
	ErrInvalidOwnerAddr                   = sdkerrors.Register(ModuleName, 22, "invalid owner address")
	ErrInvalidVoterAddr                   = sdkerrors.Register(ModuleName, 23, "invalid voter address")
	ErrInvalidVoterStatus                 = sdkerrors.Register(ModuleName, 24, "invalid voter status")
	ErrNoRegistrationVotePoolFound        = sdkerrors.Register(ModuleName, 25, "registration pool does not exist")
	ErrDuplicateVoting                    = sdkerrors.Register(ModuleName, 26, "duplicate voting")
	ErrVoteExpired                        = sdkerrors.Register(ModuleName, 27, "vote expired")
	ErrInsufficientBalanceOfBondedPool    = sdkerrors.Register(ModuleName, 28, "insufficient balance of bonded pool")
	ErrInsufficientBalanceOfNotBondedPool = sdkerrors.Register(ModuleName, 29, "insufficient balance of not bonded pool")
	ErrSubAllTokens                       = sdkerrors.Register(ModuleName, 30, "can not sub all tokens since the node is still bonded")
	ErrED25519InvalidPubKey               = sdkerrors.Register(ModuleName, 31, "ED25519 public keys are unsupported")
	ErrEmptyNodeId                        = sdkerrors.Register(ModuleName, 32, "missing node id")
	ErrEmptyPubKey                        = sdkerrors.Register(ModuleName, 33, "missing public key")
	ErrInvalidGenesisToken                = sdkerrors.Register(ModuleName, 34, "invalid genesis token")
)
