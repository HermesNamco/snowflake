package snowflake

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestGetInstance(t *testing.T) {
	tests := []struct {
		name string
		want *Snowflake
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_DataCenterIdBits(t *testing.T) {
	type fields struct {
		startPoint       uint64
		sequenceBits     uint64
		workerIdBits     uint64
		dataCenterIdBits uint64
		lastTimeStamp    uint64
		mutex            sync.Mutex
	}
	type args struct {
		dataCentIdBits uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Snowflake
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Snowflake{
				startPoint:       tt.fields.startPoint,
				sequenceBits:     tt.fields.sequenceBits,
				workerIdBits:     tt.fields.workerIdBits,
				dataCenterIdBits: tt.fields.dataCenterIdBits,
				lastTimeStamp:    tt.fields.lastTimeStamp,
				mutex:            tt.fields.mutex,
			}
			if got := s.DataCenterIdBits(tt.args.dataCentIdBits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataCenterIdBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_Next(t *testing.T) {
	type fields struct {
		startPoint       uint64
		sequenceBits     uint64
		workerIdBits     uint64
		dataCenterIdBits uint64
		lastTimeStamp    uint64
		mutex            sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "first",
			fields:  fields{0, 0, 0, 0, 0, sync.Mutex{}},
			want:    uint64(time.Now().UnixMilli() << 22),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Snowflake{
				startPoint:       tt.fields.startPoint,
				sequenceBits:     tt.fields.sequenceBits,
				workerIdBits:     tt.fields.workerIdBits,
				dataCenterIdBits: tt.fields.dataCenterIdBits,
				lastTimeStamp:    tt.fields.lastTimeStamp,
				mutex:            tt.fields.mutex,
			}
			got, err := s.Next()
			if (err != nil) != tt.wantErr {
				t.Errorf("Next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Next() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_StartPoint(t *testing.T) {
	type fields struct {
		startPoint       uint64
		sequenceBits     uint64
		workerIdBits     uint64
		dataCenterIdBits uint64
		lastTimeStamp    uint64
		mutex            sync.Mutex
	}
	type args struct {
		startPoint uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Snowflake
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Snowflake{
				startPoint:       tt.fields.startPoint,
				sequenceBits:     tt.fields.sequenceBits,
				workerIdBits:     tt.fields.workerIdBits,
				dataCenterIdBits: tt.fields.dataCenterIdBits,
				lastTimeStamp:    tt.fields.lastTimeStamp,
				mutex:            tt.fields.mutex,
			}
			if got := s.StartPoint(tt.args.startPoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_WorkerIdBits(t *testing.T) {
	type fields struct {
		startPoint       uint64
		sequenceBits     uint64
		workerIdBits     uint64
		dataCenterIdBits uint64
		lastTimeStamp    uint64
		mutex            sync.Mutex
	}
	type args struct {
		workIdBits uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Snowflake
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Snowflake{
				startPoint:       tt.fields.startPoint,
				sequenceBits:     tt.fields.sequenceBits,
				workerIdBits:     tt.fields.workerIdBits,
				dataCenterIdBits: tt.fields.dataCenterIdBits,
				lastTimeStamp:    tt.fields.lastTimeStamp,
				mutex:            tt.fields.mutex,
			}
			if got := s.WorkerIdBits(tt.args.workIdBits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WorkerIdBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_getNextMs(t *testing.T) {
	type fields struct {
		startPoint       uint64
		sequenceBits     uint64
		workerIdBits     uint64
		dataCenterIdBits uint64
		lastTimeStamp    uint64
		mutex            sync.Mutex
	}
	type args struct {
		timeStamp uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Snowflake{
				startPoint:       tt.fields.startPoint,
				sequenceBits:     tt.fields.sequenceBits,
				workerIdBits:     tt.fields.workerIdBits,
				dataCenterIdBits: tt.fields.dataCenterIdBits,
				lastTimeStamp:    tt.fields.lastTimeStamp,
				mutex:            tt.fields.mutex,
			}
			if got := s.getNextMs(tt.args.timeStamp); got != tt.want {
				t.Errorf("getNextMs() = %v, want %v", got, tt.want)
			}
		})
	}
}
