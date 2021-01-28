/*
 * Collection utility of NullableSeriesDataBulkRequest Struct
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

type NullableSeriesDataBulkRequestStream []NullableSeriesDataBulkRequest

func NullableSeriesDataBulkRequestStreamOf(arg ...NullableSeriesDataBulkRequest) NullableSeriesDataBulkRequestStream {
	return arg
}
func NullableSeriesDataBulkRequestStreamFrom(arg []NullableSeriesDataBulkRequest) NullableSeriesDataBulkRequestStream {
	return arg
}
func CreateNullableSeriesDataBulkRequestStream(arg ...NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	tmp := NullableSeriesDataBulkRequestStreamOf(arg...)
	return &tmp
}
func GenerateNullableSeriesDataBulkRequestStream(arg []NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	tmp := NullableSeriesDataBulkRequestStreamFrom(arg)
	return &tmp
}

func (self *NullableSeriesDataBulkRequestStream) Add(arg NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	return self.AddAll(arg)
}
func (self *NullableSeriesDataBulkRequestStream) AddAll(arg ...NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSeriesDataBulkRequestStream) AddSafe(arg *NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) Aggregate(fn func(NullableSeriesDataBulkRequest, NullableSeriesDataBulkRequest) NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	result := NullableSeriesDataBulkRequestStreamOf()
	self.ForEach(func(v NullableSeriesDataBulkRequest, i int) {
		if i == 0 {
			result.Add(fn(NullableSeriesDataBulkRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesDataBulkRequestStream) AllMatch(fn func(NullableSeriesDataBulkRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSeriesDataBulkRequestStream) AnyMatch(fn func(NullableSeriesDataBulkRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSeriesDataBulkRequestStream) Clone() *NullableSeriesDataBulkRequestStream {
	temp := make([]NullableSeriesDataBulkRequest, self.Len())
	copy(temp, *self)
	return (*NullableSeriesDataBulkRequestStream)(&temp)
}
func (self *NullableSeriesDataBulkRequestStream) Copy() *NullableSeriesDataBulkRequestStream {
	return self.Clone()
}
func (self *NullableSeriesDataBulkRequestStream) Concat(arg []NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	return self.AddAll(arg...)
}
func (self *NullableSeriesDataBulkRequestStream) Contains(arg NullableSeriesDataBulkRequest) bool {
	return self.FindIndex(func(_arg NullableSeriesDataBulkRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSeriesDataBulkRequestStream) Clean() *NullableSeriesDataBulkRequestStream {
	*self = NullableSeriesDataBulkRequestStreamOf()
	return self
}
func (self *NullableSeriesDataBulkRequestStream) Delete(index int) *NullableSeriesDataBulkRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSeriesDataBulkRequestStream) DeleteRange(startIndex, endIndex int) *NullableSeriesDataBulkRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSeriesDataBulkRequestStream) Distinct() *NullableSeriesDataBulkRequestStream {
	caches := map[string]bool{}
	result := NullableSeriesDataBulkRequestStreamOf()
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
func (self *NullableSeriesDataBulkRequestStream) Each(fn func(NullableSeriesDataBulkRequest)) *NullableSeriesDataBulkRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) EachRight(fn func(NullableSeriesDataBulkRequest)) *NullableSeriesDataBulkRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) Equals(arr []NullableSeriesDataBulkRequest) bool {
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
func (self *NullableSeriesDataBulkRequestStream) Filter(fn func(NullableSeriesDataBulkRequest, int) bool) *NullableSeriesDataBulkRequestStream {
	result := NullableSeriesDataBulkRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesDataBulkRequestStream) FilterSlim(fn func(NullableSeriesDataBulkRequest, int) bool) *NullableSeriesDataBulkRequestStream {
	result := NullableSeriesDataBulkRequestStreamOf()
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
func (self *NullableSeriesDataBulkRequestStream) Find(fn func(NullableSeriesDataBulkRequest, int) bool) *NullableSeriesDataBulkRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesDataBulkRequestStream) FindOr(fn func(NullableSeriesDataBulkRequest, int) bool, or NullableSeriesDataBulkRequest) NullableSeriesDataBulkRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSeriesDataBulkRequestStream) FindIndex(fn func(NullableSeriesDataBulkRequest, int) bool) int {
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
func (self *NullableSeriesDataBulkRequestStream) First() *NullableSeriesDataBulkRequest {
	return self.Get(0)
}
func (self *NullableSeriesDataBulkRequestStream) FirstOr(arg NullableSeriesDataBulkRequest) NullableSeriesDataBulkRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesDataBulkRequestStream) ForEach(fn func(NullableSeriesDataBulkRequest, int)) *NullableSeriesDataBulkRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) ForEachRight(fn func(NullableSeriesDataBulkRequest, int)) *NullableSeriesDataBulkRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) GroupBy(fn func(NullableSeriesDataBulkRequest, int) string) map[string][]NullableSeriesDataBulkRequest {
	m := map[string][]NullableSeriesDataBulkRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSeriesDataBulkRequestStream) GroupByValues(fn func(NullableSeriesDataBulkRequest, int) string) [][]NullableSeriesDataBulkRequest {
	var tmp [][]NullableSeriesDataBulkRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSeriesDataBulkRequestStream) IndexOf(arg NullableSeriesDataBulkRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSeriesDataBulkRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSeriesDataBulkRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSeriesDataBulkRequestStream) Last() *NullableSeriesDataBulkRequest {
	return self.Get(self.Len() - 1)
}
func (self *NullableSeriesDataBulkRequestStream) LastOr(arg NullableSeriesDataBulkRequest) NullableSeriesDataBulkRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesDataBulkRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSeriesDataBulkRequestStream) Limit(limit int) *NullableSeriesDataBulkRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSeriesDataBulkRequestStream) Map(fn func(NullableSeriesDataBulkRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Map2Int(fn func(NullableSeriesDataBulkRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Map2Int32(fn func(NullableSeriesDataBulkRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Map2Int64(fn func(NullableSeriesDataBulkRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Map2Float32(fn func(NullableSeriesDataBulkRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Map2Float64(fn func(NullableSeriesDataBulkRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Map2Bool(fn func(NullableSeriesDataBulkRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Map2Bytes(fn func(NullableSeriesDataBulkRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Map2String(fn func(NullableSeriesDataBulkRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Max(fn func(NullableSeriesDataBulkRequest, int) float64) *NullableSeriesDataBulkRequest {
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
func (self *NullableSeriesDataBulkRequestStream) Min(fn func(NullableSeriesDataBulkRequest, int) float64) *NullableSeriesDataBulkRequest {
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
func (self *NullableSeriesDataBulkRequestStream) NoneMatch(fn func(NullableSeriesDataBulkRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSeriesDataBulkRequestStream) Get(index int) *NullableSeriesDataBulkRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesDataBulkRequestStream) GetOr(index int, arg NullableSeriesDataBulkRequest) NullableSeriesDataBulkRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesDataBulkRequestStream) Peek(fn func(*NullableSeriesDataBulkRequest, int)) *NullableSeriesDataBulkRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSeriesDataBulkRequestStream) Reduce(fn func(NullableSeriesDataBulkRequest, NullableSeriesDataBulkRequest, int) NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	return self.ReduceInit(fn, NullableSeriesDataBulkRequest{})
}
func (self *NullableSeriesDataBulkRequestStream) ReduceInit(fn func(NullableSeriesDataBulkRequest, NullableSeriesDataBulkRequest, int) NullableSeriesDataBulkRequest, initialValue NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	result := NullableSeriesDataBulkRequestStreamOf()
	self.ForEach(func(v NullableSeriesDataBulkRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesDataBulkRequestStream) ReduceInterface(fn func(interface{}, NullableSeriesDataBulkRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSeriesDataBulkRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSeriesDataBulkRequestStream) ReduceString(fn func(string, NullableSeriesDataBulkRequest, int) string) []string {
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
func (self *NullableSeriesDataBulkRequestStream) ReduceInt(fn func(int, NullableSeriesDataBulkRequest, int) int) []int {
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
func (self *NullableSeriesDataBulkRequestStream) ReduceInt32(fn func(int32, NullableSeriesDataBulkRequest, int) int32) []int32 {
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
func (self *NullableSeriesDataBulkRequestStream) ReduceInt64(fn func(int64, NullableSeriesDataBulkRequest, int) int64) []int64 {
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
func (self *NullableSeriesDataBulkRequestStream) ReduceFloat32(fn func(float32, NullableSeriesDataBulkRequest, int) float32) []float32 {
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
func (self *NullableSeriesDataBulkRequestStream) ReduceFloat64(fn func(float64, NullableSeriesDataBulkRequest, int) float64) []float64 {
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
func (self *NullableSeriesDataBulkRequestStream) ReduceBool(fn func(bool, NullableSeriesDataBulkRequest, int) bool) []bool {
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
func (self *NullableSeriesDataBulkRequestStream) Reverse() *NullableSeriesDataBulkRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) Replace(fn func(NullableSeriesDataBulkRequest, int) NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	return self.ForEach(func(v NullableSeriesDataBulkRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSeriesDataBulkRequestStream) Select(fn func(NullableSeriesDataBulkRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSeriesDataBulkRequestStream) Set(index int, val NullableSeriesDataBulkRequest) *NullableSeriesDataBulkRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) Skip(skip int) *NullableSeriesDataBulkRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSeriesDataBulkRequestStream) SkippingEach(fn func(NullableSeriesDataBulkRequest, int) int) *NullableSeriesDataBulkRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) Slice(startIndex, n int) *NullableSeriesDataBulkRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSeriesDataBulkRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) Sort(fn func(i, j int) bool) *NullableSeriesDataBulkRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSeriesDataBulkRequestStream) Tail() *NullableSeriesDataBulkRequest {
	return self.Last()
}
func (self *NullableSeriesDataBulkRequestStream) TailOr(arg NullableSeriesDataBulkRequest) NullableSeriesDataBulkRequest {
	return self.LastOr(arg)
}
func (self *NullableSeriesDataBulkRequestStream) ToList() []NullableSeriesDataBulkRequest {
	return self.Val()
}
func (self *NullableSeriesDataBulkRequestStream) Unique() *NullableSeriesDataBulkRequestStream {
	return self.Distinct()
}
func (self *NullableSeriesDataBulkRequestStream) Val() []NullableSeriesDataBulkRequest {
	if self == nil {
		return []NullableSeriesDataBulkRequest{}
	}
	return *self.Copy()
}
func (self *NullableSeriesDataBulkRequestStream) While(fn func(NullableSeriesDataBulkRequest, int) bool) *NullableSeriesDataBulkRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSeriesDataBulkRequestStream) Where(fn func(NullableSeriesDataBulkRequest) bool) *NullableSeriesDataBulkRequestStream {
	result := NullableSeriesDataBulkRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesDataBulkRequestStream) WhereSlim(fn func(NullableSeriesDataBulkRequest) bool) *NullableSeriesDataBulkRequestStream {
	result := NullableSeriesDataBulkRequestStreamOf()
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
