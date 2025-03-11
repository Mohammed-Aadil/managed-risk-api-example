package enum

type RiskState int

const (
	Open RiskState = iota
	Closed
	Accepted
	Investigating
)
