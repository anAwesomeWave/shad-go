//go:build !solution

package otp

import (
	"io"
)

type XorReader struct {
	r    io.Reader
	prng io.Reader
}

func (xr *XorReader) Read(p []byte) (int, error) {
	sz := len(p)

	// принимает слайс p, декодирует n <= len(p) байт из потока r

	buf := make([]byte, sz)
	sz, err := xr.r.Read(buf)

	prng := make([]byte, sz)
	_, _ = xr.prng.Read(prng)

	for i := 0; i < sz; i++ {
		p[i] = buf[i] ^ prng[i]
	}

	return sz, err

}

type XorWriter struct {
	w    io.Writer
	prng io.Reader
}

func (xw *XorWriter) Write(p []byte) (int, error) {
	sz := len(p)
	prng := make([]byte, sz)
	_, _ = xw.prng.Read(prng)
	encoded := make([]byte, len(p))
	for i := 0; i < sz; i++ {
		encoded[i] = p[i] ^ prng[i]
	}
	n, err := xw.w.Write(encoded)
	if n != sz || err != nil {
		return n, err
	}
	return sz, nil
}

func NewReader(r io.Reader, prng io.Reader) io.Reader {
	// хотим создать reader
	// будет декодировать p байт
	return &XorReader{
		r,
		prng,
	}
}

func NewWriter(w io.Writer, prng io.Reader) io.Writer {
	return &XorWriter{
		w,
		prng,
	}
}
