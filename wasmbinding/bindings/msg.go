//nolint:revive,stylecheck  // if we change the names of var-naming things here, we harm some kind of mapping.
package bindings

import (
	icqtypes "github.com/neutron-org/neutron/x/interchainqueries/types"
)

// ProtobufAny is a hack-struct to serialize protobuf Any message into JSON object
type ProtobufAny struct {
	TypeURL string `json:"type_url"`
	Value   []byte `json:"value"`
}

// NeutronMsg is used like a sum type to hold one of custom Neutron messages.
// Follow https://github.com/neutron-org/neutron-contracts/tree/main/packages/bindings/src/msg.rs
// for more information.
type NeutronMsg struct {
	SubmitTx                  *SubmitTx                  `json:"submit_tx,omitempty"`
	RegisterInterchainAccount *RegisterInterchainAccount `json:"register_interchain_account,omitempty"`
	RegisterInterchainQuery   *RegisterInterchainQuery   `json:"register_interchain_query,omitempty"`
	UpdateInterchainQuery     *UpdateInterchainQuery     `json:"update_interchain_query,omitempty"`
	RemoveInterchainQuery     *RemoveInterchainQuery     `json:"remove_interchain_query,omitempty"`
	AddAdmin                  *AddAdmin                  `json:"add_admin,omitempty"`
	SubmitProposal            *SubmitProposal            `json:"submit_proposal,omitempty"`
}

// SubmitTx submits interchain transaction on a remote chain.
type SubmitTx struct {
	ConnectionId        string        `json:"connection_id"`
	InterchainAccountId string        `json:"interchain_account_id"`
	Msgs                []ProtobufAny `json:"msgs"`
	Memo                string        `json:"memo"`
	Timeout             uint64        `json:"timeout"`
}

// SubmitTxResponse holds response from SubmitTx.
type SubmitTxResponse struct {
	// SequenceId is a channel's sequence_id for outgoing ibc packet. Unique per a channel.
	SequenceId uint64 `json:"sequence_id"`
	// Channel is a src channel on neutron side transaction was submitted from
	Channel string `json:"channel"`
}

// RegisterInterchainAccount creates account on remote chain.
type RegisterInterchainAccount struct {
	ConnectionId        string `json:"connection_id"`
	InterchainAccountId string `json:"interchain_account_id"`
}

// RegisterInterchainAccountResponse holds response for RegisterInterchainAccount.
type RegisterInterchainAccountResponse struct{}

// RegisterInterchainQuery creates a query for remote chain.
type RegisterInterchainQuery struct {
	QueryType          string            `json:"query_type"`
	Keys               []*icqtypes.KVKey `json:"keys"`
	TransactionsFilter string            `json:"transactions_filter"`
	ConnectionId       string            `json:"connection_id"`
	UpdatePeriod       uint64            `json:"update_period"`
}

type AddAdmin struct {
	Admin string `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
}

type AddAdminResponse struct {
}

type SubmitProposal struct {
	Proposals Proposals `json:"proposals"`
}

type Proposals struct {
	TextProposal                  *TextProposal                  `json:"text_proposal,omitempty"`
	ParamChangeProposal           *ParamChangeProposal           `json:"param_change_proposal,omitempty"`
	CommunityPoolSpendProposal    *CommunityPoolSpendProposal    `json:"community_pool_spend_proposal,omitempty"`
	ClientUpdateProposal          *ClientUpdateProposal          `json:"client_update_proposal,omitempty"`
	SoftwareUpgradeProposal       *SoftwareUpgradeProposal       `json:"software_upgrade_proposal,omitempty"`
	CancelSoftwareUpgradeProposal *CancelSoftwareUpgradeProposal `json:"cancel_software_upgrade_proposal,omitempty"`
}

type TextProposal struct {
	Title       string
	Description string
}

type ParamChangeProposal struct {
	Title       string
	Description string
	Changes     []ParamChange
}

type CommunityPoolSpendProposal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Recipient   string `json:"recipient"`
	Amount      uint64 `json:"amount"`
}

type ClientUpdateProposal struct {
	Title              string `json:"title"`
	Description        string `json:"description"`
	SubjectClientId    string `json:"subject_client_id"`
	SubstituteClientId string `json:"substitute_client_id"`
}
type SoftwareUpgradeProposal struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type CancelSoftwareUpgradeProposal struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type ParamChange struct {
	Subspace string `json:"subspace"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

// RegisterInterchainQueryResponse holds response for RegisterInterchainQuery
type RegisterInterchainQueryResponse struct {
	Id uint64 `json:"id"`
}

type RemoveInterchainQuery struct {
	QueryId uint64 `json:"query_id"`
}

type RemoveInterchainQueryResponse struct{}

type UpdateInterchainQuery struct {
	QueryId         uint64            `json:"query_id,omitempty"`
	NewKeys         []*icqtypes.KVKey `json:"new_keys,omitempty"`
	NewUpdatePeriod uint64            `json:"new_update_period,omitempty"`
}

type UpdateInterchainQueryResponse struct{}
