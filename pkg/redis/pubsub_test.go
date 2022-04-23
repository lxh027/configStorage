package redis

import (
	"configStorage/pkg/config"
	"fmt"
	r "github.com/garyburd/redigo/redis"
	"gopkg.in/yaml.v2"
	"testing"
	"time"
)

const cfg = "env: dev\nhost: 127.0.0.1\nport: 6379\nauth:\ncon_type: tcp\n\nmax_idle: 4\nmax_active: 8\n\ntimeout: 300"

func TestPubSub(t *testing.T) {
	var redis config.Redis
	err := yaml.Unmarshal([]byte(cfg), &redis)
	if err != nil {
		panic(fmt.Sprintf("err: %v", err.Error()))
	}
	c1, _ := NewRedisClient(&redis)

	go func() {
		err := c1.Subscribe("testKey", func(msg r.Message) {
		})
		if err != nil {
			fmt.Printf("err: %v", err.Error())
		}
	}()

	c2, _ := NewRedisClient(&redis)
	_ = c2.Publish("testKey", "123")
	_ = c2.Publish("testKey", 123)
	_ = c2.Publish("testKey", []string{"123", "456"})
	_ = c2.Publish("testKey", struct {
		A string
		B int
	}{"123", 123})
	fmt.Printf("%v", err)
	time.Sleep(5 * time.Second)
}
