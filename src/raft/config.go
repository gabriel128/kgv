package raft

type Config struct {
	minElectionTimeMs int
	maxElectionTimeMs int
	heartBeatRateMs int
}

var config Config = Config{300, 700, 30}