package kubestatus

import "sync/atomic"

var (
	Healthy int32 = 0
	Ready   int32 = 0
)

func SetHealthy() {
	atomic.StoreInt32(&Healthy, 1)
}

func SetUnHealthy() {
	atomic.StoreInt32(&Healthy, 0)
}

func SetReady() {
	atomic.StoreInt32(&Ready, 1)
}

func SetUnReady() {
	atomic.StoreInt32(&Ready, 0)
}
