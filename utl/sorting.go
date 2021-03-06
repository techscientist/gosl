// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utl

import (
	"sort"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

// IntSort3 sorts 3 values in ascending order
func IntSort3(a, b, c *int) {
	if *b < *a {
		*a, *b = *b, *a
	}
	if *c < *b {
		*b, *c = *c, *b
	}
	if *b < *a {
		*a, *b = *b, *a
	}
}

// IntSort4 sort four values in ascending order
func IntSort4(a, b, c, d *int) {
	if *b < *a {
		*b, *a = *a, *b
	}
	if *c < *b {
		*c, *b = *b, *c
	}
	if *d < *c {
		*d, *c = *c, *d
	}
	if *b < *a {
		*b, *a = *a, *b
	}
	if *c < *b {
		*c, *b = *b, *c
	}
	if *b < *a {
		*b, *a = *a, *b
	}
}

// DblSort3 sorts 3 values in ascending order
func DblSort3(a, b, c *float64) {
	if *b < *a {
		*a, *b = *b, *a
	}
	if *c < *b {
		*b, *c = *c, *b
	}
	if *b < *a {
		*a, *b = *b, *a
	}
}

// DblDsort3 sorts 3 values in descending order
func DblDsort3(a, b, c *float64) {
	if *b > *a {
		*a, *b = *b, *a
	}
	if *c > *b {
		*b, *c = *c, *b
	}
	if *b > *a {
		*a, *b = *b, *a
	}
}

// DblSort4 sort four values in ascending order
func DblSort4(a, b, c, d *float64) {
	if *b < *a {
		*b, *a = *a, *b
	}
	if *c < *b {
		*c, *b = *b, *c
	}
	if *d < *c {
		*d, *c = *c, *d
	}
	if *b < *a {
		*b, *a = *a, *b
	}
	if *c < *b {
		*c, *b = *b, *c
	}
	if *b < *a {
		*b, *a = *a, *b
	}
}

// DblGetSorted returns a sorted (increasing) copy of 'A'
func DblGetSorted(A []float64) (Asorted []float64) {
	Asorted = make([]float64, len(A))
	copy(Asorted, A)
	sort.Float64s(Asorted)
	return
}

// Quadruple helps to sort a quadruple of 1 int and 3 float64s
type Quadruple struct {
	I int
	X float64
	Y float64
	Z float64
}

// Quadruples helps to sort quadruples
type Quadruples []*Quadruple

// Init initialise Quadruples with i, x, y, and z
//  Note: i, x, y, or z can be nil; but at least one of them must be non nil
func BuildQuadruples(i []int, x, y, z []float64) (q Quadruples) {
	ni := len(i)
	nx := len(x)
	ny := len(y)
	nz := len(z)
	_, n := IntMinMax([]int{ni, nx, ny, nz})
	q = make([]*Quadruple, n)
	var ival int
	var xval, yval, zval float64
	for k := 0; k < n; k++ {
		if ni > k {
			ival = i[k]
		}
		if nx > k {
			xval = x[k]
		}
		if ny > k {
			yval = y[k]
		}
		if nz > k {
			zval = z[k]
		}
		q[k] = &Quadruple{ival, xval, yval, zval}
	}
	return
}

// I returns the 'i' in quadruples
func (o Quadruples) I() (i []int) {
	i = make([]int, len(o))
	for k := 0; k < len(o); k++ {
		i[k] = o[k].I
	}
	return
}

// X returns the 'x' in quadruples
func (o Quadruples) X() (x []float64) {
	x = make([]float64, len(o))
	for k := 0; k < len(o); k++ {
		x[k] = o[k].X
	}
	return
}

// Y returns the 'y' in quadruples
func (o Quadruples) Y() (y []float64) {
	y = make([]float64, len(o))
	for k := 0; k < len(o); k++ {
		y[k] = o[k].Y
	}
	return
}

// Z returns the 'z' in quadruples
func (o Quadruples) Z() (z []float64) {
	z = make([]float64, len(o))
	for k := 0; k < len(o); k++ {
		z[k] = o[k].Z
	}
	return
}

// Len returns the length of Quadruples
func (o Quadruples) Len() int { return len(o) }

// Swap swaps two quadruples
func (o Quadruples) Swap(i, j int) { o[i], o[j] = o[j], o[i] }

// String returns the string representation of Quadruples
func (o Quadruples) String() string {
	res := ""
	for _, p := range o {
		res += io.Sf("%6d%20g%20g%20g\n", p.I, p.X, p.Y, p.Z)
	}
	return res
}

// Sort Quadruples by I
type QuadruplesByI struct{ Quadruples }

func (o QuadruplesByI) Less(i, j int) bool { return o.Quadruples[i].I < o.Quadruples[j].I }

// Sort Quadruples by X
type QuadruplesByX struct{ Quadruples }

func (o QuadruplesByX) Less(i, j int) bool { return o.Quadruples[i].X < o.Quadruples[j].X }

// Sort Quadruples by Y
type QuadruplesByY struct{ Quadruples }

func (o QuadruplesByY) Less(i, j int) bool { return o.Quadruples[i].Y < o.Quadruples[j].Y }

// Sort Quadruples by Z
type QuadruplesByZ struct{ Quadruples }

func (o QuadruplesByZ) Less(i, j int) bool { return o.Quadruples[i].Z < o.Quadruples[j].Z }

// SortQuadruples sorts i, x, y, and z by "i", "x", "y", or "z"
//  Note: either i, x, y, or z can be nil; i.e. at least one of them must be non nil
func SortQuadruples(i []int, x, y, z []float64, by string) (I []int, X, Y, Z []float64, err error) {
	q := BuildQuadruples(i, x, y, z)
	switch by {
	case "i":
		sort.Sort(QuadruplesByI{q})
	case "x":
		sort.Sort(QuadruplesByX{q})
	case "y":
		sort.Sort(QuadruplesByY{q})
	case "z":
		sort.Sort(QuadruplesByZ{q})
	default:
		return nil, nil, nil, nil, chk.Err("sort quadruples command must be 'i', 'x', 'y', or 'z'. by == '%s' is invalid", by)
	}
	if i != nil {
		I = q.I()
	}
	if x != nil {
		X = q.X()
	}
	if y != nil {
		Y = q.Y()
	}
	if z != nil {
		Z = q.Z()
	}
	return
}

// Str => ??? maps /////////////////////////////////////////////////////////////////////////////////

func StrIntMapSort(m map[string]int) (sorted_keys []string) {
	sorted_keys = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sorted_keys[i] = key
		i++
	}
	sort.Strings(sorted_keys)
	return
}

func StrDblMapSort(m map[string]float64) (sorted_keys []string) {
	sorted_keys = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sorted_keys[i] = key
		i++
	}
	sort.Strings(sorted_keys)
	return
}

func StrBoolMapSort(m map[string]bool) (sorted_keys []string) {
	sorted_keys = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sorted_keys[i] = key
		i++
	}
	sort.Strings(sorted_keys)
	return
}

func StrIntMapSortSplit(m map[string]int) (sorted_keys []string, sorted_vals []int) {
	sorted_keys = make([]string, len(m))
	sorted_vals = make([]int, len(m))
	i := 0
	for key, _ := range m {
		sorted_keys[i] = key
		i++
	}
	sort.Strings(sorted_keys)
	for j, key := range sorted_keys {
		sorted_vals[j] = m[key]
	}
	return
}

func StrDblMapSortSplit(m map[string]float64) (sorted_keys []string, sorted_vals []float64) {
	sorted_keys = make([]string, len(m))
	sorted_vals = make([]float64, len(m))
	i := 0
	for key, _ := range m {
		sorted_keys[i] = key
		i++
	}
	sort.Strings(sorted_keys)
	for j, key := range sorted_keys {
		sorted_vals[j] = m[key]
	}
	return
}

func StrBoolMapSortSplit(m map[string]bool) (sorted_keys []string, sorted_vals []bool) {
	sorted_keys = make([]string, len(m))
	sorted_vals = make([]bool, len(m))
	i := 0
	for key, _ := range m {
		sorted_keys[i] = key
		i++
	}
	sort.Strings(sorted_keys)
	for j, key := range sorted_keys {
		sorted_vals[j] = m[key]
	}
	return
}

// Int => ??? maps /////////////////////////////////////////////////////////////////////////////////

func IntBoolMapSort(m map[int]bool) (sorted_keys []int) {
	sorted_keys = make([]int, len(m))
	i := 0
	for key, _ := range m {
		sorted_keys[i] = key
		i++
	}
	sort.Ints(sorted_keys)
	return
}
