package enum

type ModelType uint8

const (
	InstanceStatus ModelType = iota

	// Log save log from instances
	Log
)
