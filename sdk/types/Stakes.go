package types

type Stakes struct {
	Embedded *StakesEmbedded `json:"_embedded,omitempty"`
	Stakes *[]Stake `json:"stakes,omitempty"`
}
