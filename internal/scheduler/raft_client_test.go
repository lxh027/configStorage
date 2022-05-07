package scheduler

import (
	"configStorage/api/raftrpc"
	"configStorage/internal/raft"
	"configStorage/tools/random"
	"context"
	"google.golang.org/grpc"
	"log"
	"sync"
	"testing"
	time2 "time"
)

var cfg = raft.ClientConfig{
	Size: 3,
	Addresses: []string{
		"172.17.0.1:3001",
		"172.17.0.1:3002",
		"172.17.0.1:3003",
	},
}
var instances = make([]raftrpc.StateClient, cfg.Size)

func Init() {
	for i, address := range cfg.Addresses {
		cOpts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		conn, err := grpc.Dial(address, cOpts...)
		if err != nil {
			log.Panicf("error while get conn for raft server, error: %v", err.Error())
		}
		instances[i] = raftrpc.NewStateClient(conn)
	}
}

func TestStopServer(t *testing.T) {
	Init()
	instances[1].StopServer(context.Background(), &raftrpc.ControllerMsg{})
}

func TestRestartServer(t *testing.T) {
	Init()
	instances[1].StartServer(context.Background(), &raftrpc.ControllerMsg{})
}

func TestMultiSetOps(t *testing.T) {
	client := NewSchedulerClient("172.17.0.1:2888")

	const PNUM = 100
	const NAMESPACE = "name"
	const KEY = "UbGcExAzeAkAxvuE"
	failNum, successNum := 0, 0
	var timeout time2.Duration = 0
	wg := sync.WaitGroup{}
	wg.Add(PNUM)
	for i := 0; i < PNUM; i++ {
		go func() {
			op := Log{Key: random.RandString(5), Value: random.RandString(5)}
			time := time2.Now()
			if _, err := client.Commit(NAMESPACE, KEY, []Log{op}); err != nil {
				failNum++
			} else {
				timeout += time2.Now().Sub(time)
				successNum++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	var successRate float64 = float64(successNum) / float64(PNUM) * 100
	var avgTimeout float64 = float64(timeout.Milliseconds()) / float64(successNum)
	log.Printf(`successRate = %f%s\n ToTalNum = %d\n SuccessNum = %d\n`, successRate, "%", PNUM, successNum)
	log.Printf("avg timeout = %f", avgTimeout)
}
