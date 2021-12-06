package fsm

// StateMachine implementation of state machine
type StateMachine struct {
	// data a alice of data
	data []Type

	// commit is a copy of the latest commit
	commit commit

	// commits of machine
	history []commit
}

type commit struct {
	id   string
	data []Type
}

type Operation struct {
	Name  OperationName
	Path  []string
	Value interface{}
}

type machinery interface {
	New(operation Operation) error
	Remove(operation Operation) error
	Set(operation Operation) error
	Get(operation Operation) (Type, error)
	Snapshot() commit

	Commit() error
	Rollback(id string) error

	orderToOperation(ops string) (Operation, error)
}
