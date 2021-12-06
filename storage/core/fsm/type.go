package fsm

type TypeName int
type OperationName int

const (
	String TypeName = iota
	Map
	Array
)

const (
	New OperationName = iota
	Remove
	Set
	Get
)

type Type interface {
	GetTypeName() TypeName
}

type TypeString struct {
	Key   string
	Value int
}

type TypeMap struct {
	Key   string
	Value []Type
}

type TypeArray struct {
	Data []string
}

func (t *TypeString) GetTypeName() TypeName { return String }
func (t *TypeMap) GetTypeName() TypeName    { return Map }
func (t *TypeArray) GetTypeName() TypeName  { return Array }
