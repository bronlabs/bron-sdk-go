package types

type IntentsQuoteEmbedded struct {
	SlippageBps *string `json:"slippageBps,omitempty"`
	SlippageReason *SlippageReason `json:"slippageReason,omitempty"`
	StockSession *StockSession `json:"stockSession,omitempty"`
}
