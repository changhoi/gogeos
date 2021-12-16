package geos

/*
#include "geos.h"
*/
import "C"

import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

type wkbDecoder struct {
	r *C.GEOSWKBReader
}

func newWkbDecoder() *wkbDecoder {
	return &wkbDecoder{}
}

func (d *wkbDecoder) decode(wkb []byte) (*Geometry, error) {
	r := cGEOSWKBReader_create()
	if r == nil {
		return nil, fmt.Errorf("cannot initialize decoder")
	}

	defer cGEOSWKBReader_destroy(r)

	var cwkb []C.uchar
	for i := range wkb {
		cwkb = append(cwkb, C.uchar(wkb[i]))
	}
	g := cGEOSWKBReader_read(r, &cwkb[0], C.size_t(len(wkb)))
	if g == nil {
		return nil, Error()
	}

	return geomFromPtr(g), nil
}

func (d *wkbDecoder) decodeHex(wkbHex string) (*Geometry, error) {
	wkb, err := hex.DecodeString(wkbHex)
	if err != nil {
		return nil, err
	}
	return d.decode(wkb)
}

type wkbEncoder struct {
}

func newWkbEncoder() *wkbEncoder {
	return &wkbEncoder{}
}

func encodeWkb(g *Geometry, fn func(*C.GEOSWKBWriter, *C.GEOSGeometry, *C.size_t) *C.uchar) ([]byte, error) {
	w := cGEOSWKBWriter_create()
	if w == nil {
		return nil, fmt.Errorf("cannot initialize encoder")
	}

	defer cGEOSWKBWriter_destroy(w)

	var size C.size_t
	bytes := fn(w, g.g, &size)
	if bytes == nil {
		return nil, Error()
	}

	ptr := unsafe.Pointer(bytes)
	defer C.free(ptr)

	l := int(size)
	var out []byte
	for i := 0; i < l; i++ {
		el := unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(C.uchar(0))*uintptr(i))
		out = append(out, byte(*(*C.uchar)(el)))
	}

	return out, nil
}

func (e *wkbEncoder) encode(g *Geometry) ([]byte, error) {
	return encodeWkb(g, cGEOSWKBWriter_write)
}

func (e *wkbEncoder) encodeHex(g *Geometry) ([]byte, error) {
	return encodeWkb(g, cGEOSWKBWriter_writeHEX)
}
