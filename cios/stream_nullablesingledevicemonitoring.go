/*
 * Collection utility of NullableSingleDeviceMonitoring Struct
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

type NullableSingleDeviceMonitoringStream []NullableSingleDeviceMonitoring

func NullableSingleDeviceMonitoringStreamOf(arg ...NullableSingleDeviceMonitoring) NullableSingleDeviceMonitoringStream {
	return arg
}
func NullableSingleDeviceMonitoringStreamFrom(arg []NullableSingleDeviceMonitoring) NullableSingleDeviceMonitoringStream {
	return arg
}
func CreateNullableSingleDeviceMonitoringStream(arg ...NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	tmp := NullableSingleDeviceMonitoringStreamOf(arg...)
	return &tmp
}
func GenerateNullableSingleDeviceMonitoringStream(arg []NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	tmp := NullableSingleDeviceMonitoringStreamFrom(arg)
	return &tmp
}

func (self *NullableSingleDeviceMonitoringStream) Add(arg NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	return self.AddAll(arg)
}
func (self *NullableSingleDeviceMonitoringStream) AddAll(arg ...NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSingleDeviceMonitoringStream) AddSafe(arg *NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) Aggregate(fn func(NullableSingleDeviceMonitoring, NullableSingleDeviceMonitoring) NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	result := NullableSingleDeviceMonitoringStreamOf()
	self.ForEach(func(v NullableSingleDeviceMonitoring, i int) {
		if i == 0 {
			result.Add(fn(NullableSingleDeviceMonitoring{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleDeviceMonitoringStream) AllMatch(fn func(NullableSingleDeviceMonitoring, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSingleDeviceMonitoringStream) AnyMatch(fn func(NullableSingleDeviceMonitoring, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSingleDeviceMonitoringStream) Clone() *NullableSingleDeviceMonitoringStream {
	temp := make([]NullableSingleDeviceMonitoring, self.Len())
	copy(temp, *self)
	return (*NullableSingleDeviceMonitoringStream)(&temp)
}
func (self *NullableSingleDeviceMonitoringStream) Copy() *NullableSingleDeviceMonitoringStream {
	return self.Clone()
}
func (self *NullableSingleDeviceMonitoringStream) Concat(arg []NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	return self.AddAll(arg...)
}
func (self *NullableSingleDeviceMonitoringStream) Contains(arg NullableSingleDeviceMonitoring) bool {
	return self.FindIndex(func(_arg NullableSingleDeviceMonitoring, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSingleDeviceMonitoringStream) Clean() *NullableSingleDeviceMonitoringStream {
	*self = NullableSingleDeviceMonitoringStreamOf()
	return self
}
func (self *NullableSingleDeviceMonitoringStream) Delete(index int) *NullableSingleDeviceMonitoringStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSingleDeviceMonitoringStream) DeleteRange(startIndex, endIndex int) *NullableSingleDeviceMonitoringStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSingleDeviceMonitoringStream) Distinct() *NullableSingleDeviceMonitoringStream {
	caches := map[string]bool{}
	result := NullableSingleDeviceMonitoringStreamOf()
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
func (self *NullableSingleDeviceMonitoringStream) Each(fn func(NullableSingleDeviceMonitoring)) *NullableSingleDeviceMonitoringStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) EachRight(fn func(NullableSingleDeviceMonitoring)) *NullableSingleDeviceMonitoringStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) Equals(arr []NullableSingleDeviceMonitoring) bool {
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
func (self *NullableSingleDeviceMonitoringStream) Filter(fn func(NullableSingleDeviceMonitoring, int) bool) *NullableSingleDeviceMonitoringStream {
	result := NullableSingleDeviceMonitoringStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleDeviceMonitoringStream) FilterSlim(fn func(NullableSingleDeviceMonitoring, int) bool) *NullableSingleDeviceMonitoringStream {
	result := NullableSingleDeviceMonitoringStreamOf()
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
func (self *NullableSingleDeviceMonitoringStream) Find(fn func(NullableSingleDeviceMonitoring, int) bool) *NullableSingleDeviceMonitoring {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSingleDeviceMonitoringStream) FindOr(fn func(NullableSingleDeviceMonitoring, int) bool, or NullableSingleDeviceMonitoring) NullableSingleDeviceMonitoring {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSingleDeviceMonitoringStream) FindIndex(fn func(NullableSingleDeviceMonitoring, int) bool) int {
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
func (self *NullableSingleDeviceMonitoringStream) First() *NullableSingleDeviceMonitoring {
	return self.Get(0)
}
func (self *NullableSingleDeviceMonitoringStream) FirstOr(arg NullableSingleDeviceMonitoring) NullableSingleDeviceMonitoring {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceMonitoringStream) ForEach(fn func(NullableSingleDeviceMonitoring, int)) *NullableSingleDeviceMonitoringStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) ForEachRight(fn func(NullableSingleDeviceMonitoring, int)) *NullableSingleDeviceMonitoringStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) GroupBy(fn func(NullableSingleDeviceMonitoring, int) string) map[string][]NullableSingleDeviceMonitoring {
	m := map[string][]NullableSingleDeviceMonitoring{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSingleDeviceMonitoringStream) GroupByValues(fn func(NullableSingleDeviceMonitoring, int) string) [][]NullableSingleDeviceMonitoring {
	var tmp [][]NullableSingleDeviceMonitoring
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSingleDeviceMonitoringStream) IndexOf(arg NullableSingleDeviceMonitoring) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSingleDeviceMonitoringStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSingleDeviceMonitoringStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSingleDeviceMonitoringStream) Last() *NullableSingleDeviceMonitoring {
	return self.Get(self.Len() - 1)
}
func (self *NullableSingleDeviceMonitoringStream) LastOr(arg NullableSingleDeviceMonitoring) NullableSingleDeviceMonitoring {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceMonitoringStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSingleDeviceMonitoringStream) Limit(limit int) *NullableSingleDeviceMonitoringStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSingleDeviceMonitoringStream) Map(fn func(NullableSingleDeviceMonitoring, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Map2Int(fn func(NullableSingleDeviceMonitoring, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Map2Int32(fn func(NullableSingleDeviceMonitoring, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Map2Int64(fn func(NullableSingleDeviceMonitoring, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Map2Float32(fn func(NullableSingleDeviceMonitoring, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Map2Float64(fn func(NullableSingleDeviceMonitoring, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Map2Bool(fn func(NullableSingleDeviceMonitoring, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Map2Bytes(fn func(NullableSingleDeviceMonitoring, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Map2String(fn func(NullableSingleDeviceMonitoring, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Max(fn func(NullableSingleDeviceMonitoring, int) float64) *NullableSingleDeviceMonitoring {
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
func (self *NullableSingleDeviceMonitoringStream) Min(fn func(NullableSingleDeviceMonitoring, int) float64) *NullableSingleDeviceMonitoring {
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
func (self *NullableSingleDeviceMonitoringStream) NoneMatch(fn func(NullableSingleDeviceMonitoring, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSingleDeviceMonitoringStream) Get(index int) *NullableSingleDeviceMonitoring {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSingleDeviceMonitoringStream) GetOr(index int, arg NullableSingleDeviceMonitoring) NullableSingleDeviceMonitoring {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceMonitoringStream) Peek(fn func(*NullableSingleDeviceMonitoring, int)) *NullableSingleDeviceMonitoringStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSingleDeviceMonitoringStream) Reduce(fn func(NullableSingleDeviceMonitoring, NullableSingleDeviceMonitoring, int) NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	return self.ReduceInit(fn, NullableSingleDeviceMonitoring{})
}
func (self *NullableSingleDeviceMonitoringStream) ReduceInit(fn func(NullableSingleDeviceMonitoring, NullableSingleDeviceMonitoring, int) NullableSingleDeviceMonitoring, initialValue NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	result := NullableSingleDeviceMonitoringStreamOf()
	self.ForEach(func(v NullableSingleDeviceMonitoring, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleDeviceMonitoringStream) ReduceInterface(fn func(interface{}, NullableSingleDeviceMonitoring, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSingleDeviceMonitoring{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSingleDeviceMonitoringStream) ReduceString(fn func(string, NullableSingleDeviceMonitoring, int) string) []string {
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
func (self *NullableSingleDeviceMonitoringStream) ReduceInt(fn func(int, NullableSingleDeviceMonitoring, int) int) []int {
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
func (self *NullableSingleDeviceMonitoringStream) ReduceInt32(fn func(int32, NullableSingleDeviceMonitoring, int) int32) []int32 {
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
func (self *NullableSingleDeviceMonitoringStream) ReduceInt64(fn func(int64, NullableSingleDeviceMonitoring, int) int64) []int64 {
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
func (self *NullableSingleDeviceMonitoringStream) ReduceFloat32(fn func(float32, NullableSingleDeviceMonitoring, int) float32) []float32 {
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
func (self *NullableSingleDeviceMonitoringStream) ReduceFloat64(fn func(float64, NullableSingleDeviceMonitoring, int) float64) []float64 {
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
func (self *NullableSingleDeviceMonitoringStream) ReduceBool(fn func(bool, NullableSingleDeviceMonitoring, int) bool) []bool {
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
func (self *NullableSingleDeviceMonitoringStream) Reverse() *NullableSingleDeviceMonitoringStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) Replace(fn func(NullableSingleDeviceMonitoring, int) NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	return self.ForEach(func(v NullableSingleDeviceMonitoring, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSingleDeviceMonitoringStream) Select(fn func(NullableSingleDeviceMonitoring) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSingleDeviceMonitoringStream) Set(index int, val NullableSingleDeviceMonitoring) *NullableSingleDeviceMonitoringStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) Skip(skip int) *NullableSingleDeviceMonitoringStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSingleDeviceMonitoringStream) SkippingEach(fn func(NullableSingleDeviceMonitoring, int) int) *NullableSingleDeviceMonitoringStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) Slice(startIndex, n int) *NullableSingleDeviceMonitoringStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSingleDeviceMonitoring{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) Sort(fn func(i, j int) bool) *NullableSingleDeviceMonitoringStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSingleDeviceMonitoringStream) Tail() *NullableSingleDeviceMonitoring {
	return self.Last()
}
func (self *NullableSingleDeviceMonitoringStream) TailOr(arg NullableSingleDeviceMonitoring) NullableSingleDeviceMonitoring {
	return self.LastOr(arg)
}
func (self *NullableSingleDeviceMonitoringStream) ToList() []NullableSingleDeviceMonitoring {
	return self.Val()
}
func (self *NullableSingleDeviceMonitoringStream) Unique() *NullableSingleDeviceMonitoringStream {
	return self.Distinct()
}
func (self *NullableSingleDeviceMonitoringStream) Val() []NullableSingleDeviceMonitoring {
	if self == nil {
		return []NullableSingleDeviceMonitoring{}
	}
	return *self.Copy()
}
func (self *NullableSingleDeviceMonitoringStream) While(fn func(NullableSingleDeviceMonitoring, int) bool) *NullableSingleDeviceMonitoringStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSingleDeviceMonitoringStream) Where(fn func(NullableSingleDeviceMonitoring) bool) *NullableSingleDeviceMonitoringStream {
	result := NullableSingleDeviceMonitoringStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleDeviceMonitoringStream) WhereSlim(fn func(NullableSingleDeviceMonitoring) bool) *NullableSingleDeviceMonitoringStream {
	result := NullableSingleDeviceMonitoringStreamOf()
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
