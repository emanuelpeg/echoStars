package concurrencyType

type ConcurrencyType uint

const (
	Sequential ConcurrencyType = iota
	ConcurrentWG
	ConcurrentChan
)
