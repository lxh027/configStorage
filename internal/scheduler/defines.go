package scheduler

import "errors"

type RegisterError error
type KvError error

var GetDataErr RegisterError = errors.New("get data error")
var SetDataErr RegisterError = errors.New("set data error")
var DelDataErr RegisterError = errors.New("delete data error")
var RaftFullErr RegisterError = errors.New("raft cluster already full")
var RaftInstanceExistedErr RegisterError = errors.New("instance uid already existed in the cluster")
var RaftInstanceNotExistedErr RegisterError = errors.New("instance uid is not existed in the cluster")
var RaftEmptyErr RegisterError = errors.New("raft cluster is empty")

var NamespaceExistedErr KvError = errors.New("namespace existed")
var NamespaceNotExistedErr KvError = errors.New("namespace not existed")
var PrivateKeyUnPatchErr KvError = errors.New("private key not patch")

var NamespaceCommittingErr KvError = errors.New("namespace is committing")

var ClusterNotExistedErr RegisterError = errors.New("cluster not existed")

type clusterStatus int

const (
	Unready clusterStatus = iota
	Ready
	Changed
	Renew
)
