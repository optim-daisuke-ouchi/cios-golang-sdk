/*
 * Collection utility of NullableDeviceMonitoringIDsRequest Struct
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

type NullableDeviceMonitoringIDsRequestStream []NullableDeviceMonitoringIDsRequest

func NullableDeviceMonitoringIDsRequestStreamOf(arg ...NullableDeviceMonitoringIDsRequest) NullableDeviceMonitoringIDsRequestStream {
	return arg
}
func NullableDeviceMonitoringIDsRequestStreamFrom(arg []NullableDeviceMonitoringIDsRequest) NullableDeviceMonitoringIDsRequestStream {
	return arg
}
func CreateNullableDeviceMonitoringIDsRequestStream(arg ...NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	tmp := NullableDeviceMonitoringIDsRequestStreamOf(arg...)
	return &tmp
}
func GenerateNullableDeviceMonitoringIDsRequestStream(arg []NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	tmp := NullableDeviceMonitoringIDsRequestStreamFrom(arg)
	return &tmp
}

func (self *NullableDeviceMonitoringIDsRequestStream) Add(arg NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	return self.AddAll(arg)
}
func (self *NullableDeviceMonitoringIDsRequestStream) AddAll(arg ...NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) AddSafe(arg *NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) Aggregate(fn func(NullableDeviceMonitoringIDsRequest, NullableDeviceMonitoringIDsRequest) NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	result := NullableDeviceMonitoringIDsRequestStreamOf()
	self.ForEach(func(v NullableDeviceMonitoringIDsRequest, i int) {
		if i == 0 {
			result.Add(fn(NullableDeviceMonitoringIDsRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) AllMatch(fn func(NullableDeviceMonitoringIDsRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableDeviceMonitoringIDsRequestStream) AnyMatch(fn func(NullableDeviceMonitoringIDsRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableDeviceMonitoringIDsRequestStream) Clone() *NullableDeviceMonitoringIDsRequestStream {
	temp := make([]NullableDeviceMonitoringIDsRequest, self.Len())
	copy(temp, *self)
	return (*NullableDeviceMonitoringIDsRequestStream)(&temp)
}
func (self *NullableDeviceMonitoringIDsRequestStream) Copy() *NullableDeviceMonitoringIDsRequestStream {
	return self.Clone()
}
func (self *NullableDeviceMonitoringIDsRequestStream) Concat(arg []NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	return self.AddAll(arg...)
}
func (self *NullableDeviceMonitoringIDsRequestStream) Contains(arg NullableDeviceMonitoringIDsRequest) bool {
	return self.FindIndex(func(_arg NullableDeviceMonitoringIDsRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableDeviceMonitoringIDsRequestStream) Clean() *NullableDeviceMonitoringIDsRequestStream {
	*self = NullableDeviceMonitoringIDsRequestStreamOf()
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) Delete(index int) *NullableDeviceMonitoringIDsRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableDeviceMonitoringIDsRequestStream) DeleteRange(startIndex, endIndex int) *NullableDeviceMonitoringIDsRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) Distinct() *NullableDeviceMonitoringIDsRequestStream {
	caches := map[string]bool{}
	result := NullableDeviceMonitoringIDsRequestStreamOf()
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
func (self *NullableDeviceMonitoringIDsRequestStream) Each(fn func(NullableDeviceMonitoringIDsRequest)) *NullableDeviceMonitoringIDsRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) EachRight(fn func(NullableDeviceMonitoringIDsRequest)) *NullableDeviceMonitoringIDsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) Equals(arr []NullableDeviceMonitoringIDsRequest) bool {
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
func (self *NullableDeviceMonitoringIDsRequestStream) Filter(fn func(NullableDeviceMonitoringIDsRequest, int) bool) *NullableDeviceMonitoringIDsRequestStream {
	result := NullableDeviceMonitoringIDsRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) FilterSlim(fn func(NullableDeviceMonitoringIDsRequest, int) bool) *NullableDeviceMonitoringIDsRequestStream {
	result := NullableDeviceMonitoringIDsRequestStreamOf()
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
func (self *NullableDeviceMonitoringIDsRequestStream) Find(fn func(NullableDeviceMonitoringIDsRequest, int) bool) *NullableDeviceMonitoringIDsRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableDeviceMonitoringIDsRequestStream) FindOr(fn func(NullableDeviceMonitoringIDsRequest, int) bool, or NullableDeviceMonitoringIDsRequest) NullableDeviceMonitoringIDsRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableDeviceMonitoringIDsRequestStream) FindIndex(fn func(NullableDeviceMonitoringIDsRequest, int) bool) int {
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
func (self *NullableDeviceMonitoringIDsRequestStream) First() *NullableDeviceMonitoringIDsRequest {
	return self.Get(0)
}
func (self *NullableDeviceMonitoringIDsRequestStream) FirstOr(arg NullableDeviceMonitoringIDsRequest) NullableDeviceMonitoringIDsRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceMonitoringIDsRequestStream) ForEach(fn func(NullableDeviceMonitoringIDsRequest, int)) *NullableDeviceMonitoringIDsRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) ForEachRight(fn func(NullableDeviceMonitoringIDsRequest, int)) *NullableDeviceMonitoringIDsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) GroupBy(fn func(NullableDeviceMonitoringIDsRequest, int) string) map[string][]NullableDeviceMonitoringIDsRequest {
	m := map[string][]NullableDeviceMonitoringIDsRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableDeviceMonitoringIDsRequestStream) GroupByValues(fn func(NullableDeviceMonitoringIDsRequest, int) string) [][]NullableDeviceMonitoringIDsRequest {
	var tmp [][]NullableDeviceMonitoringIDsRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableDeviceMonitoringIDsRequestStream) IndexOf(arg NullableDeviceMonitoringIDsRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableDeviceMonitoringIDsRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableDeviceMonitoringIDsRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableDeviceMonitoringIDsRequestStream) Last() *NullableDeviceMonitoringIDsRequest {
	return self.Get(self.Len() - 1)
}
func (self *NullableDeviceMonitoringIDsRequestStream) LastOr(arg NullableDeviceMonitoringIDsRequest) NullableDeviceMonitoringIDsRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceMonitoringIDsRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableDeviceMonitoringIDsRequestStream) Limit(limit int) *NullableDeviceMonitoringIDsRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableDeviceMonitoringIDsRequestStream) Map(fn func(NullableDeviceMonitoringIDsRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Map2Int(fn func(NullableDeviceMonitoringIDsRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Map2Int32(fn func(NullableDeviceMonitoringIDsRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Map2Int64(fn func(NullableDeviceMonitoringIDsRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Map2Float32(fn func(NullableDeviceMonitoringIDsRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Map2Float64(fn func(NullableDeviceMonitoringIDsRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Map2Bool(fn func(NullableDeviceMonitoringIDsRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Map2Bytes(fn func(NullableDeviceMonitoringIDsRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Map2String(fn func(NullableDeviceMonitoringIDsRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Max(fn func(NullableDeviceMonitoringIDsRequest, int) float64) *NullableDeviceMonitoringIDsRequest {
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
func (self *NullableDeviceMonitoringIDsRequestStream) Min(fn func(NullableDeviceMonitoringIDsRequest, int) float64) *NullableDeviceMonitoringIDsRequest {
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
func (self *NullableDeviceMonitoringIDsRequestStream) NoneMatch(fn func(NullableDeviceMonitoringIDsRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableDeviceMonitoringIDsRequestStream) Get(index int) *NullableDeviceMonitoringIDsRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableDeviceMonitoringIDsRequestStream) GetOr(index int, arg NullableDeviceMonitoringIDsRequest) NullableDeviceMonitoringIDsRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceMonitoringIDsRequestStream) Peek(fn func(*NullableDeviceMonitoringIDsRequest, int)) *NullableDeviceMonitoringIDsRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableDeviceMonitoringIDsRequestStream) Reduce(fn func(NullableDeviceMonitoringIDsRequest, NullableDeviceMonitoringIDsRequest, int) NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	return self.ReduceInit(fn, NullableDeviceMonitoringIDsRequest{})
}
func (self *NullableDeviceMonitoringIDsRequestStream) ReduceInit(fn func(NullableDeviceMonitoringIDsRequest, NullableDeviceMonitoringIDsRequest, int) NullableDeviceMonitoringIDsRequest, initialValue NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	result := NullableDeviceMonitoringIDsRequestStreamOf()
	self.ForEach(func(v NullableDeviceMonitoringIDsRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) ReduceInterface(fn func(interface{}, NullableDeviceMonitoringIDsRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableDeviceMonitoringIDsRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableDeviceMonitoringIDsRequestStream) ReduceString(fn func(string, NullableDeviceMonitoringIDsRequest, int) string) []string {
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
func (self *NullableDeviceMonitoringIDsRequestStream) ReduceInt(fn func(int, NullableDeviceMonitoringIDsRequest, int) int) []int {
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
func (self *NullableDeviceMonitoringIDsRequestStream) ReduceInt32(fn func(int32, NullableDeviceMonitoringIDsRequest, int) int32) []int32 {
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
func (self *NullableDeviceMonitoringIDsRequestStream) ReduceInt64(fn func(int64, NullableDeviceMonitoringIDsRequest, int) int64) []int64 {
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
func (self *NullableDeviceMonitoringIDsRequestStream) ReduceFloat32(fn func(float32, NullableDeviceMonitoringIDsRequest, int) float32) []float32 {
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
func (self *NullableDeviceMonitoringIDsRequestStream) ReduceFloat64(fn func(float64, NullableDeviceMonitoringIDsRequest, int) float64) []float64 {
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
func (self *NullableDeviceMonitoringIDsRequestStream) ReduceBool(fn func(bool, NullableDeviceMonitoringIDsRequest, int) bool) []bool {
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
func (self *NullableDeviceMonitoringIDsRequestStream) Reverse() *NullableDeviceMonitoringIDsRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) Replace(fn func(NullableDeviceMonitoringIDsRequest, int) NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	return self.ForEach(func(v NullableDeviceMonitoringIDsRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableDeviceMonitoringIDsRequestStream) Select(fn func(NullableDeviceMonitoringIDsRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableDeviceMonitoringIDsRequestStream) Set(index int, val NullableDeviceMonitoringIDsRequest) *NullableDeviceMonitoringIDsRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) Skip(skip int) *NullableDeviceMonitoringIDsRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableDeviceMonitoringIDsRequestStream) SkippingEach(fn func(NullableDeviceMonitoringIDsRequest, int) int) *NullableDeviceMonitoringIDsRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) Slice(startIndex, n int) *NullableDeviceMonitoringIDsRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableDeviceMonitoringIDsRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) Sort(fn func(i, j int) bool) *NullableDeviceMonitoringIDsRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableDeviceMonitoringIDsRequestStream) Tail() *NullableDeviceMonitoringIDsRequest {
	return self.Last()
}
func (self *NullableDeviceMonitoringIDsRequestStream) TailOr(arg NullableDeviceMonitoringIDsRequest) NullableDeviceMonitoringIDsRequest {
	return self.LastOr(arg)
}
func (self *NullableDeviceMonitoringIDsRequestStream) ToList() []NullableDeviceMonitoringIDsRequest {
	return self.Val()
}
func (self *NullableDeviceMonitoringIDsRequestStream) Unique() *NullableDeviceMonitoringIDsRequestStream {
	return self.Distinct()
}
func (self *NullableDeviceMonitoringIDsRequestStream) Val() []NullableDeviceMonitoringIDsRequest {
	if self == nil {
		return []NullableDeviceMonitoringIDsRequest{}
	}
	return *self.Copy()
}
func (self *NullableDeviceMonitoringIDsRequestStream) While(fn func(NullableDeviceMonitoringIDsRequest, int) bool) *NullableDeviceMonitoringIDsRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) Where(fn func(NullableDeviceMonitoringIDsRequest) bool) *NullableDeviceMonitoringIDsRequestStream {
	result := NullableDeviceMonitoringIDsRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDeviceMonitoringIDsRequestStream) WhereSlim(fn func(NullableDeviceMonitoringIDsRequest) bool) *NullableDeviceMonitoringIDsRequestStream {
	result := NullableDeviceMonitoringIDsRequestStreamOf()
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
