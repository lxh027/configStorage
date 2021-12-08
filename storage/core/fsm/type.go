package fsm

type TypeName string
type OperationName string

const (
	String        TypeName = "string"
	Map           TypeName = "map"
	Array         TypeName = "array"
	UndefinedType TypeName = "undefined"
)

const (
	New         OperationName = "new"
	Remove      OperationName = "rm"
	Set         OperationName = "set"
	UndefinedOp OperationName = "undefined"
)

func operationName(str string) OperationName {
	switch str {
	case "set":
		return Set
	case "rm":
		return Remove
	case "new":
		return New
	default:
		return UndefinedOp
	}
}

func typeName(str string) TypeName {
	switch str {
	case "map":
		return Map
	case "string":
		return String
	case "array":
		return Array
	default:
		return UndefinedType
	}
}

type Type interface {
	GetTypeName() TypeName
	GetKey() string
	GetValue() interface{}
	This() interface{}
}

type TypeString struct {
	Key   string
	Value string
}

type TypeMap struct {
	Key   string
	Value []Type
}

type TypeArray struct {
	Key   string
	Value []string
}

func (t *TypeString) GetTypeName() TypeName { return String }
func (t *TypeString) GetKey() string        { return t.Key }
func (t *TypeString) GetValue() interface{} { return t.Value }
func (t *TypeString) This() interface{}     { return t }
func (t *TypeMap) GetTypeName() TypeName    { return Map }
func (t *TypeMap) GetKey() string           { return t.Key }
func (t *TypeMap) GetValue() interface{}    { return t.Value }
func (t *TypeMap) This() interface{}        { return t }
func (t *TypeArray) GetTypeName() TypeName  { return Array }
func (t *TypeArray) GetKey() string         { return t.Key }
func (t *TypeArray) GetValue() interface{}  { return t.Value }
func (t *TypeArray) This() interface{}      { return t }
