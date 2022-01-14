package scheduler

import "errors"

type RegisterError error

var GetDataErr RegisterError = errors.New("get data error")
var SetDataErr RegisterError = errors.New("set data error")
var RaftFullErr RegisterError = errors.New("raft cluster already full")
