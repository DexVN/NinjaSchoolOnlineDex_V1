package proto

import (
	"encoding/binary"
)

type Writer struct {
	buf []byte
}

func NewWriter() *Writer {
	return &Writer{buf: make([]byte, 0, 128)}
}

func (w *Writer) WriteByte(b byte) error {
	w.buf = append(w.buf, b)
	return nil
}

func (w *Writer) WriteInt8(v int8) {
	w.buf = append(w.buf, byte(v))
}

func (w *Writer) WriteInt16(v int16) {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], uint16(v))
	w.buf = append(w.buf, b[:]...)
}

func (w *Writer) WriteInt32(v int32) {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], uint32(v))
	w.buf = append(w.buf, b[:]...)
}

// alias tương thích với code cũ
func (w *Writer) WriteInt(v int32) {
	w.WriteInt32(v)
}

func (w *Writer) WriteBool(v bool) {
	if v {
		w.buf = append(w.buf, 1)
	} else {
		w.buf = append(w.buf, 0)
	}
}

func (w *Writer) WriteUTF(s string) {
	strBytes := []byte(s)
	length := len(strBytes)
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], uint16(length))
	w.buf = append(w.buf, b[:]...)
	w.buf = append(w.buf, strBytes...)
}

func (w *Writer) WriteBytes(data []byte) {
	w.buf = append(w.buf, data...)
}

func (w *Writer) GetData() []byte {
	return w.buf
}
