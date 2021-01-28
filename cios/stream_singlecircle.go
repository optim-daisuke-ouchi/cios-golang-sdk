/*
 * Collection utility of SingleCircle Struct
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

type SingleCircleStream []SingleCircle

func SingleCircleStreamOf(arg ...SingleCircle) SingleCircleStream {
	return arg
}
func SingleCircleStreamFrom(arg []SingleCircle) SingleCircleStream {
	return arg
}
func CreateSingleCircleStream(arg ...SingleCircle) *SingleCircleStream {
	tmp := SingleCircleStreamOf(arg...)
	return &tmp
}
func GenerateSingleCircleStream(arg []SingleCircle) *SingleCircleStream {
	tmp := SingleCircleStreamFrom(arg)
	return &tmp
}

func (self *SingleCircleStream) Add(arg SingleCircle) *SingleCircleStream {
	return self.AddAll(arg)
}
func (self *SingleCircleStream) AddAll(arg ...SingleCircle) *SingleCircleStream {
	*self = append(*self, arg...)
	return self
}
func (self *SingleCircleStream) AddSafe(arg *SingleCircle) *SingleCircleStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SingleCircleStream) Aggregate(fn func(SingleCircle, SingleCircle) SingleCircle) *SingleCircleStream {
	result := SingleCircleStreamOf()
	self.ForEach(func(v SingleCircle, i int) {
		if i == 0 {
			result.Add(fn(SingleCircle{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *SingleCircleStream) AllMatch(fn func(SingleCircle, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SingleCircleStream) AnyMatch(fn func(SingleCircle, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SingleCircleStream) Clone() *SingleCircleStream {
	temp := make([]SingleCircle, self.Len())
	copy(temp, *self)
	return (*SingleCircleStream)(&temp)
}
func (self *SingleCircleStream) Copy() *SingleCircleStream {
	return self.Clone()
}
func (self *SingleCircleStream) Concat(arg []SingleCircle) *SingleCircleStream {
	return self.AddAll(arg...)
}
func (self *SingleCircleStream) Contains(arg SingleCircle) bool {
	return self.FindIndex(func(_arg SingleCircle, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SingleCircleStream) Clean() *SingleCircleStream {
	*self = SingleCircleStreamOf()
	return self
}
func (self *SingleCircleStream) Delete(index int) *SingleCircleStream {
	return self.DeleteRange(index, index)
}
func (self *SingleCircleStream) DeleteRange(startIndex, endIndex int) *SingleCircleStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SingleCircleStream) Distinct() *SingleCircleStream {
	caches := map[string]bool{}
	result := SingleCircleStreamOf()
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
func (self *SingleCircleStream) Each(fn func(SingleCircle)) *SingleCircleStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SingleCircleStream) EachRight(fn func(SingleCircle)) *SingleCircleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SingleCircleStream) Equals(arr []SingleCircle) bool {
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
func (self *SingleCircleStream) Filter(fn func(SingleCircle, int) bool) *SingleCircleStream {
	result := SingleCircleStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleCircleStream) FilterSlim(fn func(SingleCircle, int) bool) *SingleCircleStream {
	result := SingleCircleStreamOf()
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
func (self *SingleCircleStream) Find(fn func(SingleCircle, int) bool) *SingleCircle {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SingleCircleStream) FindOr(fn func(SingleCircle, int) bool, or SingleCircle) SingleCircle {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SingleCircleStream) FindIndex(fn func(SingleCircle, int) bool) int {
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
func (self *SingleCircleStream) First() *SingleCircle {
	return self.Get(0)
}
func (self *SingleCircleStream) FirstOr(arg SingleCircle) SingleCircle {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SingleCircleStream) ForEach(fn func(SingleCircle, int)) *SingleCircleStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SingleCircleStream) ForEachRight(fn func(SingleCircle, int)) *SingleCircleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SingleCircleStream) GroupBy(fn func(SingleCircle, int) string) map[string][]SingleCircle {
	m := map[string][]SingleCircle{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SingleCircleStream) GroupByValues(fn func(SingleCircle, int) string) [][]SingleCircle {
	var tmp [][]SingleCircle
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SingleCircleStream) IndexOf(arg SingleCircle) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SingleCircleStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SingleCircleStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SingleCircleStream) Last() *SingleCircle {
	return self.Get(self.Len() - 1)
}
func (self *SingleCircleStream) LastOr(arg SingleCircle) SingleCircle {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SingleCircleStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SingleCircleStream) Limit(limit int) *SingleCircleStream {
	self.Slice(0, limit)
	return self
}

func (self *SingleCircleStream) Map(fn func(SingleCircle, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleCircleStream) Map2Int(fn func(SingleCircle, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleCircleStream) Map2Int32(fn func(SingleCircle, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleCircleStream) Map2Int64(fn func(SingleCircle, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleCircleStream) Map2Float32(fn func(SingleCircle, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleCircleStream) Map2Float64(fn func(SingleCircle, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleCircleStream) Map2Bool(fn func(SingleCircle, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleCircleStream) Map2Bytes(fn func(SingleCircle, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleCircleStream) Map2String(fn func(SingleCircle, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleCircleStream) Max(fn func(SingleCircle, int) float64) *SingleCircle {
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
func (self *SingleCircleStream) Min(fn func(SingleCircle, int) float64) *SingleCircle {
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
func (self *SingleCircleStream) NoneMatch(fn func(SingleCircle, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SingleCircleStream) Get(index int) *SingleCircle {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SingleCircleStream) GetOr(index int, arg SingleCircle) SingleCircle {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SingleCircleStream) Peek(fn func(*SingleCircle, int)) *SingleCircleStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *SingleCircleStream) Reduce(fn func(SingleCircle, SingleCircle, int) SingleCircle) *SingleCircleStream {
	return self.ReduceInit(fn, SingleCircle{})
}
func (self *SingleCircleStream) ReduceInit(fn func(SingleCircle, SingleCircle, int) SingleCircle, initialValue SingleCircle) *SingleCircleStream {
	result := SingleCircleStreamOf()
	self.ForEach(func(v SingleCircle, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SingleCircleStream) ReduceInterface(fn func(interface{}, SingleCircle, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(SingleCircle{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SingleCircleStream) ReduceString(fn func(string, SingleCircle, int) string) []string {
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
func (self *SingleCircleStream) ReduceInt(fn func(int, SingleCircle, int) int) []int {
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
func (self *SingleCircleStream) ReduceInt32(fn func(int32, SingleCircle, int) int32) []int32 {
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
func (self *SingleCircleStream) ReduceInt64(fn func(int64, SingleCircle, int) int64) []int64 {
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
func (self *SingleCircleStream) ReduceFloat32(fn func(float32, SingleCircle, int) float32) []float32 {
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
func (self *SingleCircleStream) ReduceFloat64(fn func(float64, SingleCircle, int) float64) []float64 {
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
func (self *SingleCircleStream) ReduceBool(fn func(bool, SingleCircle, int) bool) []bool {
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
func (self *SingleCircleStream) Reverse() *SingleCircleStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SingleCircleStream) Replace(fn func(SingleCircle, int) SingleCircle) *SingleCircleStream {
	return self.ForEach(func(v SingleCircle, i int) { self.Set(i, fn(v, i)) })
}
func (self *SingleCircleStream) Select(fn func(SingleCircle) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *SingleCircleStream) Set(index int, val SingleCircle) *SingleCircleStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SingleCircleStream) Skip(skip int) *SingleCircleStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SingleCircleStream) SkippingEach(fn func(SingleCircle, int) int) *SingleCircleStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SingleCircleStream) Slice(startIndex, n int) *SingleCircleStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []SingleCircle{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SingleCircleStream) Sort(fn func(i, j int) bool) *SingleCircleStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SingleCircleStream) Tail() *SingleCircle {
	return self.Last()
}
func (self *SingleCircleStream) TailOr(arg SingleCircle) SingleCircle {
	return self.LastOr(arg)
}
func (self *SingleCircleStream) ToList() []SingleCircle {
	return self.Val()
}
func (self *SingleCircleStream) Unique() *SingleCircleStream {
	return self.Distinct()
}
func (self *SingleCircleStream) Val() []SingleCircle {
	if self == nil {
		return []SingleCircle{}
	}
	return *self.Copy()
}
func (self *SingleCircleStream) While(fn func(SingleCircle, int) bool) *SingleCircleStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *SingleCircleStream) Where(fn func(SingleCircle) bool) *SingleCircleStream {
	result := SingleCircleStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleCircleStream) WhereSlim(fn func(SingleCircle) bool) *SingleCircleStream {
	result := SingleCircleStreamOf()
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
