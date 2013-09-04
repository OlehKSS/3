package engine

import (
	"code.google.com/p/mx3/cuda"
	"code.google.com/p/mx3/data"
	"code.google.com/p/mx3/util"
	"reflect"
)

// A buffered quantity is stored in GPU memory at all times.
// E.g.: magnetization.
type bufferedQuant struct {
	info
	buffer *data.Slice
}

// constructor
func buffered(slice *data.Slice, name, unit string) bufferedQuant {
	return bufferedQuant{mkInfo(slice.NComp(), name, unit, slice.Mesh()), slice}
}

// get buffer (on GPU, no need to recycle)
func (b *bufferedQuant) Get() (q *data.Slice, recycle bool) {
	return b.buffer, false
}

// Replace the data by src. Auto rescales if needed.
func (b *bufferedQuant) Set(src *data.Slice) {
	if src.Mesh().Size() != b.buffer.Mesh().Size() {
		src = data.Resample(src, b.buffer.Mesh().Size())
	}
	data.Copy(b.buffer, src)
}

// Set the value of one cell.
func (b *bufferedQuant) SetCell(ix, iy, iz int, v ...float64) {
	nComp := b.NComp()
	util.Argument(len(v) == nComp)
	for c := 0; c < nComp; c++ {
		cuda.SetCell(b.buffer, util.SwapIndex(c, nComp), iz, iy, ix, float32(v[c]))
	}
}

// Get the value of one cell
//func (b *bufferedQuant) GetCell(comp, ix, iy, iz int) float64 {
//	return float64(cuda.GetCell(b.buffer, util.SwapIndex(comp, b.NComp()), iz, iy, ix))
//}

// Shift the data over (shx, shy, shz cells), clamping boundary values.
// Typically used in a PostStep function to center the magnetization on
// the simulation window.
func (b *bufferedQuant) Shift(shx, shy, shz int) {
	m2 := cuda.GetBuffer(1, b.buffer.Mesh())
	defer cuda.RecycleBuffer(m2)
	for c := 0; c < b.NComp(); c++ {
		comp := b.buffer.Comp(c)
		cuda.Shift(m2, comp, [3]int{shz, shy, shx}) // ZYX !
		data.Copy(comp, m2)
	}
}

func (b *bufferedQuant) GetVec() []float64       { return Average(b) }
func (b *bufferedQuant) SetValue(v interface{})  { b.Set(v.(*data.Slice)) }
func (b *bufferedQuant) Eval() interface{}       { return b }
func (b *bufferedQuant) InputType() reflect.Type { return reflect.TypeOf(new(data.Slice)) }
