//go:build armbe || arm64be || ppc64 || mips || mips64 || mips64p32 || ppc || sparc || sparc64 || s390 || s390x
// +build armbe arm64be ppc64 mips mips64 mips64p32 ppc sparc sparc64 s390 s390x

package serialize_wrapper

import (
	"encoding/binary"
	"math"
	"unsafe"
)

// Uint16 reads a little endian uint16 from the underlying buffer.
func (r *Reader) Uint16(x *uint16) {
	b := make([]byte, 2)
	if _, err := r.r.Read(b); err != nil {
		r.panic(err)
	}
	*x = binary.BigEndian.Uint16(b)
}

// Int16 reads a little endian int16 from the underlying buffer.
func (r *Reader) Int16(x *int16) {
	b := make([]byte, 2)
	if _, err := r.r.Read(b); err != nil {
		r.panic(err)
	}
	*x = int16(binary.BigEndian.Uint16(b))
}

// Uint32 reads a little endian uint32 from the underlying buffer.
func (r *Reader) Uint32(x *uint32) {
	b := make([]byte, 4)
	if _, err := r.r.Read(b); err != nil {
		r.panic(err)
	}
	*x = binary.BigEndian.Uint32(b)
}

// Int32 reads a little endian int32 from the underlying buffer.
func (r *Reader) Int32(x *int32) {
	b := make([]byte, 4)
	if _, err := r.r.Read(b); err != nil {
		r.panic(err)
	}
	*x = int32(binary.BigEndian.Uint32(b))
}

// BEInt32 reads a big endian int32 from the underlying buffer.
func (r *Reader) BEInt32(x *int32) {
	b := make([]byte, 4)
	if _, err := r.r.Read(b); err != nil {
		r.panic(err)
	}
	*x = *(*int32)(unsafe.Pointer(&b[0]))
}

// Uint64 reads a little endian uint64 from the underlying buffer.
func (r *Reader) Uint64(x *uint64) {
	b := make([]byte, 8)
	if _, err := r.r.Read(b); err != nil {
		r.panic(err)
	}
	*x = binary.BigEndian.Uint64(b)
}

// Int64 reads a little endian int64 from the underlying buffer.
func (r *Reader) Int64(x *int64) {
	b := make([]byte, 8)
	if _, err := r.r.Read(b); err != nil {
		r.panic(err)
	}
	*x = int64(binary.BigEndian.Uint64(b))
}

// Float32 reads a little endian float32 from the underlying buffer.
func (r *Reader) Float32(x *float32) {
	b := make([]byte, 4)
	if _, err := r.r.Read(b); err != nil {
		r.panic(err)
	}
	*x = math.Float32frombits(binary.BigEndian.Uint32(b))
}
