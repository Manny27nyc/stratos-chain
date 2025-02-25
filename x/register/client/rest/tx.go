package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	stratos "github.com/stratosnet/stratos-chain/types"
	"github.com/stratosnet/stratos-chain/x/register/types"
	"net/http"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/register/createResourceNode",
		postCreateResourceNodeHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/register/removeResourceNode",
		postRemoveResourceNodeHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/register/updateResourceNode",
		postUpdateResourceNodeHandlerFn(cliCtx),
	).Methods("POST")

	r.HandleFunc(
		"/register/createIndexingNode",
		postCreateIndexingNodeHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/register/removeIndexingNode",
		postRemoveIndexingNodeHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/register/updateIndexingNode",
		postUpdateIndexingNodeHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/register/indexingNodeRegVote",
		postIndexingNodeRegVoteFn(cliCtx),
	).Methods("POST")
}

type (
	CreateResourceNodeRequest struct {
		BaseReq     rest.BaseReq      `json:"base_req" yaml:"base_req"`
		NetworkID   string            `json:"network_id" yaml:"network_id"`
		PubKey      string            `json:"pubkey" yaml:"pubkey"` // in bech32
		Amount      sdk.Coin          `json:"amount" yaml:"amount"`
		Description types.Description `json:"description" yaml:"description"`
		NodeType    int               `json:"node_type" yaml:"node_type"`
	}

	RemoveResourceNodeRequest struct {
		BaseReq             rest.BaseReq `json:"base_req" yaml:"base_req"`
		ResourceNodeAddress string       `json:"resource_node_address" yaml:"resource_node_address"` // in bech32
	}

	UpdateResourceNodeRequest struct {
		BaseReq        rest.BaseReq      `json:"base_req" yaml:"base_req"`
		NetworkID      string            `json:"network_id" yaml:"network_id"`
		Description    types.Description `json:"description" yaml:"description"`
		NodeType       int               `json:"node_type" yaml:"node_type"`
		NetworkAddress string            `json:"network_address" yaml:"network_address"`
	}

	CreateIndexingNodeRequest struct {
		BaseReq     rest.BaseReq      `json:"base_req" yaml:"base_req"`
		NetworkID   string            `json:"network_id" yaml:"network_id"`
		PubKey      string            `json:"pubkey" yaml:"pubkey"` // in bech32
		Amount      sdk.Coin          `json:"amount" yaml:"amount"`
		Description types.Description `json:"description" yaml:"description"`
	}

	RemoveIndexingNodeRequest struct {
		BaseReq             rest.BaseReq `json:"base_req" yaml:"base_req"`
		IndexingNodeAddress string       `json:"indexing_node_address" yaml:"indexing_node_address"` // in bech32
	}

	UpdateIndexingNodeRequest struct {
		BaseReq        rest.BaseReq      `json:"base_req" yaml:"base_req"`
		NetworkID      string            `json:"network_id" yaml:"network_id"`
		Description    types.Description `json:"description" yaml:"description"`
		NetworkAddress string            `json:"network_address" yaml:"network_address"`
	}

	IndexingNodeRegVoteRequest struct {
		BaseReq                 rest.BaseReq `json:"base_req" yaml:"base_req"`
		CandidateNetworkAddress string       `json:"candidate_network_address" yaml:"candidate_network_address"`
		CandidateOwnerAddress   string       `json:"candidate_owner_address" yaml:"candidate_owner_address"`
		Opinion                 bool         `json:"opinion" yaml:"opinion"`
		VoterNetworkAddress     string       `json:"voter_network_address" yaml:"voter_network_address"`
	}
)

func postCreateResourceNodeHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateResourceNodeRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		pubKey, err := stratos.GetPubKeyFromBech32(stratos.Bech32PubKeyTypeSdsP2PPub, req.PubKey)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		nodeTypeRef := req.NodeType
		ownerAddr, er := sdk.AccAddressFromBech32(req.BaseReq.From)
		if er != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, er.Error())
			return
		}
		if t := types.NodeType(nodeTypeRef).Type(); t == "UNKNOWN" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "node type(s) not supported")
			return
		}
		msg := types.NewMsgCreateResourceNode(req.NetworkID, pubKey, req.Amount, ownerAddr, req.Description,
			fmt.Sprintf("%d: %s", nodeTypeRef, types.NodeType(nodeTypeRef).Type()))
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postCreateIndexingNodeHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateIndexingNodeRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		pubKey, err := stratos.GetPubKeyFromBech32(stratos.Bech32PubKeyTypeSdsP2PPub, req.PubKey)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		ownerAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgCreateIndexingNode(req.NetworkID, pubKey, req.Amount, ownerAddr, req.Description)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postRemoveResourceNodeHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RemoveResourceNodeRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		nodeAddr, err := sdk.AccAddressFromBech32(req.ResourceNodeAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		ownerAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgRemoveResourceNode(nodeAddr, ownerAddr)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postRemoveIndexingNodeHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RemoveIndexingNodeRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		nodeAddr, err := sdk.AccAddressFromBech32(req.IndexingNodeAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		ownerAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgRemoveIndexingNode(nodeAddr, ownerAddr)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postUpdateResourceNodeHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateResourceNodeRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		nodeTypeRef := req.NodeType

		networkAddr, err := sdk.AccAddressFromBech32(req.NetworkAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		ownerAddr, er := sdk.AccAddressFromBech32(req.BaseReq.From)
		if er != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, er.Error())
			return
		}
		if t := types.NodeType(nodeTypeRef).Type(); t == "UNKNOWN" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "node type(s) not supported")
			return
		}
		msg := types.NewMsgUpdateResourceNode(req.NetworkID, req.Description,
			fmt.Sprintf("%d: %s", nodeTypeRef, types.NodeType(nodeTypeRef).Type()), networkAddr, ownerAddr)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postUpdateIndexingNodeHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateIndexingNodeRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		networkAddr, err := sdk.AccAddressFromBech32(req.NetworkAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		ownerAddr, er := sdk.AccAddressFromBech32(req.BaseReq.From)
		if er != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, er.Error())
			return
		}

		msg := types.NewMsgUpdateIndexingNode(req.NetworkID, req.Description, networkAddr, ownerAddr)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postIndexingNodeRegVoteFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req IndexingNodeRegVoteRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		candidateNetworkAddr, err := sdk.AccAddressFromBech32(req.CandidateNetworkAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		candidateOwnerAddr, err := sdk.AccAddressFromBech32(req.CandidateOwnerAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		voteOpinion := types.VoteOpinionFromBool(req.Opinion)

		voterNetworkAddr, err := sdk.AccAddressFromBech32(req.VoterNetworkAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		voterOwnerAddr, er := sdk.AccAddressFromBech32(req.BaseReq.From)
		if er != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, er.Error())
			return
		}

		msg := types.NewMsgIndexingNodeRegistrationVote(candidateNetworkAddr, candidateOwnerAddr, voteOpinion, voterNetworkAddr, voterOwnerAddr)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
