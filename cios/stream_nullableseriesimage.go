/*
 * Collection utility of NullableSeriesImage Struct
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

type NullableSeriesImageStream []NullableSeriesImage

func NullableSeriesImageStreamOf(arg ...NullableSeriesImage) NullableSeriesImageStream {
	return arg
}
func NullableSeriesImageStreamFrom(arg []NullableSeriesImage) NullableSeriesImageStream {
	return arg
}
func CreateNullableSeriesImageStream(arg ...NullableSeriesImage) *NullableSeriesImageStream {
	tmp := NullableSeriesImageStreamOf(arg...)
	return &tmp
}
func GenerateNullableSeriesImageStream(arg []NullableSeriesImage) *NullableSeriesImageStream {
	tmp := NullableSeriesImageStreamFrom(arg)
	return &tmp
}

func (self *NullableSeriesImageStream) Add(arg NullableSeriesImage) *NullableSeriesImageStream {
	return self.AddAll(arg)
}
func (self *NullableSeriesImageStream) AddAll(arg ...NullableSeriesImage) *NullableSeriesImageStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSeriesImageStream) AddSafe(arg *NullableSeriesImage) *NullableSeriesImageStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSeriesImageStream) Aggregate(fn func(NullableSeriesImage, NullableSeriesImage) NullableSeriesImage) *NullableSeriesImageStream {
	result := NullableSeriesImageStreamOf()
	self.ForEach(func(v NullableSeriesImage, i int) {
		if i == 0 {
			result.Add(fn(NullableSeriesImage{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesImageStream) AllMatch(fn func(NullableSeriesImage, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSeriesImageStream) AnyMatch(fn func(NullableSeriesImage, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSeriesImageStream) Clone() *NullableSeriesImageStream {
	temp := make([]NullableSeriesImage, self.Len())
	copy(temp, *self)
	return (*NullableSeriesImageStream)(&temp)
}
func (self *NullableSeriesImageStream) Copy() *NullableSeriesImageStream {
	return self.Clone()
}
func (self *NullableSeriesImageStream) Concat(arg []NullableSeriesImage) *NullableSeriesImageStream {
	return self.AddAll(arg...)
}
func (self *NullableSeriesImageStream) Contains(arg NullableSeriesImage) bool {
	return self.FindIndex(func(_arg NullableSeriesImage, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSeriesImageStream) Clean() *NullableSeriesImageStream {
	*self = NullableSeriesImageStreamOf()
	return self
}
func (self *NullableSeriesImageStream) Delete(index int) *NullableSeriesImageStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSeriesImageStream) DeleteRange(startIndex, endIndex int) *NullableSeriesImageStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSeriesImageStream) Distinct() *NullableSeriesImageStream {
	caches := map[string]bool{}
	result := NullableSeriesImageStreamOf()
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
func (self *NullableSeriesImageStream) Each(fn func(NullableSeriesImage)) *NullableSeriesImageStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSeriesImageStream) EachRight(fn func(NullableSeriesImage)) *NullableSeriesImageStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSeriesImageStream) Equals(arr []NullableSeriesImage) bool {
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
func (self *NullableSeriesImageStream) Filter(fn func(NullableSeriesImage, int) bool) *NullableSeriesImageStream {
	result := NullableSeriesImageStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesImageStream) FilterSlim(fn func(NullableSeriesImage, int) bool) *NullableSeriesImageStream {
	result := NullableSeriesImageStreamOf()
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
func (self *NullableSeriesImageStream) Find(fn func(NullableSeriesImage, int) bool) *NullableSeriesImage {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesImageStream) FindOr(fn func(NullableSeriesImage, int) bool, or NullableSeriesImage) NullableSeriesImage {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSeriesImageStream) FindIndex(fn func(NullableSeriesImage, int) bool) int {
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
func (self *NullableSeriesImageStream) First() *NullableSeriesImage {
	return self.Get(0)
}
func (self *NullableSeriesImageStream) FirstOr(arg NullableSeriesImage) NullableSeriesImage {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesImageStream) ForEach(fn func(NullableSeriesImage, int)) *NullableSeriesImageStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSeriesImageStream) ForEachRight(fn func(NullableSeriesImage, int)) *NullableSeriesImageStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSeriesImageStream) GroupBy(fn func(NullableSeriesImage, int) string) map[string][]NullableSeriesImage {
	m := map[string][]NullableSeriesImage{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSeriesImageStream) GroupByValues(fn func(NullableSeriesImage, int) string) [][]NullableSeriesImage {
	var tmp [][]NullableSeriesImage
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSeriesImageStream) IndexOf(arg NullableSeriesImage) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSeriesImageStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSeriesImageStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSeriesImageStream) Last() *NullableSeriesImage {
	return self.Get(self.Len() - 1)
}
func (self *NullableSeriesImageStream) LastOr(arg NullableSeriesImage) NullableSeriesImage {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesImageStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSeriesImageStream) Limit(limit int) *NullableSeriesImageStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSeriesImageStream) Map(fn func(NullableSeriesImage, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesImageStream) Map2Int(fn func(NullableSeriesImage, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesImageStream) Map2Int32(fn func(NullableSeriesImage, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesImageStream) Map2Int64(fn func(NullableSeriesImage, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesImageStream) Map2Float32(fn func(NullableSeriesImage, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesImageStream) Map2Float64(fn func(NullableSeriesImage, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesImageStream) Map2Bool(fn func(NullableSeriesImage, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesImageStream) Map2Bytes(fn func(NullableSeriesImage, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesImageStream) Map2String(fn func(NullableSeriesImage, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesImageStream) Max(fn func(NullableSeriesImage, int) float64) *NullableSeriesImage {
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
func (self *NullableSeriesImageStream) Min(fn func(NullableSeriesImage, int) float64) *NullableSeriesImage {
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
func (self *NullableSeriesImageStream) NoneMatch(fn func(NullableSeriesImage, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSeriesImageStream) Get(index int) *NullableSeriesImage {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesImageStream) GetOr(index int, arg NullableSeriesImage) NullableSeriesImage {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesImageStream) Peek(fn func(*NullableSeriesImage, int)) *NullableSeriesImageStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSeriesImageStream) Reduce(fn func(NullableSeriesImage, NullableSeriesImage, int) NullableSeriesImage) *NullableSeriesImageStream {
	return self.ReduceInit(fn, NullableSeriesImage{})
}
func (self *NullableSeriesImageStream) ReduceInit(fn func(NullableSeriesImage, NullableSeriesImage, int) NullableSeriesImage, initialValue NullableSeriesImage) *NullableSeriesImageStream {
	result := NullableSeriesImageStreamOf()
	self.ForEach(func(v NullableSeriesImage, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesImageStream) ReduceInterface(fn func(interface{}, NullableSeriesImage, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSeriesImage{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSeriesImageStream) ReduceString(fn func(string, NullableSeriesImage, int) string) []string {
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
func (self *NullableSeriesImageStream) ReduceInt(fn func(int, NullableSeriesImage, int) int) []int {
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
func (self *NullableSeriesImageStream) ReduceInt32(fn func(int32, NullableSeriesImage, int) int32) []int32 {
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
func (self *NullableSeriesImageStream) ReduceInt64(fn func(int64, NullableSeriesImage, int) int64) []int64 {
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
func (self *NullableSeriesImageStream) ReduceFloat32(fn func(float32, NullableSeriesImage, int) float32) []float32 {
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
func (self *NullableSeriesImageStream) ReduceFloat64(fn func(float64, NullableSeriesImage, int) float64) []float64 {
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
func (self *NullableSeriesImageStream) ReduceBool(fn func(bool, NullableSeriesImage, int) bool) []bool {
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
func (self *NullableSeriesImageStream) Reverse() *NullableSeriesImageStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSeriesImageStream) Replace(fn func(NullableSeriesImage, int) NullableSeriesImage) *NullableSeriesImageStream {
	return self.ForEach(func(v NullableSeriesImage, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSeriesImageStream) Select(fn func(NullableSeriesImage) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSeriesImageStream) Set(index int, val NullableSeriesImage) *NullableSeriesImageStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSeriesImageStream) Skip(skip int) *NullableSeriesImageStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSeriesImageStream) SkippingEach(fn func(NullableSeriesImage, int) int) *NullableSeriesImageStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSeriesImageStream) Slice(startIndex, n int) *NullableSeriesImageStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSeriesImage{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSeriesImageStream) Sort(fn func(i, j int) bool) *NullableSeriesImageStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSeriesImageStream) Tail() *NullableSeriesImage {
	return self.Last()
}
func (self *NullableSeriesImageStream) TailOr(arg NullableSeriesImage) NullableSeriesImage {
	return self.LastOr(arg)
}
func (self *NullableSeriesImageStream) ToList() []NullableSeriesImage {
	return self.Val()
}
func (self *NullableSeriesImageStream) Unique() *NullableSeriesImageStream {
	return self.Distinct()
}
func (self *NullableSeriesImageStream) Val() []NullableSeriesImage {
	if self == nil {
		return []NullableSeriesImage{}
	}
	return *self.Copy()
}
func (self *NullableSeriesImageStream) While(fn func(NullableSeriesImage, int) bool) *NullableSeriesImageStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSeriesImageStream) Where(fn func(NullableSeriesImage) bool) *NullableSeriesImageStream {
	result := NullableSeriesImageStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesImageStream) WhereSlim(fn func(NullableSeriesImage) bool) *NullableSeriesImageStream {
	result := NullableSeriesImageStreamOf()
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
