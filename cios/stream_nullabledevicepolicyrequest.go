/*
 * Collection utility of NullableDevicePolicyRequest Struct
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

type NullableDevicePolicyRequestStream []NullableDevicePolicyRequest

func NullableDevicePolicyRequestStreamOf(arg ...NullableDevicePolicyRequest) NullableDevicePolicyRequestStream {
	return arg
}
func NullableDevicePolicyRequestStreamFrom(arg []NullableDevicePolicyRequest) NullableDevicePolicyRequestStream {
	return arg
}
func CreateNullableDevicePolicyRequestStream(arg ...NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	tmp := NullableDevicePolicyRequestStreamOf(arg...)
	return &tmp
}
func GenerateNullableDevicePolicyRequestStream(arg []NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	tmp := NullableDevicePolicyRequestStreamFrom(arg)
	return &tmp
}

func (self *NullableDevicePolicyRequestStream) Add(arg NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	return self.AddAll(arg)
}
func (self *NullableDevicePolicyRequestStream) AddAll(arg ...NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableDevicePolicyRequestStream) AddSafe(arg *NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) Aggregate(fn func(NullableDevicePolicyRequest, NullableDevicePolicyRequest) NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	result := NullableDevicePolicyRequestStreamOf()
	self.ForEach(func(v NullableDevicePolicyRequest, i int) {
		if i == 0 {
			result.Add(fn(NullableDevicePolicyRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableDevicePolicyRequestStream) AllMatch(fn func(NullableDevicePolicyRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableDevicePolicyRequestStream) AnyMatch(fn func(NullableDevicePolicyRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableDevicePolicyRequestStream) Clone() *NullableDevicePolicyRequestStream {
	temp := make([]NullableDevicePolicyRequest, self.Len())
	copy(temp, *self)
	return (*NullableDevicePolicyRequestStream)(&temp)
}
func (self *NullableDevicePolicyRequestStream) Copy() *NullableDevicePolicyRequestStream {
	return self.Clone()
}
func (self *NullableDevicePolicyRequestStream) Concat(arg []NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	return self.AddAll(arg...)
}
func (self *NullableDevicePolicyRequestStream) Contains(arg NullableDevicePolicyRequest) bool {
	return self.FindIndex(func(_arg NullableDevicePolicyRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableDevicePolicyRequestStream) Clean() *NullableDevicePolicyRequestStream {
	*self = NullableDevicePolicyRequestStreamOf()
	return self
}
func (self *NullableDevicePolicyRequestStream) Delete(index int) *NullableDevicePolicyRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableDevicePolicyRequestStream) DeleteRange(startIndex, endIndex int) *NullableDevicePolicyRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableDevicePolicyRequestStream) Distinct() *NullableDevicePolicyRequestStream {
	caches := map[string]bool{}
	result := NullableDevicePolicyRequestStreamOf()
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
func (self *NullableDevicePolicyRequestStream) Each(fn func(NullableDevicePolicyRequest)) *NullableDevicePolicyRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) EachRight(fn func(NullableDevicePolicyRequest)) *NullableDevicePolicyRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) Equals(arr []NullableDevicePolicyRequest) bool {
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
func (self *NullableDevicePolicyRequestStream) Filter(fn func(NullableDevicePolicyRequest, int) bool) *NullableDevicePolicyRequestStream {
	result := NullableDevicePolicyRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDevicePolicyRequestStream) FilterSlim(fn func(NullableDevicePolicyRequest, int) bool) *NullableDevicePolicyRequestStream {
	result := NullableDevicePolicyRequestStreamOf()
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
func (self *NullableDevicePolicyRequestStream) Find(fn func(NullableDevicePolicyRequest, int) bool) *NullableDevicePolicyRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableDevicePolicyRequestStream) FindOr(fn func(NullableDevicePolicyRequest, int) bool, or NullableDevicePolicyRequest) NullableDevicePolicyRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableDevicePolicyRequestStream) FindIndex(fn func(NullableDevicePolicyRequest, int) bool) int {
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
func (self *NullableDevicePolicyRequestStream) First() *NullableDevicePolicyRequest {
	return self.Get(0)
}
func (self *NullableDevicePolicyRequestStream) FirstOr(arg NullableDevicePolicyRequest) NullableDevicePolicyRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDevicePolicyRequestStream) ForEach(fn func(NullableDevicePolicyRequest, int)) *NullableDevicePolicyRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) ForEachRight(fn func(NullableDevicePolicyRequest, int)) *NullableDevicePolicyRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) GroupBy(fn func(NullableDevicePolicyRequest, int) string) map[string][]NullableDevicePolicyRequest {
	m := map[string][]NullableDevicePolicyRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableDevicePolicyRequestStream) GroupByValues(fn func(NullableDevicePolicyRequest, int) string) [][]NullableDevicePolicyRequest {
	var tmp [][]NullableDevicePolicyRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableDevicePolicyRequestStream) IndexOf(arg NullableDevicePolicyRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableDevicePolicyRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableDevicePolicyRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableDevicePolicyRequestStream) Last() *NullableDevicePolicyRequest {
	return self.Get(self.Len() - 1)
}
func (self *NullableDevicePolicyRequestStream) LastOr(arg NullableDevicePolicyRequest) NullableDevicePolicyRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDevicePolicyRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableDevicePolicyRequestStream) Limit(limit int) *NullableDevicePolicyRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableDevicePolicyRequestStream) Map(fn func(NullableDevicePolicyRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Map2Int(fn func(NullableDevicePolicyRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Map2Int32(fn func(NullableDevicePolicyRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Map2Int64(fn func(NullableDevicePolicyRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Map2Float32(fn func(NullableDevicePolicyRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Map2Float64(fn func(NullableDevicePolicyRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Map2Bool(fn func(NullableDevicePolicyRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Map2Bytes(fn func(NullableDevicePolicyRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Map2String(fn func(NullableDevicePolicyRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Max(fn func(NullableDevicePolicyRequest, int) float64) *NullableDevicePolicyRequest {
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
func (self *NullableDevicePolicyRequestStream) Min(fn func(NullableDevicePolicyRequest, int) float64) *NullableDevicePolicyRequest {
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
func (self *NullableDevicePolicyRequestStream) NoneMatch(fn func(NullableDevicePolicyRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableDevicePolicyRequestStream) Get(index int) *NullableDevicePolicyRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableDevicePolicyRequestStream) GetOr(index int, arg NullableDevicePolicyRequest) NullableDevicePolicyRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDevicePolicyRequestStream) Peek(fn func(*NullableDevicePolicyRequest, int)) *NullableDevicePolicyRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableDevicePolicyRequestStream) Reduce(fn func(NullableDevicePolicyRequest, NullableDevicePolicyRequest, int) NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	return self.ReduceInit(fn, NullableDevicePolicyRequest{})
}
func (self *NullableDevicePolicyRequestStream) ReduceInit(fn func(NullableDevicePolicyRequest, NullableDevicePolicyRequest, int) NullableDevicePolicyRequest, initialValue NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	result := NullableDevicePolicyRequestStreamOf()
	self.ForEach(func(v NullableDevicePolicyRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableDevicePolicyRequestStream) ReduceInterface(fn func(interface{}, NullableDevicePolicyRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableDevicePolicyRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableDevicePolicyRequestStream) ReduceString(fn func(string, NullableDevicePolicyRequest, int) string) []string {
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
func (self *NullableDevicePolicyRequestStream) ReduceInt(fn func(int, NullableDevicePolicyRequest, int) int) []int {
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
func (self *NullableDevicePolicyRequestStream) ReduceInt32(fn func(int32, NullableDevicePolicyRequest, int) int32) []int32 {
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
func (self *NullableDevicePolicyRequestStream) ReduceInt64(fn func(int64, NullableDevicePolicyRequest, int) int64) []int64 {
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
func (self *NullableDevicePolicyRequestStream) ReduceFloat32(fn func(float32, NullableDevicePolicyRequest, int) float32) []float32 {
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
func (self *NullableDevicePolicyRequestStream) ReduceFloat64(fn func(float64, NullableDevicePolicyRequest, int) float64) []float64 {
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
func (self *NullableDevicePolicyRequestStream) ReduceBool(fn func(bool, NullableDevicePolicyRequest, int) bool) []bool {
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
func (self *NullableDevicePolicyRequestStream) Reverse() *NullableDevicePolicyRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) Replace(fn func(NullableDevicePolicyRequest, int) NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	return self.ForEach(func(v NullableDevicePolicyRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableDevicePolicyRequestStream) Select(fn func(NullableDevicePolicyRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableDevicePolicyRequestStream) Set(index int, val NullableDevicePolicyRequest) *NullableDevicePolicyRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) Skip(skip int) *NullableDevicePolicyRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableDevicePolicyRequestStream) SkippingEach(fn func(NullableDevicePolicyRequest, int) int) *NullableDevicePolicyRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) Slice(startIndex, n int) *NullableDevicePolicyRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableDevicePolicyRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) Sort(fn func(i, j int) bool) *NullableDevicePolicyRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableDevicePolicyRequestStream) Tail() *NullableDevicePolicyRequest {
	return self.Last()
}
func (self *NullableDevicePolicyRequestStream) TailOr(arg NullableDevicePolicyRequest) NullableDevicePolicyRequest {
	return self.LastOr(arg)
}
func (self *NullableDevicePolicyRequestStream) ToList() []NullableDevicePolicyRequest {
	return self.Val()
}
func (self *NullableDevicePolicyRequestStream) Unique() *NullableDevicePolicyRequestStream {
	return self.Distinct()
}
func (self *NullableDevicePolicyRequestStream) Val() []NullableDevicePolicyRequest {
	if self == nil {
		return []NullableDevicePolicyRequest{}
	}
	return *self.Copy()
}
func (self *NullableDevicePolicyRequestStream) While(fn func(NullableDevicePolicyRequest, int) bool) *NullableDevicePolicyRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableDevicePolicyRequestStream) Where(fn func(NullableDevicePolicyRequest) bool) *NullableDevicePolicyRequestStream {
	result := NullableDevicePolicyRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDevicePolicyRequestStream) WhereSlim(fn func(NullableDevicePolicyRequest) bool) *NullableDevicePolicyRequestStream {
	result := NullableDevicePolicyRequestStreamOf()
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
