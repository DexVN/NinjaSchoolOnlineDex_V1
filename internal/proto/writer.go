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

func (w *Writer) WriteBytes(data []byte) {
	w.buf = append(w.buf, data...)
}

func (w *Writer) WriteInt(v int32) {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], uint32(v))
	w.buf = append(w.buf, b[:]...)
}

func (w *Writer) GetData() []byte {
	return w.buf
}
