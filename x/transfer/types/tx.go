package types

import (
	"github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (msg *MsgTransfer) ValidateBasic() error {
	if err := msg.Fee.Validate(); err != nil {
		return err
	}

	sdkMsg := types.NewMsgTransfer(msg.SourcePort, msg.SourceChannel, msg.Token, msg.Sender, msg.Receiver, msg.TimeoutHeight, msg.TimeoutTimestamp, msg.Memo)
	return sdkMsg.ValidateBasic()
}

func (msg *MsgTransfer) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{fromAddress}
}
