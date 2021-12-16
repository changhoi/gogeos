package geos

// Created mechanically from C API header - DO NOT EDIT

/*
#include <geos_c.h>
*/
import "C"

import (
	"unsafe"
)

func cGEOS_interruptRegisterCallback(cb *C.GEOSInterruptCallback) *C.GEOSInterruptCallback {
	return C.GEOS_interruptRegisterCallback(cb)
}

func cGEOS_interruptRequest() {
	C.GEOS_interruptRequest()
}

func cGEOS_interruptCancel() {
	C.GEOS_interruptCancel()
}

func cGEOSversion() *C.char {
	return C.GEOSversion()
}

func cinitGEOS(notice_function C.GEOSMessageHandler, error_function C.GEOSMessageHandler) C.GEOSContextHandle_t {
	return C.initGEOS_r(notice_function, error_function)
}

func cfinishGEOS() {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.finishGEOS_r(handle)
}

func cGEOSCoordSeq_create(size C.uint, dims C.uint) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_create_r(handle, size, dims)
}

func cGEOSCoordSeq_copyFromBuffer(buf *C.double, size C.uint, hasZ C.int, hasM C.int) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_copyFromBuffer_r(handle, buf, size, hasZ, hasM)
}

func cGEOSCoordSeq_copyFromArrays(x *C.double, y *C.double, z *C.double, m *C.double, size C.uint) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_copyFromArrays_r(handle, x, y, z, m, size)
}

func cGEOSCoordSeq_copyToBuffer(s *C.GEOSCoordSequence, buf *C.double, hasZ C.int, hasM C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_copyToBuffer_r(handle, s, buf, hasZ, hasM)
}

func cGEOSCoordSeq_copyToArrays(s *C.GEOSCoordSequence, x *C.double, y *C.double, z *C.double, m *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_copyToArrays_r(handle, s, x, y, z, m)
}

func cGEOSCoordSeq_clone(s *C.GEOSCoordSequence) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_clone_r(handle, s)
}

func cGEOSCoordSeq_destroy(s *C.GEOSCoordSequence) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSCoordSeq_destroy_r(handle, s)
}

func cGEOSCoordSeq_setX(s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setX_r(handle, s, idx, val)
}

func cGEOSCoordSeq_setY(s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setY_r(handle, s, idx, val)
}

func cGEOSCoordSeq_setZ(s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setZ_r(handle, s, idx, val)
}

func cGEOSCoordSeq_setXY(s *C.GEOSCoordSequence, idx C.uint, x C.double, y C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setXY_r(handle, s, idx, x, y)
}

func cGEOSCoordSeq_setXYZ(s *C.GEOSCoordSequence, idx C.uint, x C.double, y C.double, z C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setXYZ_r(handle, s, idx, x, y, z)
}

func cGEOSCoordSeq_setOrdinate(s *C.GEOSCoordSequence, idx C.uint, dim C.uint, val C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setOrdinate_r(handle, s, idx, dim, val)
}

func cGEOSCoordSeq_getX(s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getX_r(handle, s, idx, val)
}

func cGEOSCoordSeq_getY(s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getY_r(handle, s, idx, val)
}

func cGEOSCoordSeq_getZ(s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getZ_r(handle, s, idx, val)
}

func cGEOSCoordSeq_getXY(s *C.GEOSCoordSequence, idx C.uint, x *C.double, y *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getXY_r(handle, s, idx, x, y)
}

func cGEOSCoordSeq_getXYZ(s *C.GEOSCoordSequence, idx C.uint, x *C.double, y *C.double, z *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getXYZ_r(handle, s, idx, x, y, z)
}

func cGEOSCoordSeq_getOrdinate(s *C.GEOSCoordSequence, idx C.uint, dim C.uint, val *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getOrdinate_r(handle, s, idx, dim, val)
}

func cGEOSCoordSeq_getSize(s *C.GEOSCoordSequence, size *C.uint) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getSize_r(handle, s, size)
}

func cGEOSCoordSeq_getDimensions(s *C.GEOSCoordSequence, dims *C.uint) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_getDimensions_r(handle, s, dims)
}

func cGEOSCoordSeq_isCCW(s *C.GEOSCoordSequence, is_ccw *C.char) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_isCCW_r(handle, s, is_ccw)
}

func cGEOSProject(line *C.GEOSGeometry, point *C.GEOSGeometry) C.double {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSProject_r(handle, line, point)
}

func cGEOSInterpolate(line *C.GEOSGeometry, d C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSInterpolate_r(handle, line, d)
}

func cGEOSProjectNormalized(g *C.GEOSGeometry, p *C.GEOSGeometry) C.double {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSProjectNormalized_r(handle, g, p)
}

func cGEOSInterpolateNormalized(g *C.GEOSGeometry, d C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSInterpolateNormalized_r(handle, g, d)
}

func cGEOSBuffer(g *C.GEOSGeometry, width C.double, quadsegs C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBuffer_r(handle, g, width, quadsegs)
}

func cGEOSBufferParams_create() *C.GEOSBufferParams {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_create_r(handle)
}

func cGEOSBufferParams_destroy(parms *C.GEOSBufferParams) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSBufferParams_destroy_r(handle, parms)
}

func cGEOSBufferParams_setEndCapStyle(p *C.GEOSBufferParams, style C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setEndCapStyle_r(handle, p, style)
}

func cGEOSBufferParams_setJoinStyle(p *C.GEOSBufferParams, joinStyle C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setJoinStyle_r(handle, p, joinStyle)
}

func cGEOSBufferParams_setMitreLimit(p *C.GEOSBufferParams, mitreLimit C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setMitreLimit_r(handle, p, mitreLimit)
}

func cGEOSBufferParams_setQuadrantSegments(p *C.GEOSBufferParams, quadSegs C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setQuadrantSegments_r(handle, p, quadSegs)
}

func cGEOSBufferParams_setSingleSided(p *C.GEOSBufferParams, singleSided C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setSingleSided_r(handle, p, singleSided)
}

func cGEOSBufferWithParams(g *C.GEOSGeometry, p *C.GEOSBufferParams, width C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferWithParams_r(handle, g, p, width)
}

func cGEOSBufferWithStyle(g *C.GEOSGeometry, width C.double, quadsegs C.int, endCapStyle C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferWithStyle_r(handle, g, width, quadsegs, endCapStyle, joinStyle, mitreLimit)
}

func cGEOSDensify(g *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDensify_r(handle, g, tolerance)
}

func cGEOSOffsetCurve(g *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSOffsetCurve_r(handle, g, width, quadsegs, joinStyle, mitreLimit)
}

func cGEOSGeom_createPoint(s *C.GEOSCoordSequence) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createPoint_r(handle, s)
}

func cGEOSGeom_createPointFromXY(x C.double, y C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createPointFromXY_r(handle, x, y)
}

func cGEOSGeom_createEmptyPoint() *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createEmptyPoint_r(handle)
}

func cGEOSGeom_createLinearRing(s *C.GEOSCoordSequence) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createLinearRing_r(handle, s)
}

func cGEOSGeom_createLineString(s *C.GEOSCoordSequence) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createLineString_r(handle, s)
}

func cGEOSGeom_createEmptyLineString() *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createEmptyLineString_r(handle)
}

func cGEOSGeom_createEmptyPolygon() *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createEmptyPolygon_r(handle)
}

func cGEOSGeom_createPolygon(shell *C.GEOSGeometry, holes **C.GEOSGeometry, nholes C.uint) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createPolygon_r(handle, shell, holes, nholes)
}

func cGEOSGeom_createCollection(_type C.int, geoms **C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createCollection_r(handle, _type, geoms, ngeoms)
}

func cGEOSGeom_createEmptyCollection(_type C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_createEmptyCollection_r(handle, _type)
}

func cGEOSGeom_clone(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_clone_r(handle, g)
}

func cGEOSGeom_destroy(g *C.GEOSGeometry) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSGeom_destroy_r(handle, g)
}

func cGEOSEnvelope(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSEnvelope_r(handle, g)
}

func cGEOSIntersection(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSIntersection_r(handle, g1, g2)
}

func cGEOSIntersectionPrec(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, gridSize C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSIntersectionPrec_r(handle, g1, g2, gridSize)
}

func cGEOSDifference(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDifference_r(handle, g1, g2)
}

func cGEOSDifferencePrec(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, gridSize C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDifferencePrec_r(handle, g1, g2, gridSize)
}

func cGEOSSymDifference(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSymDifference_r(handle, g1, g2)
}

func cGEOSSymDifferencePrec(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, gridSize C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSymDifferencePrec_r(handle, g1, g2, gridSize)
}

func cGEOSUnion(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSUnion_r(handle, g1, g2)
}

func cGEOSUnionPrec(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, gridSize C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSUnionPrec_r(handle, g1, g2, gridSize)
}

func cGEOSUnaryUnion(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSUnaryUnion_r(handle, g)
}

func cGEOSUnaryUnionPrec(g *C.GEOSGeometry, gridSize C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSUnaryUnionPrec_r(handle, g, gridSize)
}

func cGEOSBoundary(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBoundary_r(handle, g)
}

func cGEOSConvexHull(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSConvexHull_r(handle, g)
}

func cGEOSMinimumRotatedRectangle(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMinimumRotatedRectangle_r(handle, g)
}

func cGEOSMaximumInscribedCircle(g *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMaximumInscribedCircle_r(handle, g, tolerance)
}

func cGEOSLargestEmptyCircle(g *C.GEOSGeometry, boundary *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSLargestEmptyCircle_r(handle, g, boundary, tolerance)
}

func cGEOSMinimumWidth(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMinimumWidth_r(handle, g)
}

func cGEOSMinimumClearance(g *C.GEOSGeometry, distance *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMinimumClearance_r(handle, g, distance)
}

func cGEOSMinimumClearanceLine(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMinimumClearanceLine_r(handle, g)
}

func cGEOSCoverageUnion(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoverageUnion_r(handle, g)
}

func cGEOSPointOnSurface(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPointOnSurface_r(handle, g)
}

func cGEOSGetCentroid(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetCentroid_r(handle, g)
}

func cGEOSMinimumBoundingCircle(g *C.GEOSGeometry, radius *C.double, center **C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMinimumBoundingCircle_r(handle, g, radius, center)
}

func cGEOSNode(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSNode_r(handle, g)
}

func cGEOSClipByRect(g *C.GEOSGeometry, xmin C.double, ymin C.double, xmax C.double, ymax C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSClipByRect_r(handle, g, xmin, ymin, xmax, ymax)
}

func cGEOSPolygonize(geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPolygonize_r(handle, &geoms[0], ngeoms)
}

func cGEOSPolygonize_valid(geoms []*C.GEOSGeometry, ngems C.uint) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPolygonize_valid_r(handle, &geoms[0], ngems)
}

func cGEOSPolygonizer_getCutEdges(geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPolygonizer_getCutEdges_r(handle, &geoms[0], ngeoms)
}

func cGEOSPolygonize_full(input *C.GEOSGeometry, cuts **C.GEOSGeometry, dangles **C.GEOSGeometry, invalidRings **C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPolygonize_full_r(handle, input, cuts, dangles, invalidRings)
}

func cGEOSBuildArea(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBuildArea_r(handle, g)
}

func cGEOSLineMerge(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSLineMerge_r(handle, g)
}

func cGEOSReverse(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSReverse_r(handle, g)
}

func cGEOSSimplify(g *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSimplify_r(handle, g, tolerance)
}

func cGEOSTopologyPreserveSimplify(g *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSTopologyPreserveSimplify_r(handle, g, tolerance)
}

func cGEOSGeom_extractUniquePoints(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_extractUniquePoints_r(handle, g)
}

func cGEOSSharedPaths(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSharedPaths_r(handle, g1, g2)
}

func cGEOSSnap(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSnap_r(handle, g1, g2, tolerance)
}

func cGEOSDelaunayTriangulation(g *C.GEOSGeometry, tolerance C.double, onlyEdges C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDelaunayTriangulation_r(handle, g, tolerance, onlyEdges)
}

func cGEOSConstrainedDelaunayTriangulation(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSConstrainedDelaunayTriangulation_r(handle, g)
}

func cGEOSVoronoiDiagram(g *C.GEOSGeometry, env *C.GEOSGeometry, tolerance C.double, onlyEdges C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSVoronoiDiagram_r(extHandle, g, env, tolerance, onlyEdges)
}

func cGEOSSegmentIntersection(ax0 C.double, ay0 C.double, ax1 C.double, ay1 C.double, bx0 C.double, by0 C.double, bx1 C.double, by1 C.double, cx *C.double, cy *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSegmentIntersection_r(extHandle, ax0, ay0, ax1, ay1, bx0, by0, bx1, by1, cx, cy)
}

func cGEOSDisjoint(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDisjoint_r(handle, g1, g2)
}

func cGEOSTouches(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSTouches_r(handle, g1, g2)
}

func cGEOSIntersects(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSIntersects_r(handle, g1, g2)
}

func cGEOSCrosses(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCrosses_r(handle, g1, g2)
}

func cGEOSWithin(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWithin_r(handle, g1, g2)
}

func cGEOSContains(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSContains_r(handle, g1, g2)
}

func cGEOSOverlaps(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSOverlaps_r(handle, g1, g2)
}

func cGEOSEquals(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSEquals_r(handle, g1, g2)
}

func cGEOSCovers(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCovers_r(handle, g1, g2)
}

func cGEOSCoveredBy(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoveredBy_r(handle, g1, g2)
}

func cGEOSEqualsExact(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSEqualsExact_r(handle, g1, g2, tolerance)
}

func cGEOSPrepare(g *C.GEOSGeometry) *C.GEOSPreparedGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPrepare_r(handle, g)
}

func cGEOSPreparedGeom_destroy(g *C.GEOSPreparedGeometry) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSPreparedGeom_destroy_r(handle, g)
}

func cGEOSPreparedContains(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedContains_r(handle, pg1, g2)
}

func cGEOSPreparedContainsProperly(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedContainsProperly_r(handle, pg1, g2)
}

func cGEOSPreparedCoveredBy(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedCoveredBy_r(handle, pg1, g2)
}

func cGEOSPreparedCovers(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedCovers_r(handle, pg1, g2)
}

func cGEOSPreparedCrosses(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedCrosses_r(handle, pg1, g2)
}

func cGEOSPreparedDisjoint(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedDisjoint_r(handle, pg1, g2)
}

func cGEOSPreparedIntersects(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedIntersects_r(handle, pg1, g2)
}

func cGEOSPreparedOverlaps(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedOverlaps_r(handle, pg1, g2)
}

func cGEOSPreparedTouches(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedTouches_r(handle, pg1, g2)
}

func cGEOSPreparedWithin(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedWithin_r(handle, pg1, g2)
}

func cGEOSPreparedNearestPoints(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedNearestPoints_r(handle, pg1, g2)
}

func cGEOSPreparedDistance(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedDistance_r(handle, pg1, g2, dist)
}

func cGEOSPreparedDistanceWithin(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry, dist C.double) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSPreparedDistanceWithin_r(handle, pg1, g2, dist)
}

func cGEOSSTRtree_create(nodeCapacity C.size_t) *C.GEOSSTRtree {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSTRtree_create_r(handle, nodeCapacity)
}

func cGEOSSTRtree_insert(tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSTRtree_insert_r(handle, tree, g, unsafe.Pointer(item))
}

func cGEOSSTRtree_query(tree *C.GEOSSTRtree, g *C.GEOSGeometry, callback C.GEOSQueryCallback, userdata *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSTRtree_query_r(handle, tree, g, callback, unsafe.Pointer(userdata))
}

func cGEOSSTRtree_nearest(tree *C.GEOSSTRtree, geom *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSTRtree_nearest_r(handle, tree, geom)
}

func cGEOSSTRtree_nearest_generic(tree *C.GEOSSTRtree, item *C.void, itemEnvelope *C.GEOSGeometry, distancefn C.GEOSDistanceCallback, userdata *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSTRtree_nearest_generic_r(handle, tree, unsafe.Pointer(item), itemEnvelope, distancefn, unsafe.Pointer(userdata))
}

func cGEOSSTRtree_iterate(tree *C.GEOSSTRtree, callback C.GEOSQueryCallback, userdata *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSTRtree_iterate_r(handle, tree, callback, unsafe.Pointer(userdata))
}

func cGEOSSTRtree_remove(tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSTRtree_remove_r(handle, tree, g, unsafe.Pointer(item))
}

func cGEOSSTRtree_destroy(tree *C.GEOSSTRtree) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSTRtree_destroy_r(handle, tree)
}

func cGEOSisEmpty(g *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisEmpty_r(handle, g)
}

func cGEOSisSimple(g *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisSimple_r(handle, g)
}

func cGEOSisRing(g *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisRing_r(handle, g)
}

func cGEOSHasZ(g *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSHasZ_r(handle, g)
}

func cGEOSisClosed(g *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisClosed_r(handle, g)
}

func cGEOSRelatePattern(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, pat *C.char) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSRelatePattern_r(handle, g1, g2, pat)
}

func cGEOSRelate(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSRelate_r(handle, g1, g2)
}

func cGEOSRelatePatternMatch(mat *C.char, pat *C.char) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSRelatePatternMatch_r(handle, mat, pat)
}

func cGEOSRelateBoundaryNodeRule(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, bnr C.int) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSRelateBoundaryNodeRule_r(handle, g1, g2, bnr)
}

func cGEOSisValid(g *C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisValid_r(handle, g)
}

func cGEOSisValidReason(g *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisValidReason_r(handle, g)
}

func cGEOSisValidDetail(g *C.GEOSGeometry, flags C.int, reason **C.char, location **C.GEOSGeometry) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSisValidDetail_r(handle, g, flags, reason, location)
}

func cGEOSMakeValid(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMakeValid_r(handle, g)
}

func cGEOSMakeValidWithParams(g *C.GEOSGeometry, makeValidParams *C.GEOSMakeValidParams) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMakeValidWithParams_r(handle, g, makeValidParams)
}

func cGEOSMakeValidParams_create() *C.GEOSMakeValidParams {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMakeValidParams_create_r(extHandle)
}

func cGEOSMakeValidParams_destroy(parms *C.GEOSMakeValidParams) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSMakeValidParams_destroy_r(handle, parms)
}

func cGEOSMakeValidParams_setMethod(p *C.GEOSMakeValidParams, method C.enum_GEOSMakeValidMethods) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMakeValidParams_setMethod_r(handle, p, method)
}

func cGEOSMakeValidParams_setKeepCollapsed(p *C.GEOSMakeValidParams, style C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSMakeValidParams_setKeepCollapsed_r(handle, p, style)
}

func cGEOSGeomType(g *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomType_r(handle, g)
}

func cGEOSGeomTypeId(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomTypeId_r(handle, g)
}

func cGEOSGetSRID(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetSRID_r(handle, g)
}

func cGEOSSetSRID(g *C.GEOSGeometry, SRID C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSSetSRID_r(handle, g, SRID)
}

func cGEOSGeom_getUserData(g *C.GEOSGeometry) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSGeom_getUserData_r(handle, g)
}

func cGEOSGeom_setUserData(g *C.GEOSGeometry, userData *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSGeom_setUserData_r(handle, g, unsafe.Pointer(userData))
}

func cGEOSGetNumGeometries(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetNumGeometries_r(handle, g)
}

func cGEOSGetGeometryN(g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetGeometryN_r(handle, g, n)
}

func cGEOSNormalize(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSNormalize_r(handle, g)
}

func cGEOSGeom_setPrecision(g *C.GEOSGeometry, gridSize C.double, flags C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_setPrecision_r(handle, g, gridSize, flags)
}

func cGEOSGeom_getPrecision(g *C.GEOSGeometry) C.double {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getPrecision_r(handle, g)
}

func cGEOSGetNumInteriorRings(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetNumInteriorRings_r(handle, g)
}

func cGEOSGeomGetNumPoints(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetNumPoints_r(handle, g)
}

func cGEOSGeomGetX(g *C.GEOSGeometry, x *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetX_r(handle, g, x)
}

func cGEOSGeomGetY(g *C.GEOSGeometry, y *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetY_r(handle, g, y)
}

func cGEOSGeomGetZ(g *C.GEOSGeometry, z *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetZ_r(handle, g, z)
}

func cGEOSGetInteriorRingN(g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetInteriorRingN_r(handle, g, n)
}

func cGEOSGetExteriorRing(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetExteriorRing_r(handle, g)
}

func cGEOSGetNumCoordinates(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGetNumCoordinates_r(handle, g)
}

func cGEOSGeom_getCoordSeq(g *C.GEOSGeometry) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getCoordSeq_r(handle, g)
}

func cGEOSGeom_getDimensions(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getDimensions_r(handle, g)
}

func cGEOSGeom_getCoordinateDimension(g *C.GEOSGeometry) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getCoordinateDimension_r(handle, g)
}

func cGEOSGeom_getXMin(g *C.GEOSGeometry, value *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getXMin_r(handle, g, value)
}

func cGEOSGeom_getYMin(g *C.GEOSGeometry, value *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getYMin_r(handle, g, value)
}

func cGEOSGeom_getXMax(g *C.GEOSGeometry, value *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getXMax_r(handle, g, value)
}

func cGEOSGeom_getYMax(g *C.GEOSGeometry, value *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeom_getYMax_r(handle, g, value)
}

func cGEOSGeomGetPointN(g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetPointN_r(handle, g, n)
}

func cGEOSGeomGetStartPoint(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetStartPoint_r(handle, g)
}

func cGEOSGeomGetEndPoint(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetEndPoint_r(handle, g)
}

func cGEOSArea(g *C.GEOSGeometry, area *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSArea_r(handle, g, area)
}

func cGEOSLength(g *C.GEOSGeometry, length *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSLength_r(handle, g, length)
}

func cGEOSDistance(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDistance_r(handle, g1, g2, dist)
}

func cGEOSDistanceWithin(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist C.double) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDistanceWithin_r(handle, g1, g2, dist)
}

func cGEOSDistanceIndexed(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSDistanceIndexed_r(handle, g1, g2, dist)
}

func cGEOSHausdorffDistance(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSHausdorffDistance_r(handle, g1, g2, dist)
}

func cGEOSHausdorffDistanceDensify(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, densifyFrac C.double, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSHausdorffDistanceDensify_r(handle, g1, g2, densifyFrac, dist)
}

func cGEOSFrechetDistance(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSFrechetDistance_r(handle, g1, g2, dist)
}

func cGEOSFrechetDistanceDensify(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, densifyFrac C.double, dist *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSFrechetDistanceDensify_r(handle, g1, g2, densifyFrac, dist)
}

func cGEOSGeomGetLength(g *C.GEOSGeometry, length *C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomGetLength_r(handle, g, length)
}

func cGEOSNearestPoints(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSCoordSequence {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSNearestPoints_r(handle, g1, g2)
}

func cGEOSOrientationIndex(Ax C.double, Ay C.double, Bx C.double, By C.double, Px C.double, Py C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSOrientationIndex_r(handle, Ax, Ay, Bx, By, Px, Py)
}

func cGEOSWKTReader_create() *C.GEOSWKTReader {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTReader_create_r(handle)
}

func cGEOSWKTReader_destroy(reader *C.GEOSWKTReader) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTReader_destroy_r(handle, reader)
}

func cGEOSWKTReader_read(reader *C.GEOSWKTReader, wkt *C.char) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTReader_read_r(handle, reader, wkt)
}

func cGEOSWKTWriter_create() *C.GEOSWKTWriter {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTWriter_create_r(handle)
}

func cGEOSWKTWriter_destroy(writer *C.GEOSWKTWriter) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_destroy_r(handle, writer)
}

func cGEOSWKTWriter_write(writer *C.GEOSWKTWriter, g *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTWriter_write_r(handle, writer, g)
}

func cGEOSWKTWriter_setTrim(writer *C.GEOSWKTWriter, trim C.char) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_setTrim_r(handle, writer, trim)
}

func cGEOSWKTWriter_setRoundingPrecision(writer *C.GEOSWKTWriter, precision C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_setRoundingPrecision_r(handle, writer, precision)
}

func cGEOSWKTWriter_setOutputDimension(writer *C.GEOSWKTWriter, dim C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_setOutputDimension_r(handle, writer, dim)
}

func cGEOSWKTWriter_getOutputDimension(writer *C.GEOSWKTWriter) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKTWriter_getOutputDimension_r(handle, writer)
}

func cGEOSWKTWriter_setOld3D(writer *C.GEOSWKTWriter, useOld3D C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKTWriter_setOld3D_r(handle, writer, useOld3D)
}

func cGEOSWKBReader_create() *C.GEOSWKBReader {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBReader_create_r(handle)
}

func cGEOSWKBReader_destroy(reader *C.GEOSWKBReader) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBReader_destroy_r(handle, reader)
}

func cGEOSWKBReader_read(reader *C.GEOSWKBReader, wkb *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBReader_read_r(handle, reader, wkb, size)
}

func cGEOSWKBReader_readHEX(reader *C.GEOSWKBReader, hex *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBReader_readHEX_r(handle, reader, hex, size)
}

func cGEOSWKBWriter_create() *C.GEOSWKBWriter {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_create_r(handle)
}

func cGEOSWKBWriter_destroy(writer *C.GEOSWKBWriter) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBWriter_destroy_r(handle, writer)
}

func cGEOSWKBWriter_write(writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_write_r(handle, writer, g, size)
}

func cGEOSWKBWriter_writeHEX(writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_writeHEX_r(handle, writer, g, size)
}

func cGEOSWKBWriter_getOutputDimension(writer *C.GEOSWKBWriter) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_getOutputDimension_r(handle, writer)
}

func cGEOSWKBWriter_setOutputDimension(writer *C.GEOSWKBWriter, newDimension C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBWriter_setOutputDimension_r(handle, writer, newDimension)
}

func cGEOSWKBWriter_getByteOrder(writer *C.GEOSWKBWriter) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_getByteOrder_r(handle, writer)
}

func cGEOSWKBWriter_setByteOrder(writer *C.GEOSWKBWriter, byteOrder C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBWriter_setByteOrder_r(handle, writer, byteOrder)
}

func cGEOSWKBWriter_getFlavor(writer *C.GEOSWKBWriter) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_getFlavor_r(handle, writer)
}

func cGEOSWKBWriter_setFlavor(writer *C.GEOSWKBWriter, flavor C.int) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBWriter_setFlavor_r(handle, writer, flavor)
}

func cGEOSWKBWriter_getIncludeSRID(writer *C.GEOSWKBWriter) C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSWKBWriter_getIncludeSRID_r(handle, writer)
}

func cGEOSWKBWriter_setIncludeSRID(writer *C.GEOSWKBWriter, writeSRID C.char) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSWKBWriter_setIncludeSRID_r(handle, writer, writeSRID)
}

func cGEOSFree(buffer *C.void) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSFree_r(handle, unsafe.Pointer(buffer))
}

func cGEOSGeoJSONReader_create() *C.GEOSGeoJSONReader {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeoJSONReader_create_r(handle)
}

func cGEOSGeoJSONReader_destroy(reader *C.GEOSGeoJSONReader) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSGeoJSONReader_destroy_r(handle, reader)
}

func cGEOSGeoJSONReader_readGeometry(reader *C.GEOSGeoJSONReader, geojson *C.char) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeoJSONReader_readGeometry_r(handle, reader, geojson)
}

func cGEOSGeoJSONWriter_create() *C.GEOSGeoJSONWriter {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeoJSONWriter_create_r(handle)
}

func cGEOSGeoJSONWriter_destroy(writer *C.GEOSGeoJSONWriter) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSGeoJSONWriter_destroy_r(handle, writer)
}

func cGEOSGeoJSONWriter_writeGeometry(writer *C.GEOSGeoJSONWriter, g *C.GEOSGeometry, indent C.int) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeoJSONWriter_writeGeometry_r(handle, writer, g, indent)
}

func cGEOSSingleSidedBuffer(g *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double, leftSide C.int) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSSingleSidedBuffer_r(handle, g, width, quadsegs, joinStyle, mitreLimit, leftSide)
}

func cGEOSGeomFromWKT(wkt *C.char) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomFromWKT_r(handle, wkt)
}

func cGEOSGeomToWKT(g *C.GEOSGeometry) *C.char {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomToWKT_r(handle, g)
}

func cGEOS_getWKBOutputDims() C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_getWKBOutputDims_r(handle)
}

func cGEOS_setWKBOutputDims(newDims C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_setWKBOutputDims_r(handle, newDims)
}

func cGEOS_getWKBByteOrder() C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_getWKBByteOrder_r(handle)
}

func cGEOS_setWKBByteOrder(byteOrder C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_setWKBByteOrder_r(handle, byteOrder)
}

func cGEOSGeomFromWKB_buf(wkb *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomFromWKB_buf_r(handle, wkb, size)
}

func cGEOSGeomToWKB_buf(g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomToWKB_buf_r(handle, g, size)
}

func cGEOSGeomFromHEX_buf(hex *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomFromHEX_buf_r(handle, hex, size)
}

func cGEOSGeomToHEX_buf(g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomToHEX_buf_r(handle, g, size)
}

func cGEOSUnionCascaded(g *C.GEOSGeometry) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSUnionCascaded_r(handle, g)
}
