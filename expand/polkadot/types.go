package polkadot

import (
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
	"github.com/Platdot-Network/substrate-go/expand/base"
	"github.com/Platdot-Network/substrate-go/expand/bridge"
	"github.com/Platdot-Network/substrate-go/expand/extra"
)

type PolkadotEventRecords struct {
	base.BaseEventRecords
	bridge.BridgeEvents
}

func (p PolkadotEventRecords) GetMultisigNewMultisig() []types.EventMultisigNewMultisig {
	return p.Multisig_NewMultisig
}

func (p PolkadotEventRecords) GetMultisigApproval() []types.EventMultisigApproval {
	return p.Multisig_MultisigApproval
}

func (p PolkadotEventRecords) GetMultisigExecuted() []types.EventMultisigExecuted {
	return p.Multisig_MultisigExecuted
}

func (p PolkadotEventRecords) GetMultisigCancelled() []types.EventMultisigCancelled {
	return p.Multisig_MultisigCancelled
}

func (p PolkadotEventRecords) GetUtilityBatchCompleted() []types.EventUtilityBatchCompleted {
	return p.Utility_BatchCompleted
}

func (p PolkadotEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return p.Balances_Transfer
}

func (p PolkadotEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return p.System_ExtrinsicSuccess
}

func (p PolkadotEventRecords) GetSystemExtrinsicFailed() []extra.EventSystemExtrinsicFailed {
	return p.System_ExtrinsicFailed
}
