package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/ethereum/go-ethereum/common"

	"github.com/bitdao-io/bitos/x/erc20/types"

	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

// OnRecvPacket will get the denom name from ibc ,generate by port/channel/denom
func (k Keeper) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	ack exported.Acknowledgement,
) exported.Acknowledgement {
	var err error
	cctx, write := ctx.CacheContext()

	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		ctx.Logger().Error(err.Error())
		return nil
	}
	transferAmount, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		return nil
	}
	receiver, _ := sdk.AccAddressFromBech32(data.Receiver)

	denom, err := types.IBCDenom(packet.GetDestPort(), packet.GetDestChannel(), data.Denom)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil
	}

	if !k.IsDenomRegistered(ctx, denom) {
		return nil
	}
	msg := types.NewMsgConvertCoin(
		sdk.NewCoin(denom, transferAmount),
		common.BytesToAddress(receiver.Bytes()),
		receiver,
	)
	// use cctx to ConvertCoin
	context := sdk.WrapSDKContext(cctx)
	_, err = k.ConvertCoin(context, msg)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil
	}
	// get minting state
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	pair, err := k.MintingEnabled(ctx, sender, receiver.Bytes(), msg.Coin.Denom)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil
	}
	write()
	emit(ctx,
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
		sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Coin.Amount.String()),
		sdk.NewAttribute(types.AttributeKeyCosmosCoin, msg.Coin.Denom),
		sdk.NewAttribute(types.AttributeKeyERC20Token, pair.Erc20Address),
	)
	return nil
}

func (k Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
) error {
	// nothing to do
	return nil
}

func (k Keeper) SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet exported.PacketI) error {
	return k.ics4Wrapper.SendPacket(ctx, channelCap, packet)
}

func (k Keeper) WriteAcknowledgement(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet exported.PacketI, ack exported.Acknowledgement) error {
	return k.ics4Wrapper.WriteAcknowledgement(ctx, channelCap, packet, ack)
}

func emit(ctx sdk.Context, attributes ...sdk.Attribute) {
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertCoin,
				attributes...,
			),
		},
	)
}
