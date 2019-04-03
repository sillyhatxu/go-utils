package goredis

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	InitialRedisConfig("127.0.0.1:6379", "", 0)
	redisClient, err := RedisConf.GetClient()
	assert.Nil(t, err)
	assert.NotNil(t, redisClient)
}

func TestGetAndSet(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	InitialRedisConfig("127.0.0.1:6379", "", 0)
	RedisConf.Set("test_01", "P_1001")
	RedisConf.Set("test_02", "P_1002")
	RedisConf.Set("test_03", "P_1003")
	RedisConf.Set("test_04", "P_1004")
	RedisConf.Set("test_05", "P_1005")
	v1, err := RedisConf.Get("test_01")
	assert.Nil(t, err)
	assert.EqualValues(t, v1, "P_1001")
	v2, err := RedisConf.Get("test_02")
	assert.Nil(t, err)
	assert.EqualValues(t, v2, "P_1002")
	v3, err := RedisConf.Get("test_03")
	assert.Nil(t, err)
	assert.EqualValues(t, v3, "P_1003")
	v4, err := RedisConf.Get("test_04")
	assert.Nil(t, err)
	assert.EqualValues(t, v4, "P_1004")
	v5, err := RedisConf.Get("test_05")
	assert.Nil(t, err)
	assert.EqualValues(t, v5, "P_1005")
	exist1, err := RedisConf.Exists("test_01")
	assert.Nil(t, err)
	assert.EqualValues(t, exist1, true)
	exist2, err := RedisConf.Exists("test_02")
	assert.Nil(t, err)
	assert.EqualValues(t, exist2, true)
	exist3, err := RedisConf.Exists("test_03")
	assert.Nil(t, err)
	assert.EqualValues(t, exist3, true)
	exist4, err := RedisConf.Exists("test_04")
	assert.Nil(t, err)
	assert.EqualValues(t, exist4, true)
	exist5, err := RedisConf.Exists("test_05")
	assert.Nil(t, err)
	assert.EqualValues(t, exist5, true)
}