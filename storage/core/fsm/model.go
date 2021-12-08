package fsm

// StateMachine implementation of state machine
type StateMachine struct {
	// data a alice of data
	data TypeMap

	// commit is a copy of the latest commit
	commit commit

	// commits of machine
	history []commit
}

type commit struct {
	id   string
	data TypeMap
}

type Operation struct {
	Name  OperationName
	Path  []string
	Value Type
}

type machinery interface {
	Operation(order string) error
	Get(path []string) interface{}

	// TODO Snapshot
	// Snapshot() commit

	Commit() string
	Rollback(id string) error
}
