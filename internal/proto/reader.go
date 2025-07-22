package proto

import (
	"encoding/binary"
	"io"
)

type Reader struct {
	buf []byte
	pos int
}

func NewReader(data []byte) *Reader {
	return &Reader{buf: data}
}
func (r *Reader) ReadByte() (byte, error) {
	if r.pos >= len(r.buf) {
		return 0, io.EOF
	}
	b := r.buf[r.pos]
	r.pos++
	return b, nil
}

func (r *Reader) ReadBytes(n int) []byte {
	if r.pos+n > len(r.buf) {
		return nil
	}
	b := r.buf[r.pos : r.pos+n]
	r.pos += n
	return b
}

func (r *Reader) ReadBool() (bool, error) {
	b, err := r.ReadByte()
	if err != nil {
		return false, err
	}
	return b != 0, nil
}

func (r *Reader) ReadInt() int32 {
	if r.pos+4 > len(r.buf) {
		return 0
	}
	v := binary.BigEndian.Uint32(r.buf[r.pos:])
	r.pos += 4
	return int32(v)
}

func (r *Reader) ReadUTF() (string, error) {
	if r.pos+2 > len(r.buf) {
		return "", io.EOF
	}
	length := int(binary.BigEndian.Uint16(r.buf[r.pos:]))
	r.pos += 2

	if r.pos+length > len(r.buf) {
		return "", io.EOF
	}
	strBuf := r.buf[r.pos : r.pos+length]
	r.pos += length
	return string(strBuf), nil
}

func (r *Reader) Len() int {
	return len(r.buf) - r.pos
}
