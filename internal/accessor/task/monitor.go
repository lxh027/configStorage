package task

import (
	monitor2 "configStorage/internal/accessor/app/monitor"
	"configStorage/internal/accessor/global"
	"configStorage/internal/raft"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
	"sync"
	"time"
)

func GetMonitorData() {
	var monitors []monitor2.Monitor
	mu := sync.Mutex{}
	go func() {
		for {
			time.Sleep(3 * time.Second)
			mu.Lock()
			if len(monitors) > 5 {
				global.MysqlClient.Create(&monitors)
				monitors = monitors[:0]
			}
			mu.Unlock()
		}
	}()
	for {
		err := global.RedisClient.Subscribe(raft.ReportChan, func(msg redis.Message) {
			var rm raft.ReportMsg
			err := json.Unmarshal(msg.Data, &rm)
			if err != nil {
				log.Printf("parse monitor data error: %v\n", err.Error())
			}
			monitor := monitor2.Monitor{
				RaftID:          rm.RaftID,
				PeerID:          rm.Id,
				IsLeader:        rm.IsLeader,
				Status:          int(rm.Status),
				ConfigVersion:   rm.CfgVersion,
				CurrentTerm:     rm.CurrentTerm,
				CurrentIndex:    rm.CurrentIndex,
				CommitIndex:     rm.CommitIndex,
				MemoryTotal:     rm.MemoryTotal,
				MemoryAvailable: rm.MemoryAvailable,
				MemoryUsed:      rm.MemoryUsed,
				MemoryCur:       rm.MemoryCur,
				Time:            rm.Now,
			}
			mu.Lock()
			monitors = append(monitors, monitor)
			mu.Unlock()
		})
		if err != nil {
			log.Printf("get monitor data error: %v\n", err.Error())
		}
	}
}

func DeleteMonitorData() {
	for {
		global.MysqlClient.
			Where("time < ?", time.Now().Add(-30*time.Minute)).Delete(&monitor2.Monitor{})
		time.Sleep(time.Minute)
	}
}
