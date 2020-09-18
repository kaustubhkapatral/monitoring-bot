package types

import "time"

type Validators struct {
	Result []struct {
		ID          string `json:"id"`
		VotingPower int64  `json:"voting_power"`
	} `json:"result"`
}

type Block_response struct {
	Result struct {
		Height  int `json:"height"`
		Round   int `json:"round"`
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Total int    `json:"total"`
				Hash  string `json:"hash"`
			} `json:"parts"`
		} `json:"block_id"`
		Signatures []struct {
			BlockIDFlag      int       `json:"block_id_flag"`
			ValidatorAddress string    `json:"validator_address"`
			Timestamp        time.Time `json:"timestamp"`
			Signature        string    `json:"signature"`
		} `json:"signatures"`
	} `json:"result"`
}
