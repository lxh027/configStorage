package fsm

import (
	"log"
	"testing"
)

func TestFsmBasic(T *testing.T) {
	fsm := NewStateMachine()

	_ = fsm.Operation("new string k v")
	// commit
	commitId := fsm.Commit()
	log.Printf("commit id: %v", commitId)

	path := []string{"k"}
	v := fsm.Get(path)

	log.Printf("get key of path `k`: %v", v)
}

func TestFsmMap(T *testing.T) {
	fsm := NewStateMachine()

	_ = fsm.Operation("new map mp")
	_ = fsm.Operation("new string mp.k v")

	commitId := fsm.Commit()
	log.Printf("commit id: %v", commitId)

	path := []string{"mp", "k"}
	v := fsm.Get(path)

	log.Printf("get key of path `mp.k`: %v", v)

}
