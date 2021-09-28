package expand

import (
	"strings"

	"github.com/Platdot-Network/substrate-go/expand/bifrost"
	"github.com/Platdot-Network/substrate-go/expand/chainx"
	"github.com/Platdot-Network/substrate-go/expand/extra"
	"github.com/Platdot-Network/substrate-go/expand/polkadot"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

type IEventRecords interface {
	GetMultisigNewMultisig() []types.EventMultisigNewMultisig
	GetMultisigApproval() []types.EventMultisigApproval
	GetMultisigExecuted() []types.EventMultisigExecuted
	GetMultisigCancelled() []types.EventMultisigCancelled
	GetBalancesTransfer() []types.EventBalancesTransfer
	GetUtilityBatchCompleted() []types.EventUtilityBatchCompleted
	GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess
	GetSystemExtrinsicFailed() []extra.EventSystemExtrinsicFailed
	///ChainX
}

func DecodeEventRecords(meta *types.Metadata, rawData string, chainName string) (IEventRecords, error) {
	e := types.EventRecordsRaw(types.MustHexDecodeString(rawData))
	var ier IEventRecords
	switch strings.ToLower(chainName) {
	case "polkadot":
		var events polkadot.PolkadotEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "kusama":
		var events polkadot.PolkadotEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "westend":
		var events polkadot.PolkadotEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case ClientNameChainX:
		var events chainx.ChainXEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case ClientNameChainXAsset:
		var events chainx.ChainXEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "bifrost":
		//default
		var events bifrost.BifrostEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	default:
		var events chainx.ChainXEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	}

	return ier, nil
}

