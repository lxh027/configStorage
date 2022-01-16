package scheduler

import "errors"

type RegisterError error

var GetDataErr RegisterError = errors.New("get data error")
var SetDataErr RegisterError = errors.New("set data error")
var DelDataErr RegisterError = errors.New("delete data error")
var RaftFullErr RegisterError = errors.New("raft cluster already full")
var RaftEmptyErr RegisterError = errors.New("raft cluster is empty")