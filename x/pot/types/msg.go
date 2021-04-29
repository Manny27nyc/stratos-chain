package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const MsgType = "node_volume_report"

// verify interface at compile time
var (
	_ sdk.Msg = &MsgVolumeReport{}
)

type MsgVolumeReport struct {
	NodesVolume         []SingleNodeVolume `json:"nodes_volume" yaml:"nodes_volume"`               // volume report
	Reporter            sdk.AccAddress     `json:"volume_reporter" yaml:"volume_reporter"`         // volume reporter
	Epoch               sdk.Int            `json:"volume_report_epoch" yaml:"volume_report_epoch"` // volume report epoch
	ReportReferenceHash string             `json:"volume_report_hash" yaml:"volume_report_hash"`   // volume report reference
}

// NewMsgVolumeReport creates a new Msg<Action> instance
func NewMsgVolumeReport(
	nodesVolume []SingleNodeVolume,
	reporter sdk.AccAddress,
	epoch sdk.Int,
	reportReferenceHash string,
) MsgVolumeReport {
	return MsgVolumeReport{
		NodesVolume:         nodesVolume,
		Reporter:            reporter,
		Epoch:               epoch,
		ReportReferenceHash: reportReferenceHash,
	}
}

// Route Implement
func (msg MsgVolumeReport) Route() string { return RouterKey }

// GetSigners Implement
func (msg MsgVolumeReport) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Reporter}
}

// Type Implement
func (msg MsgVolumeReport) Type() string { return MsgType }

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgVolumeReport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgVolumeReport) ValidateBasic() error {
	if msg.Reporter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing reporter address")
	}
	if !(len(msg.NodesVolume) > 0) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no node reports volume")
	}

	if !(msg.Epoch.IsPositive()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid report epoch")
	}

	if !(len(msg.ReportReferenceHash) > 0) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid report reference hash")
	}

	for _, item := range msg.NodesVolume {
		if item.Volume.IsNegative() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "report volume is negative")
		}
		if item.NodeAddress.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing node address")
		}
	}
	return nil
}
