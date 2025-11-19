package Quick

import (
	"errors"
)

const (
	MaxStreamID int64 = (int64(1) << 62) - 1 // 2^62-1
)

type StreamType = int64

const (
	ClientInitiatedBidi StreamType = iota + 0b_00 // 0b_00
	ServerInitiatedBidi                           // 0b_01
	ClientInitiatedUndi                           // 0b_10
	ServerInitiatedUndi                           // 0b_11
)

type StreamID struct {
	streamID int64
}

func NewStreamID(streamType StreamType) StreamID {
	streamId := StreamID{
		streamID: int64(streamType),
	}

	return streamId
}

var (
	IntegerOverflow error = errors.New("Integer overflow")
)

func (si *StreamID) Increment() error {
	si.streamID += 4
	if si.streamID > MaxStreamID {
		si.streamID -= 4
		return IntegerOverflow
	}
	return nil
}

func (si *StreamID) StreamType() StreamType {
	return si.streamID & 0b_11
}

func (si *StreamID) ToVariableLengthInt() ([]byte, error) {
	return ToVarint(si.streamID)
}
