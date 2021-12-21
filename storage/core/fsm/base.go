package fsm

import (
	"encoding/json"
	"github.com/google/uuid"
	"strings"
)

func NewStateMachine() *StateMachine {
	sm := StateMachine{
		data:    TypeMap{},
		commit:  commit{},
		history: make([]commit, 0),
	}

	return &sm
}

// Operation do an operation
// parse operation first, then make data do the change
// format: operation type path value
func (sm *StateMachine) Operation(order string) error {
	// parse operation
	operation, err := strToOp(order)
	if err != nil {
		return err
	}
	switch operation.Name {
	case New:
		return sm.data.new(operation.Path, operation.Value)
	case Remove:
		return sm.data.remove(operation.Path)
	case Set:
		return sm.data.set(operation.Path, operation.Value)
	default:
		return ErrOperationUndefined
	}
}

// Get the data by key path
func (sm *StateMachine) Get(path []string) interface{} {
	data := sm.commit.data
	for index, key := range path {
		for _, kv := range data.Value {
			if kv.GetKey() == key {
				if index == len(path)-1 {
					return kv.GetValue()
				}
				if kv.GetTypeName() != Map {
					return nil
				}
				data = kv.This().(TypeMap)
				continue
			}
		}
		return nil
	}
	return nil
}

// Commit to commit a new change, using uuid to
// guarantee the unique id
func (sm *StateMachine) Commit() string {
	uid := uuid.New().String()
	newCommit := commit{
		id:   uid,
		data: TypeMap{sm.data.Key, sm.data.Value},
	}
	sm.history = append(sm.history, newCommit)
	sm.commit = newCommit
	return uid
}

// Rollback to a history commit
func (sm *StateMachine) Rollback(id string) error {
	for index, commit := range sm.history {
		if commit.id == id {
			sm.commit = commit
			sm.data = commit.data
			sm.history = sm.history[:index+1]
			return nil
		}
	}
	return ErrCommitNotFound
}

func (t *TypeMap) new(path []string, value Type) error {
	if len(path) == 0 {
		return ErrWrongArgs
	}
	// at the last point of path
	if len(path) == 1 {
		// check for conflict
		for _, item := range t.Value {
			if item.GetKey() == value.GetKey() {
				return ErrItemExisted
			}
		}
		t.Value = append(t.Value, value)
		return nil
	}

	for index, item := range t.Value {
		if item.GetKey() == path[0] {
			if item.GetTypeName() != Map {
				return ErrPathNotMatch
			}
			itemMap, ok := item.This().(TypeMap)
			if !ok {
				return ErrWrongArgs
			}
			if err := itemMap.new(path[1:], value); err != nil {
				return err
			}
			t.Value[index] = &itemMap
			return nil
		}
	}
	return ErrPathNotMatch
}

func (t *TypeMap) remove(path []string) error {
	if len(path) == 0 {
		return ErrWrongArgs
	}
	// at the last point of path
	if len(path) == 1 {
		// find end point
		for index, item := range t.Value {
			if item.GetKey() == path[0] {
				items := make([]Type, 0)
				if index > 0 {
					items = append(items, t.Value[:index]...)
				}
				if index < len(t.Value)-1 {
					items = append(items, t.Value[index+1:]...)
				}
				t.Value = items
				return nil
			}
		}
		return ErrItemNotExisted
	}

	for index, item := range t.Value {
		if item.GetKey() == path[0] {
			if item.GetTypeName() != Map {
				return ErrPathNotMatch
			}
			itemMap, ok := item.This().(TypeMap)
			if !ok {
				return ErrWrongArgs
			}
			if err := itemMap.remove(path[1:]); err != nil {
				return err
			}
			t.Value[index] = &itemMap
			return nil
		}
	}
	return ErrPathNotMatch
}

func (t *TypeMap) set(path []string, value Type) error {
	if len(path) == 0 {
		return ErrWrongArgs
	}
	// at the last point of path
	if len(path) == 1 {
		// find end point
		for index, item := range t.Value {
			if item.GetKey() == path[0] {
				t.Value[index] = value
				return nil
			}
		}
		return ErrItemNotExisted
	}
	for index, item := range t.Value {
		if item.GetKey() == path[0] {
			if item.GetTypeName() != Map {
				return ErrPathNotMatch
			}
			itemMap, ok := item.This().(TypeMap)
			if !ok {
				return ErrWrongArgs
			}
			if err := itemMap.set(path[1:], value); err != nil {
				return err
			}
			t.Value[index] = &itemMap
			return nil
		}
	}
	return ErrPathNotMatch
}

func strToOp(order string) (Operation, error) {
	keys := strings.Split(order, " ")
	l := len(keys)

	if l <= 3 || l >= 5 {
		return Operation{}, ErrWrongArgs
	}

	op := operationName(keys[0])
	tp := typeName(keys[1])

	if tp == UndefinedType || op == UndefinedOp {
		return Operation{}, ErrWrongArgs
	}

	path := strings.Split(keys[2], ".")
	pathL := len(path)
	typeMap := TypeMap{Key: path[pathL-1], Value: make([]Type, 0)}
	typeArray := TypeArray{Key: path[pathL-1], Value: make([]string, 0)}
	typeString := TypeString{Key: path[pathL-1], Value: ""}
	// get value
	switch tp {
	case Map:
		if l >= 4 {
			return Operation{}, ErrWrongArgs
		}
		return Operation{Name: op, Path: path, Value: &typeMap}, nil
	case Array:
		if (op == Remove && l >= 4) ||
			((op == New || op == Set) && l < 4) {
			return Operation{}, ErrWrongArgs
		}
		if l == 4 {
			v := keys[3]
			var data []string
			if json.Unmarshal([]byte(v), &data) != nil {
				return Operation{}, ErrWrongArgs
			}
			typeArray.Value = data
		}
		return Operation{Name: op, Path: path, Value: &typeArray}, nil
	case String:
		if (op == Remove && l >= 4) ||
			((op == New || op == Set) && l < 4) {
			return Operation{}, ErrWrongArgs
		}
		if l == 4 {
			typeString.Value = keys[3]
		}
		return Operation{Name: op, Path: path, Value: &typeString}, nil
	}
	return Operation{}, ErrWrongArgs
}
