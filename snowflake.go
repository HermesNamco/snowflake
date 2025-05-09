package snowflake

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	timeStampMask   uint64 = 1<<41 - 1
	maxWorkerId     uint64 = 1<<5 - 1
	maxDataCenterId uint64 = 1<<5 - 1
	sequenceMask    uint64 = 1<<12 - 1
)

type Snowflake struct {
	// 起始时间戳
	startPoint uint64
	// 序列号位数
	sequenceBits uint64
	// 机器ID
	workerIdBits uint64
	// 数据中心ID位数
	dataCenterIdBits uint64
	// 上次记录时钟
	lastTimeStamp uint64

	mutex sync.Mutex
}

func (s *Snowflake) StartPoint(startPoint uint64) *Snowflake {
	s.startPoint = startPoint
	return s
}

func (s *Snowflake) WorkerIdBits(workIdBits uint64) *Snowflake {
	if workIdBits > maxWorkerId {
		log.Fatal("workIdBits is greater than maxWorkerId")
		return s
	}
	s.workerIdBits = workIdBits
	return s
}

func (s *Snowflake) DataCenterIdBits(dataCentIdBits uint64) *Snowflake {
	if dataCentIdBits > maxDataCenterId {
		log.Fatal("dataCentIdBits is greater than maxDataCenterId")
		return s
	}
	s.dataCenterIdBits = dataCentIdBits
	return s
}

func (s *Snowflake) Next() (uint64, error) {
	currentTime := uint64(time.Now().UnixMilli())
	if currentTime < s.lastTimeStamp {
		return 0, fmt.Errorf("current time is too far in the future")
	}

	if currentTime == s.lastTimeStamp {
		s.mutex.Lock()
		s.sequenceBits = (s.sequenceBits + 1) & sequenceMask
		if s.sequenceBits == 0 {
			currentTime = s.getNextMs(s.lastTimeStamp)
		}
		s.mutex.Unlock()
	}

	currentTime = currentTime & timeStampMask
	// 更新时间戳
	s.lastTimeStamp = currentTime

	// 生成本次id
	var nextId uint64 = (currentTime-s.startPoint)<<22 |
		s.dataCenterIdBits<<17 |
		s.workerIdBits<<12 |
		s.sequenceBits
	return nextId, nil
}

// 阻塞取下一毫秒
func (s *Snowflake) getNextMs(timeStamp uint64) uint64 {
	currentTime := uint64(time.Now().UnixMilli())
	for currentTime <= timeStamp {
		currentTime = uint64(time.Now().UnixMilli())
	}
	return currentTime
}

var snowflakeGenerator *Snowflake
var once sync.Once

func GetInstance() *Snowflake {
	once.Do(func() {
		snowflakeGenerator = &Snowflake{
			startPoint:       uint64(time.Now().UnixMilli()),
			sequenceBits:     0,
			workerIdBits:     0,
			dataCenterIdBits: 0,
			mutex:            sync.Mutex{},
		}
	})
	return snowflakeGenerator
}
