package proto

import (
	"encoding/binary"
	"io"
)

type Message struct {
	Command int8    // giống sbyte
	reader  *Reader // đọc từ []byte
	writer  *Writer // ghi ra []byte
}

// Constructor: chỉ có opcode
func NewMessage(cmd int8) *Message {
	return &Message{
		Command: cmd,
		writer:  NewWriter(),
	}
}

func NewEmptyMessage() *Message {
	return &Message{
		writer: NewWriter(),
	}
}

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

func (m *Message) Cleanup() {
	m.reader = nil
	m.writer = nil
}

func xorByte(b byte, key []byte, index *int) byte {
	if len(key) == 0 {
		return b
	}
	res := b ^ key[*index]
	*index++
	if *index >= len(key) {
		*index = 0
	}
	return res
}

func WriteMessage(w io.Writer, msg Message, key []byte, index *int) error {
	data := msg.GetData()
	length := uint16(len(data))
	buf := make([]byte, 1+2+len(data))
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
