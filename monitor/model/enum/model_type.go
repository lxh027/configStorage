package enum

type ModelType uint8

const (
	InstanceStatus ModelType = iota

	// save log from instances
	Log
)
