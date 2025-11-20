package Quick

import (
	"io"
)

/*
STREAM Frame {
  Type (i) = 0x08..0x0f,
  Stream ID (i),
  [Offset (i)],
  [Length (i)],
  Stream Data (..),
}
*/

// 0x08 to 0x0f
type StreamFrameType uint8

func (sft StreamFrameType) Offset() bool {
	notImplemented()
	return false
}

func (sft StreamFrameType) Length() bool {
	notImplemented()
	return false
}

func (sft StreamFrameType) Fin() bool {
	notImplemented()
	return false
}

// STREAM frames implicitly create a stream and carry stream data.
type StreamFrame struct {
	Type     StreamFrameType //Half the byte is empty. Only LS 4 bytes is in use.
	StreamID StreamID

	// Offset is starting index which the StreamData should be place in the stream.
	Offset Int62
	// Offset is optional and uses variable length encoding.
	// If Offset of the frame is not specified in the Type of the frame. Offset of the frame is considered 0.

	// When the Length is 0, the Offset in the STREAM frame is the offset of the next byte that would be sent.
	Length Int62 // Length is optional and uses variable length encoding.
	// Offset of the the stream and the Length of the frame cannot overflow int62.

	StreamData io.Reader
}
