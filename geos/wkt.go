package geos

/*
#include "geos.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Reads the WKT serialization and produces geometries
type wktDecoder struct {
	r *C.GEOSWKTReader
}

// Creates a new WKT decoder, can be nil if initialization in the C API fails
func newWktDecoder() *wktDecoder {
	return &wktDecoder{}
}

// decode decodes the WKT string and returns a geometry
func (d *wktDecoder) decode(wkt string) (*Geometry, error) {
	r := cGEOSWKTReader_create()
	if r == nil {
		return nil, fmt.Errorf("cannot initialize decoder")
	}

	defer cGEOSWKTReader_destroy(r)

	cstr := C.CString(wkt)
	defer C.free(unsafe.Pointer(cstr))

	g := cGEOSWKTReader_read(r, cstr)
	if g == nil {
		return nil, Error()
	}

	return geomFromPtr(g), nil
}

type wktEncoder struct {
}

func newWktEncoder() *wktEncoder {
	return &wktEncoder{}
}

// Encode returns a string that is the geometry encoded as WKT
func (e *wktEncoder) encode(g *Geometry) (string, error) {
	w := cGEOSWKTWriter_create()
	if w == nil {
		return "", fmt.Errorf("cannot initialize encoder")
	}

	defer cGEOSWKTWriter_destroy(w)

	cstr := cGEOSWKTWriter_write(w, g.g)
	if cstr == nil {
		return "", Error()
	}

	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr), nil
}
