/*
 * Collection utility of NullableDevicePolicy Struct
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

type NullableDevicePolicyStream []NullableDevicePolicy

func NullableDevicePolicyStreamOf(arg ...NullableDevicePolicy) NullableDevicePolicyStream {
	return arg
}
func NullableDevicePolicyStreamFrom(arg []NullableDevicePolicy) NullableDevicePolicyStream {
	return arg
}
func CreateNullableDevicePolicyStream(arg ...NullableDevicePolicy) *NullableDevicePolicyStream {
	tmp := NullableDevicePolicyStreamOf(arg...)
	return &tmp
}
func GenerateNullableDevicePolicyStream(arg []NullableDevicePolicy) *NullableDevicePolicyStream {
	tmp := NullableDevicePolicyStreamFrom(arg)
	return &tmp
}

func (self *NullableDevicePolicyStream) Add(arg NullableDevicePolicy) *NullableDevicePolicyStream {
	return self.AddAll(arg)
}
func (self *NullableDevicePolicyStream) AddAll(arg ...NullableDevicePolicy) *NullableDevicePolicyStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableDevicePolicyStream) AddSafe(arg *NullableDevicePolicy) *NullableDevicePolicyStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableDevicePolicyStream) Aggregate(fn func(NullableDevicePolicy, NullableDevicePolicy) NullableDevicePolicy) *NullableDevicePolicyStream {
	result := NullableDevicePolicyStreamOf()
	self.ForEach(func(v NullableDevicePolicy, i int) {
		if i == 0 {
			result.Add(fn(NullableDevicePolicy{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableDevicePolicyStream) AllMatch(fn func(NullableDevicePolicy, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableDevicePolicyStream) AnyMatch(fn func(NullableDevicePolicy, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableDevicePolicyStream) Clone() *NullableDevicePolicyStream {
	temp := make([]NullableDevicePolicy, self.Len())
	copy(temp, *self)
	return (*NullableDevicePolicyStream)(&temp)
}
func (self *NullableDevicePolicyStream) Copy() *NullableDevicePolicyStream {
	return self.Clone()
}
func (self *NullableDevicePolicyStream) Concat(arg []NullableDevicePolicy) *NullableDevicePolicyStream {
	return self.AddAll(arg...)
}
func (self *NullableDevicePolicyStream) Contains(arg NullableDevicePolicy) bool {
	return self.FindIndex(func(_arg NullableDevicePolicy, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableDevicePolicyStream) Clean() *NullableDevicePolicyStream {
	*self = NullableDevicePolicyStreamOf()
	return self
}
func (self *NullableDevicePolicyStream) Delete(index int) *NullableDevicePolicyStream {
	return self.DeleteRange(index, index)
}
func (self *NullableDevicePolicyStream) DeleteRange(startIndex, endIndex int) *NullableDevicePolicyStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableDevicePolicyStream) Distinct() *NullableDevicePolicyStream {
	caches := map[string]bool{}
	result := NullableDevicePolicyStreamOf()
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
func (self *NullableDevicePolicyStream) Each(fn func(NullableDevicePolicy)) *NullableDevicePolicyStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableDevicePolicyStream) EachRight(fn func(NullableDevicePolicy)) *NullableDevicePolicyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableDevicePolicyStream) Equals(arr []NullableDevicePolicy) bool {
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
func (self *NullableDevicePolicyStream) Filter(fn func(NullableDevicePolicy, int) bool) *NullableDevicePolicyStream {
	result := NullableDevicePolicyStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDevicePolicyStream) FilterSlim(fn func(NullableDevicePolicy, int) bool) *NullableDevicePolicyStream {
	result := NullableDevicePolicyStreamOf()
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
func (self *NullableDevicePolicyStream) Find(fn func(NullableDevicePolicy, int) bool) *NullableDevicePolicy {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableDevicePolicyStream) FindOr(fn func(NullableDevicePolicy, int) bool, or NullableDevicePolicy) NullableDevicePolicy {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableDevicePolicyStream) FindIndex(fn func(NullableDevicePolicy, int) bool) int {
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
func (self *NullableDevicePolicyStream) First() *NullableDevicePolicy {
	return self.Get(0)
}
func (self *NullableDevicePolicyStream) FirstOr(arg NullableDevicePolicy) NullableDevicePolicy {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDevicePolicyStream) ForEach(fn func(NullableDevicePolicy, int)) *NullableDevicePolicyStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableDevicePolicyStream) ForEachRight(fn func(NullableDevicePolicy, int)) *NullableDevicePolicyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableDevicePolicyStream) GroupBy(fn func(NullableDevicePolicy, int) string) map[string][]NullableDevicePolicy {
	m := map[string][]NullableDevicePolicy{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableDevicePolicyStream) GroupByValues(fn func(NullableDevicePolicy, int) string) [][]NullableDevicePolicy {
	var tmp [][]NullableDevicePolicy
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableDevicePolicyStream) IndexOf(arg NullableDevicePolicy) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableDevicePolicyStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableDevicePolicyStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableDevicePolicyStream) Last() *NullableDevicePolicy {
	return self.Get(self.Len() - 1)
}
func (self *NullableDevicePolicyStream) LastOr(arg NullableDevicePolicy) NullableDevicePolicy {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDevicePolicyStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableDevicePolicyStream) Limit(limit int) *NullableDevicePolicyStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableDevicePolicyStream) Map(fn func(NullableDevicePolicy, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Map2Int(fn func(NullableDevicePolicy, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Map2Int32(fn func(NullableDevicePolicy, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Map2Int64(fn func(NullableDevicePolicy, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Map2Float32(fn func(NullableDevicePolicy, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Map2Float64(fn func(NullableDevicePolicy, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Map2Bool(fn func(NullableDevicePolicy, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Map2Bytes(fn func(NullableDevicePolicy, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Map2String(fn func(NullableDevicePolicy, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Max(fn func(NullableDevicePolicy, int) float64) *NullableDevicePolicy {
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
func (self *NullableDevicePolicyStream) Min(fn func(NullableDevicePolicy, int) float64) *NullableDevicePolicy {
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
func (self *NullableDevicePolicyStream) NoneMatch(fn func(NullableDevicePolicy, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableDevicePolicyStream) Get(index int) *NullableDevicePolicy {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableDevicePolicyStream) GetOr(index int, arg NullableDevicePolicy) NullableDevicePolicy {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDevicePolicyStream) Peek(fn func(*NullableDevicePolicy, int)) *NullableDevicePolicyStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableDevicePolicyStream) Reduce(fn func(NullableDevicePolicy, NullableDevicePolicy, int) NullableDevicePolicy) *NullableDevicePolicyStream {
	return self.ReduceInit(fn, NullableDevicePolicy{})
}
func (self *NullableDevicePolicyStream) ReduceInit(fn func(NullableDevicePolicy, NullableDevicePolicy, int) NullableDevicePolicy, initialValue NullableDevicePolicy) *NullableDevicePolicyStream {
	result := NullableDevicePolicyStreamOf()
	self.ForEach(func(v NullableDevicePolicy, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableDevicePolicyStream) ReduceInterface(fn func(interface{}, NullableDevicePolicy, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableDevicePolicy{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableDevicePolicyStream) ReduceString(fn func(string, NullableDevicePolicy, int) string) []string {
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
func (self *NullableDevicePolicyStream) ReduceInt(fn func(int, NullableDevicePolicy, int) int) []int {
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
func (self *NullableDevicePolicyStream) ReduceInt32(fn func(int32, NullableDevicePolicy, int) int32) []int32 {
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
func (self *NullableDevicePolicyStream) ReduceInt64(fn func(int64, NullableDevicePolicy, int) int64) []int64 {
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
func (self *NullableDevicePolicyStream) ReduceFloat32(fn func(float32, NullableDevicePolicy, int) float32) []float32 {
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
func (self *NullableDevicePolicyStream) ReduceFloat64(fn func(float64, NullableDevicePolicy, int) float64) []float64 {
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
func (self *NullableDevicePolicyStream) ReduceBool(fn func(bool, NullableDevicePolicy, int) bool) []bool {
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
func (self *NullableDevicePolicyStream) Reverse() *NullableDevicePolicyStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableDevicePolicyStream) Replace(fn func(NullableDevicePolicy, int) NullableDevicePolicy) *NullableDevicePolicyStream {
	return self.ForEach(func(v NullableDevicePolicy, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableDevicePolicyStream) Select(fn func(NullableDevicePolicy) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableDevicePolicyStream) Set(index int, val NullableDevicePolicy) *NullableDevicePolicyStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableDevicePolicyStream) Skip(skip int) *NullableDevicePolicyStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableDevicePolicyStream) SkippingEach(fn func(NullableDevicePolicy, int) int) *NullableDevicePolicyStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableDevicePolicyStream) Slice(startIndex, n int) *NullableDevicePolicyStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableDevicePolicy{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableDevicePolicyStream) Sort(fn func(i, j int) bool) *NullableDevicePolicyStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableDevicePolicyStream) Tail() *NullableDevicePolicy {
	return self.Last()
}
func (self *NullableDevicePolicyStream) TailOr(arg NullableDevicePolicy) NullableDevicePolicy {
	return self.LastOr(arg)
}
func (self *NullableDevicePolicyStream) ToList() []NullableDevicePolicy {
	return self.Val()
}
func (self *NullableDevicePolicyStream) Unique() *NullableDevicePolicyStream {
	return self.Distinct()
}
func (self *NullableDevicePolicyStream) Val() []NullableDevicePolicy {
	if self == nil {
		return []NullableDevicePolicy{}
	}
	return *self.Copy()
}
func (self *NullableDevicePolicyStream) While(fn func(NullableDevicePolicy, int) bool) *NullableDevicePolicyStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableDevicePolicyStream) Where(fn func(NullableDevicePolicy) bool) *NullableDevicePolicyStream {
	result := NullableDevicePolicyStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDevicePolicyStream) WhereSlim(fn func(NullableDevicePolicy) bool) *NullableDevicePolicyStream {
	result := NullableDevicePolicyStreamOf()
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
