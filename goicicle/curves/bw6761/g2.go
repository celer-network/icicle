// Copyright 2023 Ingonyama
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by Ingonyama DO NOT EDIT

package bw6761

import (
	"encoding/binary"
	"unsafe"
)

// #cgo CFLAGS: -I./include/
// #cgo CFLAGS: -I/usr/local/cuda/include
// #cgo LDFLAGS: -L${SRCDIR}/../../ -lbw6761
// #include "projective.h"
// #include "ve_mod_mult.h"
import "C"

// G2 extension field

type G2Element [12]uint64

type G2PointAffine struct {
	X, Y G2Element
}

type G2Point struct {
	X, Y, Z G2Element
}

func (p *G2Point) Random() *G2Point {
	outC := (*C.BW6761_g2_projective_t)(unsafe.Pointer(p))
	C.random_g2_projective_bw6_761(outC)

	return p
}

func (p *G2Point) FromAffine(affine *G2PointAffine) *G2Point {
	out := (*C.BW6761_g2_projective_t)(unsafe.Pointer(p))
	in := (*C.BW6761_g2_affine_t)(unsafe.Pointer(affine))

	C.g2_projective_from_affine_bw6_761(out, in)

	return p
}

func (p *G2Point) Eq(pCompare *G2Point) bool {
	// Cast *PointBW6761 to *C.BW6761_projective_t
	// The unsafe.Pointer cast is necessary because Go doesn't allow direct casts
	// between different pointer types.
	// It's your responsibility to ensure that the types are compatible.
	pC := (*C.BW6761_g2_projective_t)(unsafe.Pointer(p))
	pCompareC := (*C.BW6761_g2_projective_t)(unsafe.Pointer(pCompare))

	// Call the C function
	// The C function doesn't keep any references to the data,
	// so it's fine if the Go garbage collector moves or deletes the data later.
	return bool(C.eq_g2_bw6_761(pC, pCompareC))
}

func (f *G2Element) ToBytesLe() []byte {
	var bytes []byte
	for _, val := range f {
		buf := make([]byte, 8) // 8 bytes because uint64 is 64-bit
		binary.LittleEndian.PutUint64(buf, val)
		bytes = append(bytes, buf...)
	}
	return bytes
}

func (p *G2PointAffine) FromProjective(projective *G2Point) *G2PointAffine {
	out := (*C.BW6761_g2_affine_t)(unsafe.Pointer(p))
	in := (*C.BW6761_g2_projective_t)(unsafe.Pointer(projective))

	C.g2_projective_to_affine_bw6_761(out, in)

	return p
}

func (p *G2Point) IsOnCurve() bool {
	// Directly copy memory from the C struct to the Go struct
	point := (*C.BW6761_g2_projective_t)(unsafe.Pointer(p))
	res := C.g2_projective_is_on_curve_bw6_761(point)

	return bool(res)
}
