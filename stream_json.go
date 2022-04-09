package response

import (
	"fmt"
	"io"
)

var (
	beginObjectBytes = []byte("{")
	endObjectBytes   = []byte("}")
	beginArrayBytes  = []byte("[")
	endArrayBytes    = []byte("]")
	separatorBytes   = []byte(",")
)

type JSONStream struct {
	Stream
	shouldWriteSeparator bool
}

type JSONSerializable interface {
	Encode(stream *JSONStream) error
}

func NewJSONStream(w io.Writer) *JSONStream {
	stream := NewStream(w)
	return &JSONStream{Stream: *stream}
}

func (s *JSONStream) BeginObject() error {
	_, err := s.writeSeparated(beginObjectBytes)
	s.shouldWriteSeparator = false
	return err
}

func (s *JSONStream) EndObject() error {
	_, err := s.Write(endObjectBytes)
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) BeginArray() error {
	_, err := s.writeSeparated(beginArrayBytes)
	s.shouldWriteSeparator = false
	return err
}

func (s *JSONStream) EndArray() error {
	_, err := s.Write(endArrayBytes)
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteKey(keyName string) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%q:", keyName)))
	s.shouldWriteSeparator = false
	return err
}

func (s *JSONStream) WriteString(value string) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%q", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteInt(value int) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteInt8(value int8) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteInt16(value int16) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteInt32(value int32) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteInt64(value int64) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteUint(value uint) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteUint8(value uint8) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteUint16(value uint16) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteUint32(value uint32) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteUint64(value uint64) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteFloat32(value float32) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%g", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteFloat64(value float64) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%g", value)))
	s.shouldWriteSeparator = true
	return err
}

func (s *JSONStream) WriteObject(obj JSONSerializable) error {
	err := obj.Encode(s)
	s.shouldWriteSeparator = true

	return err
}

func (s *JSONStream) writeSeparated(data []byte) (bytes int, err error) {
	if s.shouldWriteSeparator {
		bytes, err = s.Write(separatorBytes)
	}

	prevBytes := bytes
	bytes, err = s.Write(data)
	bytes += prevBytes

	return
}
