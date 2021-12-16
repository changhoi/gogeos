package geos

import (
	"testing"
)

var geoJSONEncoderTests = []struct{ in, middle, out string }{
	{`{"type":"Polygon","coordinates":[[[0,0],[0,2],[2,2],[2,0],[0,0]],[[1,1],[1,1.5],[1.5,1.5],[1.5,1],[1,1]]]}`,
		"POLYGON ((0 0, 0 2, 2 2, 2 0, 0 0), (1 1, 1 1.5, 1.5 1.5, 1.5 1, 1 1))",
		`{"type":"Polygon","coordinates":[[[0.0,0.0],[0.0,2.0],[2.0,2.0],[2.0,0.0],[0.0,0.0]],[[1.0,1.0],[1.0,1.5],[1.5,1.5],[1.5,1.0],[1.0,1.0]]]}`},
}

func TestGeoJSONEncoder(t *testing.T) {
	encoder := newGeoJSONEncoder()
	decoder := newGeoJSONDecoder()
	var geom *Geometry
	var err error
	for _, test := range geoJSONEncoderTests {
		geom, err = decoder.decode(test.in)
		if err != nil {
			t.Errorf("geoJSONDecoder.decode(): %v", err)
		}
		actual, err := encoder.encode(geom)
		if err != nil {
			t.Errorf("geoJSONEncoder.encode(): %v", err)
		}
		if actual != test.out {
			t.Errorf("geoJSONEncoder.encode(): want %v, got %v", test.out, actual)
		}
	}
}

