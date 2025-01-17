package types

type GetAllocationResponse struct {
	Status string `json:"status,omitempty"`
	Data   struct {
		LikelySybil                        bool    `json:"likely_sybil,omitempty"`
		ExcludedBecauseDistinctWeeksUnder3 bool    `json:"excluded_because_distinct_weeks_under_3,omitempty"`
		ExcludedBecauseObscurePairs        bool    `json:"excluded_because_obscure_pairs,omitempty"`
		ExcludedBecauseFailureRate         bool    `json:"excluded_because_failure_rate,omitempty"`
		ExcludedBecauseNotEnoughSwapPoints bool    `json:"excluded_because_not_enough_swap_points,omitempty"`
		UserPublicKey                      string  `json:"user_public_key,omitempty"`
		SwapScore                          float64 `json:"swap_score,omitempty"`
		SwapTier                           float64 `json:"swap_tier,omitempty"`
		SwapAllocationBase                 float64 `json:"swap_allocation_base,omitempty"`
		SwapConsistencyBonus               float64 `json:"swap_consistency_bonus,omitempty"`
		ExpertScore                        float64 `json:"expert_score,omitempty"`
		ExpertTier                         float64 `json:"expert_tier,omitempty"`
		ExpertAllocation                   float64 `json:"expert_allocation,omitempty"`
		StakersScore                       float64 `json:"stakers_score,omitempty"`
		StakersAllocationBase              float64 `json:"stakers_allocation_base,omitempty"`
		StakersSuperVoterBonus             float64 `json:"stakers_super_voter_bonus,omitempty"`
		StakersSuperStakerBonus            float64 `json:"stakers_super_staker_bonus,omitempty"`
		MobilePotentialBonus               float64 `json:"mobile_potential_bonus,omitempty"`
		FlaggedByChainanalysis             bool    `json:"flagged_by_chainanalysis,omitempty"`
		FlaggedAsATA                       bool    `json:"flagged_as_ATA,omitempty"`
		TotalAllocated                     float64 `json:"total_allocated,omitempty"`
	} `json:"data,omitempty"`
}
