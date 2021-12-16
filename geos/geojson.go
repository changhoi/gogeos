package geos

/*
#include "geos.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type geoJSONDecoder struct {
}

func newGeoJSONDecoder() *geoJSONDecoder {
	return &geoJSONDecoder{}
}

func (d *geoJSONDecoder) decode(geoJSON string) (*Geometry, error) {
	r := cGEOSGeoJSONReader_create()
	if r == nil {
		return nil, fmt.Errorf("cannot initialize decoder")
	}

	defer cGEOSGeoJSONReader_destroy(r)

	cstr := C.CString(geoJSON)
	defer C.free(unsafe.Pointer(cstr))

	g := cGEOSGeoJSONReader_readGeometry(r, cstr)
	if g == nil {
		return nil, Error()
	}

	return geomFromPtr(g), nil
}

type geoJSONEncoder struct {
}

func newGeoJSONEncoder() *geoJSONEncoder {
	return &geoJSONEncoder{}
}

func (e *geoJSONEncoder) encode(g *Geometry) (string, error) {
	w := cGEOSGeoJSONWriter_create()
	if w == nil {
		return "", fmt.Errorf("cannot initialize encoder")
	}

	defer cGEOSGeoJSONWriter_destroy(w)

	cstr := cGEOSGeoJSONWriter_writeGeometry(w, g.g, -1)
	if cstr == nil {
		return "", Error()
	}

	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr), nil
}
