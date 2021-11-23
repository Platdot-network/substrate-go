package extra

import (
	gsrpcTypes "github.com/Platdot-Network/go-substrate-rpc-client/v3/types"
)

type TempMemo struct {
	Account string
	ParaId  uint32
	Memo    string
}

type OriginEvents struct {
	OriginEventRecord
}

type OriginEventRecord struct {
	// refer origin
	Claims_Claimed              []gsrpcTypes.EventClaimsClaimed      //nolint:stylecheck,golint
	Balances_Endowed            []gsrpcTypes.EventBalancesEndowed    //nolint:stylecheck,golint
	Balances_DustLost           []gsrpcTypes.EventBalancesDustLost   //nolint:stylecheck,golint
	Balances_Transfer           []gsrpcTypes.EventBalancesTransfer   //nolint:stylecheck,golint
	Balances_BalanceSet         []gsrpcTypes.EventBalancesBalanceSet //nolint:stylecheck,golint
	Balances_Reserved           []gsrpcTypes.EventBalancesReserved   //nolint:stylecheck,golint
	Balances_Unreserved         []gsrpcTypes.EventBalancesUnreserved //nolint:stylecheck,golint
	Balances_ReserveRepatriated []gsrpcTypes.EventBalancesReserveRepatriated
	Balances_Deposit            []gsrpcTypes.EventBalancesDeposit //nolint:stylecheck,golint
	Balances_Withdraw           []gsrpcTypes.EventBalancesWithdraw
	Balances_Slashed            []gsrpcTypes.EventBalancesSlashed
	//Balances_ReservedRepatriated                    []gsrpcTypes.EventBalancesReserveRepatriated             //nolint:stylecheck,golint
	Grandpa_NewAuthorities             []gsrpcTypes.EventGrandpaNewAuthorities             //nolint:stylecheck,golint
	Grandpa_Paused                     []gsrpcTypes.EventGrandpaPaused                     //nolint:stylecheck,golint
	Grandpa_Resumed                    []gsrpcTypes.EventGrandpaResumed                    //nolint:stylecheck,golint
	ImOnline_HeartbeatReceived         []gsrpcTypes.EventImOnlineHeartbeatReceived         //nolint:stylecheck,golint
	ImOnline_AllGood                   []gsrpcTypes.EventImOnlineAllGood                   //nolint:stylecheck,golint
	ImOnline_SomeOffline               []gsrpcTypes.EventImOnlineSomeOffline               //nolint:stylecheck,golint
	Indices_IndexAssigned              []gsrpcTypes.EventIndicesIndexAssigned              //nolint:stylecheck,golint
	Indices_IndexFreed                 []gsrpcTypes.EventIndicesIndexFreed                 //nolint:stylecheck,golint
	Indices_IndexFrozen                []gsrpcTypes.EventIndicesIndexFrozen                //nolint:stylecheck,golint
	Offences_Offence                   []gsrpcTypes.EventOffencesOffence                   //nolint:stylecheck,golint
	Session_NewSession                 []gsrpcTypes.EventSessionNewSession                 //nolint:stylecheck,golint
	Staking_EraPayout                  []gsrpcTypes.EventStakingEraPayout                  //nolint:stylecheck,golint
	Staking_Reward                     []gsrpcTypes.EventStakingReward                     //nolint:stylecheck,golint
	Staking_Slash                      []gsrpcTypes.EventStakingSlash                      //nolint:stylecheck,golint
	Staking_OldSlashingReportDiscarded []gsrpcTypes.EventStakingOldSlashingReportDiscarded //nolint:stylecheck,golint
	Staking_StakingElection            []gsrpcTypes.EventStakingStakingElection            //nolint:stylecheck,golint
	Staking_Bonded                     []gsrpcTypes.EventStakingBonded                     //nolint:stylecheck,golint
	Staking_Unbonded                   []gsrpcTypes.EventStakingUnbonded                   //nolint:stylecheck,golint
	Staking_Withdrawn                  []gsrpcTypes.EventStakingWithdrawn                  //nolint:stylecheck,golint
	//Staking_Kicked                     []types2.EventStakingKicked                         //nolint:stylecheck,golint
	System_ExtrinsicSuccess []gsrpcTypes.EventSystemExtrinsicSuccess //nolint:stylecheck,golint
	//System_ExtrinsicFailed                          []gsrpcTypes.EventSystemExtrinsicFailed                  //nolint:stylecheck,golint
	System_CodeUpdated   []gsrpcTypes.EventSystemCodeUpdated   //nolint:stylecheck,golint
	System_NewAccount    []gsrpcTypes.EventSystemNewAccount    //nolint:stylecheck,golint
	System_KilledAccount []gsrpcTypes.EventSystemKilledAccount //nolint:stylecheck,golint
	//Assets_Created                                  []types2.EventAssetCreated                           //nolint:stylecheck,golint
	Assets_Issued      []gsrpcTypes.EventAssetIssued      //nolint:stylecheck,golint
	Assets_Transferred []gsrpcTypes.EventAssetTransferred //nolint:stylecheck,golint
	//Assets_Burned                                   []types2.EventAssetBurned                            //nolint:stylecheck,golint
	//Assets_TeamChanged                              []types2.EventAssetTeamChanged                       //nolint:stylecheck,golint
	//Assets_OwnerChanged                             []types2.EventAssetOwnerChanged                      //nolint:stylecheck,golint
	//Assets_Frozen                                   []types2.EventAssetFrozen                            //nolint:stylecheck,golint
	//Assets_Thawed                                   []types2.EventAssetThawed                            //nolint:stylecheck,golint
	//Assets_AssetFrozen                              []types2.EventAssetAssetFrozen                       //nolint:stylecheck,golint
	//Assets_AssetThawed                              []types2.EventAssetAssetThawed                       //nolint:stylecheck,golint
	Assets_Destroyed []gsrpcTypes.EventAssetDestroyed //nolint:stylecheck,golint
	//Assets_ForceCreated                             []types2.EventAssetForceCreated                      //nolint:stylecheck,golint
	//Assets_MetadataSet                              []types2.EventAssetMetadataSet                       //nolint:stylecheck,golint
	//Assets_MetadataCleared                          []types2.EventAssetMetadataCleared                   //nolint:stylecheck,golint
	//Assets_ApprovedTransfer                         []types2.EventAssetApprovedTransfer                  //nolint:stylecheck,golint
	//Assets_ApprovalCancelled                        []types2.EventAssetApprovalCancelled                 //nolint:stylecheck,golint
	//Assets_TransferredApproved                      []types2.EventAssetTransferredApproved               //nolint:stylecheck,golint
	//Assets_AssetStatusChanged                       []types2.EventAssetAssetStatusChanged                //nolint:stylecheck,golint
	Democracy_Proposed                 []gsrpcTypes.EventDemocracyProposed                 //nolint:stylecheck,golint
	Democracy_Tabled                   []gsrpcTypes.EventDemocracyTabled                   //nolint:stylecheck,golint
	Democracy_ExternalTabled           []gsrpcTypes.EventDemocracyExternalTabled           //nolint:stylecheck,golint
	Democracy_Started                  []gsrpcTypes.EventDemocracyStarted                  //nolint:stylecheck,golint
	Democracy_Passed                   []gsrpcTypes.EventDemocracyPassed                   //nolint:stylecheck,golint
	Democracy_NotPassed                []gsrpcTypes.EventDemocracyNotPassed                //nolint:stylecheck,golint
	Democracy_Cancelled                []gsrpcTypes.EventDemocracyCancelled                //nolint:stylecheck,golint
	Democracy_Executed                 []gsrpcTypes.EventDemocracyExecuted                 //nolint:stylecheck,golint
	Democracy_Delegated                []gsrpcTypes.EventDemocracyDelegated                //nolint:stylecheck,golint
	Democracy_Undelegated              []gsrpcTypes.EventDemocracyUndelegated              //nolint:stylecheck,golint
	Democracy_Vetoed                   []gsrpcTypes.EventDemocracyVetoed                   //nolint:stylecheck,golint
	Democracy_PreimageNoted            []gsrpcTypes.EventDemocracyPreimageNoted            //nolint:stylecheck,golint
	Democracy_PreimageUsed             []gsrpcTypes.EventDemocracyPreimageUsed             //nolint:stylecheck,golint
	Democracy_PreimageInvalid          []gsrpcTypes.EventDemocracyPreimageInvalid          //nolint:stylecheck,golint
	Democracy_PreimageMissing          []gsrpcTypes.EventDemocracyPreimageMissing          //nolint:stylecheck,golint
	Democracy_PreimageReaped           []gsrpcTypes.EventDemocracyPreimageReaped           //nolint:stylecheck,golint
	Democracy_Unlocked                 []gsrpcTypes.EventDemocracyUnlocked                 //nolint:stylecheck,golint
	Democracy_Blacklisted              []gsrpcTypes.EventDemocracyBlacklisted              //nolint:stylecheck,golint
	Council_Proposed                   []gsrpcTypes.EventCollectiveProposed                //nolint:stylecheck,golint
	Council_Voted                      []gsrpcTypes.EventCollectiveVoted                   //nolint:stylecheck,golint
	Council_Approved                   []gsrpcTypes.EventCollectiveApproved                //nolint:stylecheck,golint
	Council_Disapproved                []gsrpcTypes.EventCollectiveDisapproved             //nolint:stylecheck,golint
	Council_Executed                   []gsrpcTypes.EventCollectiveExecuted                //nolint:stylecheck,golint
	Council_MemberExecuted             []gsrpcTypes.EventCollectiveMemberExecuted          //nolint:stylecheck,golint
	Council_Closed                     []gsrpcTypes.EventCollectiveClosed                  //nolint:stylecheck,golint
	TechnicalCommittee_Proposed        []gsrpcTypes.EventTechnicalCommitteeProposed        //nolint:stylecheck,golint
	TechnicalCommittee_Voted           []gsrpcTypes.EventTechnicalCommitteeVoted           //nolint:stylecheck,golint
	TechnicalCommittee_Approved        []gsrpcTypes.EventTechnicalCommitteeApproved        //nolint:stylecheck,golint
	TechnicalCommittee_Disapproved     []gsrpcTypes.EventTechnicalCommitteeDisapproved     //nolint:stylecheck,golint
	TechnicalCommittee_Executed        []gsrpcTypes.EventTechnicalCommitteeExecuted        //nolint:stylecheck,golint
	TechnicalCommittee_MemberExecuted  []gsrpcTypes.EventTechnicalCommitteeMemberExecuted  //nolint:stylecheck,golint
	TechnicalCommittee_Closed          []gsrpcTypes.EventTechnicalCommitteeClosed          //nolint:stylecheck,golint
	TechnicalMembership_MemberAdded    []gsrpcTypes.EventTechnicalMembershipMemberAdded    //nolint:stylecheck,golint
	TechnicalMembership_MemberRemoved  []gsrpcTypes.EventTechnicalMembershipMemberRemoved  //nolint:stylecheck,golint
	TechnicalMembership_MembersSwapped []gsrpcTypes.EventTechnicalMembershipMembersSwapped //nolint:stylecheck,golint
	TechnicalMembership_MembersReset   []gsrpcTypes.EventTechnicalMembershipMembersReset   //nolint:stylecheck,golint
	TechnicalMembership_KeyChanged     []gsrpcTypes.EventTechnicalMembershipKeyChanged     //nolint:stylecheck,golint
	TechnicalMembership_Dummy          []gsrpcTypes.EventTechnicalMembershipDummy          //nolint:stylecheck,golint
	//ElectionProviderMultiPhase_SolutionStored       []types2.EventElectionMultiPhaseSolutionStored       //nolint:stylecheck,golint
	//ElectionProviderMultiPhase_ElectionFinalized    []types2.EventElectionMultiPhaseElectionFinalized    //nolint:stylecheck,golint
	//ElectionProviderMultiPhase_Rewarded             []types2.EventElectionMultiPhaseRewarded             //nolint:stylecheck,golint
	//ElectionProviderMultiPhase_Slashed              []types2.EventElectionMultiPhaseSlashed              //nolint:stylecheck,golint
	//ElectionProviderMultiPhase_SignedPhaseStarted   []types2.EventElectionMultiPhaseSignedPhaseStarted   //nolint:stylecheck,golint
	//ElectionProviderMultiPhase_UnsignedPhaseStarted []types2.EventElectionMultiPhaseUnsignedPhaseStarted //nolint:stylecheck,golint
	Elections_NewTerm       []gsrpcTypes.EventElectionsNewTerm       //nolint:stylecheck,golint
	Elections_EmptyTerm     []gsrpcTypes.EventElectionsEmptyTerm     //nolint:stylecheck,golint
	Elections_ElectionError []gsrpcTypes.EventElectionsElectionError //nolint:stylecheck,golint
	Elections_MemberKicked  []gsrpcTypes.EventElectionsMemberKicked  //nolint:stylecheck,golint
	//Elections_Renounced                             []types2.EventElectionsRenounced                     //nolint:stylecheck,golint
	//Elections_CandidateSlashed                      []types2.EventElectionsCandidateSlashed              //nolint:stylecheck,golint
	//Elections_SeatHolderSlashed                     []types2.EventElectionsSeatHolderSlashed             //nolint:stylecheck,golint
	Identity_IdentitySet             []gsrpcTypes.EventIdentitySet                     //nolint:stylecheck,golint
	Identity_IdentityCleared         []gsrpcTypes.EventIdentityCleared                 //nolint:stylecheck,golint
	Identity_IdentityKilled          []gsrpcTypes.EventIdentityKilled                  //nolint:stylecheck,golint
	Identity_JudgementRequested      []gsrpcTypes.EventIdentityJudgementRequested      //nolint:stylecheck,golint
	Identity_JudgementUnrequested    []gsrpcTypes.EventIdentityJudgementUnrequested    //nolint:stylecheck,golint
	Identity_JudgementGiven          []gsrpcTypes.EventIdentityJudgementGiven          //nolint:stylecheck,golint
	Identity_RegistrarAdded          []gsrpcTypes.EventIdentityRegistrarAdded          //nolint:stylecheck,golint
	Identity_SubIdentityAdded        []gsrpcTypes.EventIdentitySubIdentityAdded        //nolint:stylecheck,golint
	Identity_SubIdentityRemoved      []gsrpcTypes.EventIdentitySubIdentityRemoved      //nolint:stylecheck,golint
	Identity_SubIdentityRevoked      []gsrpcTypes.EventIdentitySubIdentityRevoked      //nolint:stylecheck,golint
	Society_Founded                  []gsrpcTypes.EventSocietyFounded                  //nolint:stylecheck,golint
	Society_Bid                      []gsrpcTypes.EventSocietyBid                      //nolint:stylecheck,golint
	Society_Vouch                    []gsrpcTypes.EventSocietyVouch                    //nolint:stylecheck,golint
	Society_AutoUnbid                []gsrpcTypes.EventSocietyAutoUnbid                //nolint:stylecheck,golint
	Society_Unbid                    []gsrpcTypes.EventSocietyUnbid                    //nolint:stylecheck,golint
	Society_Unvouch                  []gsrpcTypes.EventSocietyUnvouch                  //nolint:stylecheck,golint
	Society_Inducted                 []gsrpcTypes.EventSocietyInducted                 //nolint:stylecheck,golint
	Society_SuspendedMemberJudgement []gsrpcTypes.EventSocietySuspendedMemberJudgement //nolint:stylecheck,golint
	Society_CandidateSuspended       []gsrpcTypes.EventSocietyCandidateSuspended       //nolint:stylecheck,golint
	Society_MemberSuspended          []gsrpcTypes.EventSocietyMemberSuspended          //nolint:stylecheck,golint
	Society_Challenged               []gsrpcTypes.EventSocietyChallenged               //nolint:stylecheck,golint
	Society_Vote                     []gsrpcTypes.EventSocietyVote                     //nolint:stylecheck,golint
	Society_DefenderVote             []gsrpcTypes.EventSocietyDefenderVote             //nolint:stylecheck,golint
	Society_NewMaxMembers            []gsrpcTypes.EventSocietyNewMaxMembers            //nolint:stylecheck,golint
	Society_Unfounded                []gsrpcTypes.EventSocietyUnfounded                //nolint:stylecheck,golint
	Society_Deposit                  []gsrpcTypes.EventSocietyDeposit                  //nolint:stylecheck,golint
	Recovery_RecoveryCreated         []gsrpcTypes.EventRecoveryCreated                 //nolint:stylecheck,golint
	Recovery_RecoveryInitiated       []gsrpcTypes.EventRecoveryInitiated               //nolint:stylecheck,golint
	Recovery_RecoveryVouched         []gsrpcTypes.EventRecoveryVouched                 //nolint:stylecheck,golint
	Recovery_RecoveryClosed          []gsrpcTypes.EventRecoveryClosed                  //nolint:stylecheck,golint
	Recovery_AccountRecovered        []gsrpcTypes.EventRecoveryAccountRecovered        //nolint:stylecheck,golint
	Recovery_RecoveryRemoved         []gsrpcTypes.EventRecoveryRemoved                 //nolint:stylecheck,golint
	Vesting_VestingUpdated           []gsrpcTypes.EventVestingVestingUpdated           //nolint:stylecheck,golint
	Vesting_VestingCompleted         []gsrpcTypes.EventVestingVestingCompleted         //nolint:stylecheck,golint
	// gsrpcTypes.BlockNum is incorrect
	//Scheduler_Scheduled                             []gsrpcTypes.EventSchedulerScheduled                     //nolint:stylecheck,golint
	//Scheduler_Canceled                              []gsrpcTypes.EventSchedulerCanceled                      //nolint:stylecheck,golint
	//Scheduler_Dispatched                            []gsrpcTypes.EventSchedulerDispatched                                    //nolint:stylecheck,golint

	//Proxy_ProxyExecuted         []gsrpcTypes.EventProxyProxyExecuted       //nolint:stylecheck,golint
	Proxy_AnonymousCreated []gsrpcTypes.EventProxyAnonymousCreated //nolint:stylecheck,golint
	Proxy_Announced        []gsrpcTypes.EventProxyAnnounced        //nolint:stylecheck,golint
	Sudo_Sudid             []gsrpcTypes.EventSudoSudid             //nolint:stylecheck,golint
	Sudo_KeyChanged        []gsrpcTypes.EventSudoKeyChanged        //nolint:stylecheck,golint
	Sudo_SudoAsDone        []gsrpcTypes.EventSudoAsDone            //nolint:stylecheck,golint
	Treasury_Proposed      []gsrpcTypes.EventTreasuryProposed      //nolint:stylecheck,golint
	Treasury_Spending      []gsrpcTypes.EventTreasurySpending      //nolint:stylecheck,golint
	Treasury_Awarded       []gsrpcTypes.EventTreasuryAwarded       //nolint:stylecheck,golint
	Treasury_Rejected      []gsrpcTypes.EventTreasuryRejected      //nolint:stylecheck,golint
	Treasury_Burnt         []gsrpcTypes.EventTreasuryBurnt         //nolint:stylecheck,golint
	Treasury_Rollover      []gsrpcTypes.EventTreasuryRollover      //nolint:stylecheck,golint
	Treasury_Deposit       []gsrpcTypes.EventTreasuryDeposit       //nolint:stylecheck,golint
	//Tips_NewTip                 []types2.EventTipsNewTip                   //nolint:stylecheck,golint
	//Tips_TipClosing             []types2.EventTipsTipClosing               //nolint:stylecheck,golint
	//Tips_TipClosed              []types2.EventTipsTipClosed                //nolint:stylecheck,golint
	//Tips_TipRetracted           []types2.EventTipsTipRetracted             //nolint:stylecheck,golint
	//Tips_TipSlashed             []types2.EventTipsTipSlashed               //nolint:stylecheck,golint
	//Bounties_BountyProposed     []types2.EventBountyBountyProposed         //nolint:stylecheck,golint
	//Bounties_BountyRejected     []types2.EventBountyBountyRejected         //nolint:stylecheck,golint
	//Bounties_BountyBecameActive []types2.EventBountyBountyBecameActive     //nolint:stylecheck,golint
	//Bounties_BountyAwarded      []types2.EventBountyBountyAwarded          //nolint:stylecheck,golint
	//Bounties_BountyClaimed      []types2.EventBountyBountyClaimed          //nolint:stylecheck,golint
	//Bounties_BountyCanceled     []types2.EventBountyBountyCanceled         //nolint:stylecheck,golint
	//Bounties_BountyExtended     []types2.EventBountyBountyExtended         //nolint:stylecheck,golint
	Contracts_Instantiated []gsrpcTypes.EventContractsInstantiated //nolint:stylecheck,golint
	Contracts_Evicted      []gsrpcTypes.EventContractsEvicted      //nolint:stylecheck,golint
	//Contracts_Terminated        []types2.EventContractsTerminated          //nolint:stylecheck,golint
	Contracts_Restored        []gsrpcTypes.EventContractsRestored        //nolint:stylecheck,golint
	Contracts_CodeStored      []gsrpcTypes.EventContractsCodeStored      //nolint:stylecheck,golint
	Contracts_ScheduleUpdated []gsrpcTypes.EventContractsScheduleUpdated //nolint:stylecheck,golint
	//Contracts_ContractEmitted   []types2.EventContractsContractEmitted     //nolint:stylecheck,golint
	//Contracts_CodeRemoved       []types2.EventContractsCodeRemoved         //nolint:stylecheck,golint
	Utility_BatchInterrupted   []gsrpcTypes.EventUtilityBatchInterrupted //nolint:stylecheck,golint
	Utility_BatchCompleted     []gsrpcTypes.EventUtilityBatchCompleted   //nolint:stylecheck,golint
	Multisig_NewMultisig       []gsrpcTypes.EventMultisigNewMultisig     //nolint:stylecheck,golint
	Multisig_MultisigApproval  []gsrpcTypes.EventMultisigApproval        //nolint:stylecheck,golint
	Multisig_MultisigExecuted  []gsrpcTypes.EventMultisigExecuted        //nolint:stylecheck,golint
	Multisig_MultisigCancelled []gsrpcTypes.EventMultisigCancelled
}
