/*
 * Collection utility of Point Struct
 *
 * Generated by: Go Streamer
 */

package cios

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type PointStream []Point

func PointStreamOf(arg ...Point) PointStream {
	return arg
}
func PointStreamFrom(arg []Point) PointStream {
	return arg
}
func CreatePointStream(arg ...Point) *PointStream {
	tmp := PointStreamOf(arg...)
	return &tmp
}
func GeneratePointStream(arg []Point) *PointStream {
	tmp := PointStreamFrom(arg)
	return &tmp
}

func (self *PointStream) Add(arg Point) *PointStream {
	return self.AddAll(arg)
}
func (self *PointStream) AddAll(arg ...Point) *PointStream {
	*self = append(*self, arg...)
	return self
}
func (self *PointStream) AddSafe(arg *Point) *PointStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *PointStream) Aggregate(fn func(Point, Point) Point) *PointStream {
	result := PointStreamOf()
	self.ForEach(func(v Point, i int) {
		if i == 0 {
			result.Add(fn(Point{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *PointStream) AllMatch(fn func(Point, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *PointStream) AnyMatch(fn func(Point, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *PointStream) Clone() *PointStream {
	temp := make([]Point, self.Len())
	copy(temp, *self)
	return (*PointStream)(&temp)
}
func (self *PointStream) Copy() *PointStream {
	return self.Clone()
}
func (self *PointStream) Concat(arg []Point) *PointStream {
	return self.AddAll(arg...)
}
func (self *PointStream) Contains(arg Point) bool {
	return self.FindIndex(func(_arg Point, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *PointStream) Clean() *PointStream {
	*self = PointStreamOf()
	return self
}
func (self *PointStream) Delete(index int) *PointStream {
	return self.DeleteRange(index, index)
}
func (self *PointStream) DeleteRange(startIndex, endIndex int) *PointStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *PointStream) Distinct() *PointStream {
	caches := map[string]bool{}
	result := PointStreamOf()
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if !f {
				result = append(result, v)
			}
		} else if caches[key] = true; !f {
			result = append(result, v)
		}

	}
	*self = result
	return self
}
func (self *PointStream) Each(fn func(Point)) *PointStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *PointStream) EachRight(fn func(Point)) *PointStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *PointStream) Equals(arr []Point) bool {
	if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
		return false
	}
	for i := range *self {
		if !reflect.DeepEqual((*self)[i], arr[i]) {
			return false
		}
	}
	return true
}
func (self *PointStream) Filter(fn func(Point, int) bool) *PointStream {
	result := PointStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *PointStream) FilterSlim(fn func(Point, int) bool) *PointStream {
	result := PointStreamOf()
	caches := map[string]bool{}
	for i, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v, i); caches[key] {
			result.Add(v)

		}
	}
	*self = result
	return self
}
func (self *PointStream) Find(fn func(Point, int) bool) *Point {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *PointStream) FindOr(fn func(Point, int) bool, or Point) Point {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *PointStream) FindIndex(fn func(Point, int) bool) int {
	if self == nil {
		return -1
	}
	for i, v := range *self {
		if fn(v, i) {
			return i
		}
	}
	return -1
}
func (self *PointStream) First() *Point {
	return self.Get(0)
}
func (self *PointStream) FirstOr(arg Point) Point {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *PointStream) ForEach(fn func(Point, int)) *PointStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *PointStream) ForEachRight(fn func(Point, int)) *PointStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *PointStream) GroupBy(fn func(Point, int) string) map[string][]Point {
	m := map[string][]Point{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *PointStream) GroupByValues(fn func(Point, int) string) [][]Point {
	var tmp [][]Point
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *PointStream) IndexOf(arg Point) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *PointStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *PointStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *PointStream) Last() *Point {
	return self.Get(self.Len() - 1)
}
func (self *PointStream) LastOr(arg Point) Point {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *PointStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *PointStream) Limit(limit int) *PointStream {
	self.Slice(0, limit)
	return self
}

func (self *PointStream) Map(fn func(Point, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PointStream) Map2Int(fn func(Point, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PointStream) Map2Int32(fn func(Point, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PointStream) Map2Int64(fn func(Point, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PointStream) Map2Float32(fn func(Point, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PointStream) Map2Float64(fn func(Point, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PointStream) Map2Bool(fn func(Point, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PointStream) Map2Bytes(fn func(Point, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PointStream) Map2String(fn func(Point, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PointStream) Max(fn func(Point, int) float64) *Point {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Max(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *PointStream) Min(fn func(Point, int) float64) *Point {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Min(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *PointStream) NoneMatch(fn func(Point, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *PointStream) Get(index int) *Point {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *PointStream) GetOr(index int, arg Point) Point {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *PointStream) Peek(fn func(*Point, int)) *PointStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *PointStream) Reduce(fn func(Point, Point, int) Point) *PointStream {
	return self.ReduceInit(fn, Point{})
}
func (self *PointStream) ReduceInit(fn func(Point, Point, int) Point, initialValue Point) *PointStream {
	result := PointStreamOf()
	self.ForEach(func(v Point, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *PointStream) ReduceInterface(fn func(interface{}, Point, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Point{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *PointStream) ReduceString(fn func(string, Point, int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *PointStream) ReduceInt(fn func(int, Point, int) int) []int {
	result := []int{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *PointStream) ReduceInt32(fn func(int32, Point, int) int32) []int32 {
	result := []int32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *PointStream) ReduceInt64(fn func(int64, Point, int) int64) []int64 {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *PointStream) ReduceFloat32(fn func(float32, Point, int) float32) []float32 {
	result := []float32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *PointStream) ReduceFloat64(fn func(float64, Point, int) float64) []float64 {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *PointStream) ReduceBool(fn func(bool, Point, int) bool) []bool {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *PointStream) Reverse() *PointStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *PointStream) Replace(fn func(Point, int) Point) *PointStream {
	return self.ForEach(func(v Point, i int) { self.Set(i, fn(v, i)) })
}
func (self *PointStream) Select(fn func(Point) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *PointStream) Set(index int, val Point) *PointStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *PointStream) Skip(skip int) *PointStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *PointStream) SkippingEach(fn func(Point, int) int) *PointStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *PointStream) Slice(startIndex, n int) *PointStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Point{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *PointStream) Sort(fn func(i, j int) bool) *PointStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *PointStream) Tail() *Point {
	return self.Last()
}
func (self *PointStream) TailOr(arg Point) Point {
	return self.LastOr(arg)
}
func (self *PointStream) ToList() []Point {
	return self.Val()
}
func (self *PointStream) Unique() *PointStream {
	return self.Distinct()
}
func (self *PointStream) Val() []Point {
	if self == nil {
		return []Point{}
	}
	return *self.Copy()
}
func (self *PointStream) While(fn func(Point, int) bool) *PointStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *PointStream) Where(fn func(Point) bool) *PointStream {
	result := PointStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *PointStream) WhereSlim(fn func(Point) bool) *PointStream {
	result := PointStreamOf()
	caches := map[string]bool{}
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v); caches[key] {
			result.Add(v)
		}
	}
	*self = result
	return self
}
