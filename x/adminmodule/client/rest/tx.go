package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/admin-module/x/adminmodule/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router, phs []govrest.ProposalRESTHandler) {
	propSubRtr := r.PathPrefix("/adminmodule/proposals").Subrouter()
	for _, ph := range phs {
		propSubRtr.HandleFunc(fmt.Sprintf("/%s", ph.SubRoute), ph.Handler).Methods("POST")
	}

	r.HandleFunc("/adminmodule/addadmin", newAddAdminHandlerFn(clientCtx)).Methods("POST")
	r.HandleFunc("/adminmodule/deleteadmin", newDeleteAdminHandlerFn(clientCtx)).Methods("POST")
}

func newAddAdminHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddAdminReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		newAdmin, err := sdk.AccAddressFromBech32(req.Admin)

		msg := types.NewMsgAddAdmin(fromAddr, newAdmin)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func newDeleteAdminHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddAdminReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		removeAdmin, err := sdk.AccAddressFromBech32(req.Admin)

		msg := types.NewMsgDeleteAdmin(fromAddr, removeAdmin)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
