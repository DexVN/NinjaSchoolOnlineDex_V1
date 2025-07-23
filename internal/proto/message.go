package proto

import (
	"encoding/binary"
	"io"
)

const messageHeaderSize = 3 // 1 byte Command + 2 byte Length

type Message struct {
	Command int8
	reader  *Reader
	writer  *Writer
}

// NewMessage creates a message for writing (outgoing)
func NewMessage(cmd int8) *Message {
	return &Message{
		Command: cmd,
		writer:  NewWriter(),
	}
}

// NewEmptyMessage creates a message with empty writer
func NewEmptyMessage() *Message {
	return &Message{
		writer: NewWriter(),
	}
}

// NewMessageFromData creates a message for reading (incoming)
func NewMessageFromData(cmd int8, data []byte) *Message {
	return &Message{
		Command: cmd,
		reader:  NewReader(data),
	}
}

func (m *Message) Reader() *Reader {
	return m.reader
}

func (m *Message) Writer() *Writer {
	if m.writer == nil {
		m.writer = NewWriter()
	}
	return m.writer
}

func (m *Message) GetData() []byte {
	if m.writer == nil {
		return nil
	}
	return m.writer.GetData()
}

func (m *Message) Len() int {
	return len(m.GetData())
}

func (m *Message) Cleanup() {
	m.reader = nil
	m.writer = nil
}

func (m *Message) WriteInt8(v int8) {
	m.Writer().WriteInt8(v)
}

func (m *Message) WriteInt16(v int16) {
	m.Writer().WriteInt16(v)
}

func (m *Message) WriteInt32(v int32) {
	m.Writer().WriteInt32(v)
}

func (m *Message) WriteBool(v bool) {
	m.Writer().WriteBool(v)
}

func (m *Message) WriteUTF(s string) {
	m.Writer().WriteUTF(s)
}

func (m *Message) WriteBytes(b []byte) {
	m.Writer().WriteBytes(b)
}

func xorByte(b byte, key []byte, index *int) byte {
	if len(key) == 0 {
		return b
	}
	res := b ^ key[*index]
	*index += 1
	if *index >= len(key) {
		*index = 0
	}
	return res
}

func WriteMessage(w io.Writer, msg *Message, key []byte, index *int) error {
	data := msg.GetData()
	length := uint16(len(data))
	buf := make([]byte, messageHeaderSize+len(data))

	buf[0] = byte(msg.Command)
	binary.BigEndian.PutUint16(buf[1:3], length)
	copy(buf[3:], data)

	if len(key) > 0 {
		for i := 0; i < len(buf); i++ {
			buf[i] = xorByte(buf[i], key, index)
		}
	}

	_, err := w.Write(buf)
	return err
}

func ReadMessage(r io.Reader, key []byte, index *int) (*Message, error) {
	opcodeBuf := make([]byte, 1)
	if _, err := io.ReadFull(r, opcodeBuf); err != nil {
		return nil, err
	}
	opcode := xorByte(opcodeBuf[0], key, index)

	lenBuf := make([]byte, 2)
	if _, err := io.ReadFull(r, lenBuf); err != nil {
		return nil, err
	}
	for i := range lenBuf {
		lenBuf[i] = xorByte(lenBuf[i], key, index)
	}
	length := binary.BigEndian.Uint16(lenBuf)

	payload := make([]byte, length)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}
	for i := range payload {
		payload[i] = xorByte(payload[i], key, index)
	}

	return NewMessageFromData(int8(opcode), payload), nil
}
