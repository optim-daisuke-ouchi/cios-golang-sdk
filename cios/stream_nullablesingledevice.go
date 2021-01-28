/*
 * Collection utility of NullableSingleDevice Struct
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

type NullableSingleDeviceStream []NullableSingleDevice

func NullableSingleDeviceStreamOf(arg ...NullableSingleDevice) NullableSingleDeviceStream {
	return arg
}
func NullableSingleDeviceStreamFrom(arg []NullableSingleDevice) NullableSingleDeviceStream {
	return arg
}
func CreateNullableSingleDeviceStream(arg ...NullableSingleDevice) *NullableSingleDeviceStream {
	tmp := NullableSingleDeviceStreamOf(arg...)
	return &tmp
}
func GenerateNullableSingleDeviceStream(arg []NullableSingleDevice) *NullableSingleDeviceStream {
	tmp := NullableSingleDeviceStreamFrom(arg)
	return &tmp
}

func (self *NullableSingleDeviceStream) Add(arg NullableSingleDevice) *NullableSingleDeviceStream {
	return self.AddAll(arg)
}
func (self *NullableSingleDeviceStream) AddAll(arg ...NullableSingleDevice) *NullableSingleDeviceStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSingleDeviceStream) AddSafe(arg *NullableSingleDevice) *NullableSingleDeviceStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSingleDeviceStream) Aggregate(fn func(NullableSingleDevice, NullableSingleDevice) NullableSingleDevice) *NullableSingleDeviceStream {
	result := NullableSingleDeviceStreamOf()
	self.ForEach(func(v NullableSingleDevice, i int) {
		if i == 0 {
			result.Add(fn(NullableSingleDevice{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleDeviceStream) AllMatch(fn func(NullableSingleDevice, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSingleDeviceStream) AnyMatch(fn func(NullableSingleDevice, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSingleDeviceStream) Clone() *NullableSingleDeviceStream {
	temp := make([]NullableSingleDevice, self.Len())
	copy(temp, *self)
	return (*NullableSingleDeviceStream)(&temp)
}
func (self *NullableSingleDeviceStream) Copy() *NullableSingleDeviceStream {
	return self.Clone()
}
func (self *NullableSingleDeviceStream) Concat(arg []NullableSingleDevice) *NullableSingleDeviceStream {
	return self.AddAll(arg...)
}
func (self *NullableSingleDeviceStream) Contains(arg NullableSingleDevice) bool {
	return self.FindIndex(func(_arg NullableSingleDevice, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSingleDeviceStream) Clean() *NullableSingleDeviceStream {
	*self = NullableSingleDeviceStreamOf()
	return self
}
func (self *NullableSingleDeviceStream) Delete(index int) *NullableSingleDeviceStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSingleDeviceStream) DeleteRange(startIndex, endIndex int) *NullableSingleDeviceStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSingleDeviceStream) Distinct() *NullableSingleDeviceStream {
	caches := map[string]bool{}
	result := NullableSingleDeviceStreamOf()
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
func (self *NullableSingleDeviceStream) Each(fn func(NullableSingleDevice)) *NullableSingleDeviceStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSingleDeviceStream) EachRight(fn func(NullableSingleDevice)) *NullableSingleDeviceStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSingleDeviceStream) Equals(arr []NullableSingleDevice) bool {
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
func (self *NullableSingleDeviceStream) Filter(fn func(NullableSingleDevice, int) bool) *NullableSingleDeviceStream {
	result := NullableSingleDeviceStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleDeviceStream) FilterSlim(fn func(NullableSingleDevice, int) bool) *NullableSingleDeviceStream {
	result := NullableSingleDeviceStreamOf()
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
func (self *NullableSingleDeviceStream) Find(fn func(NullableSingleDevice, int) bool) *NullableSingleDevice {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSingleDeviceStream) FindOr(fn func(NullableSingleDevice, int) bool, or NullableSingleDevice) NullableSingleDevice {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSingleDeviceStream) FindIndex(fn func(NullableSingleDevice, int) bool) int {
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
func (self *NullableSingleDeviceStream) First() *NullableSingleDevice {
	return self.Get(0)
}
func (self *NullableSingleDeviceStream) FirstOr(arg NullableSingleDevice) NullableSingleDevice {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceStream) ForEach(fn func(NullableSingleDevice, int)) *NullableSingleDeviceStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSingleDeviceStream) ForEachRight(fn func(NullableSingleDevice, int)) *NullableSingleDeviceStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSingleDeviceStream) GroupBy(fn func(NullableSingleDevice, int) string) map[string][]NullableSingleDevice {
	m := map[string][]NullableSingleDevice{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSingleDeviceStream) GroupByValues(fn func(NullableSingleDevice, int) string) [][]NullableSingleDevice {
	var tmp [][]NullableSingleDevice
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSingleDeviceStream) IndexOf(arg NullableSingleDevice) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSingleDeviceStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSingleDeviceStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSingleDeviceStream) Last() *NullableSingleDevice {
	return self.Get(self.Len() - 1)
}
func (self *NullableSingleDeviceStream) LastOr(arg NullableSingleDevice) NullableSingleDevice {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSingleDeviceStream) Limit(limit int) *NullableSingleDeviceStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSingleDeviceStream) Map(fn func(NullableSingleDevice, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Map2Int(fn func(NullableSingleDevice, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Map2Int32(fn func(NullableSingleDevice, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Map2Int64(fn func(NullableSingleDevice, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Map2Float32(fn func(NullableSingleDevice, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Map2Float64(fn func(NullableSingleDevice, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Map2Bool(fn func(NullableSingleDevice, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Map2Bytes(fn func(NullableSingleDevice, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Map2String(fn func(NullableSingleDevice, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Max(fn func(NullableSingleDevice, int) float64) *NullableSingleDevice {
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
func (self *NullableSingleDeviceStream) Min(fn func(NullableSingleDevice, int) float64) *NullableSingleDevice {
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
func (self *NullableSingleDeviceStream) NoneMatch(fn func(NullableSingleDevice, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSingleDeviceStream) Get(index int) *NullableSingleDevice {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSingleDeviceStream) GetOr(index int, arg NullableSingleDevice) NullableSingleDevice {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceStream) Peek(fn func(*NullableSingleDevice, int)) *NullableSingleDeviceStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSingleDeviceStream) Reduce(fn func(NullableSingleDevice, NullableSingleDevice, int) NullableSingleDevice) *NullableSingleDeviceStream {
	return self.ReduceInit(fn, NullableSingleDevice{})
}
func (self *NullableSingleDeviceStream) ReduceInit(fn func(NullableSingleDevice, NullableSingleDevice, int) NullableSingleDevice, initialValue NullableSingleDevice) *NullableSingleDeviceStream {
	result := NullableSingleDeviceStreamOf()
	self.ForEach(func(v NullableSingleDevice, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleDeviceStream) ReduceInterface(fn func(interface{}, NullableSingleDevice, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSingleDevice{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSingleDeviceStream) ReduceString(fn func(string, NullableSingleDevice, int) string) []string {
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
func (self *NullableSingleDeviceStream) ReduceInt(fn func(int, NullableSingleDevice, int) int) []int {
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
func (self *NullableSingleDeviceStream) ReduceInt32(fn func(int32, NullableSingleDevice, int) int32) []int32 {
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
func (self *NullableSingleDeviceStream) ReduceInt64(fn func(int64, NullableSingleDevice, int) int64) []int64 {
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
func (self *NullableSingleDeviceStream) ReduceFloat32(fn func(float32, NullableSingleDevice, int) float32) []float32 {
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
func (self *NullableSingleDeviceStream) ReduceFloat64(fn func(float64, NullableSingleDevice, int) float64) []float64 {
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
func (self *NullableSingleDeviceStream) ReduceBool(fn func(bool, NullableSingleDevice, int) bool) []bool {
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
func (self *NullableSingleDeviceStream) Reverse() *NullableSingleDeviceStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSingleDeviceStream) Replace(fn func(NullableSingleDevice, int) NullableSingleDevice) *NullableSingleDeviceStream {
	return self.ForEach(func(v NullableSingleDevice, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSingleDeviceStream) Select(fn func(NullableSingleDevice) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSingleDeviceStream) Set(index int, val NullableSingleDevice) *NullableSingleDeviceStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSingleDeviceStream) Skip(skip int) *NullableSingleDeviceStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSingleDeviceStream) SkippingEach(fn func(NullableSingleDevice, int) int) *NullableSingleDeviceStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSingleDeviceStream) Slice(startIndex, n int) *NullableSingleDeviceStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSingleDevice{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSingleDeviceStream) Sort(fn func(i, j int) bool) *NullableSingleDeviceStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSingleDeviceStream) Tail() *NullableSingleDevice {
	return self.Last()
}
func (self *NullableSingleDeviceStream) TailOr(arg NullableSingleDevice) NullableSingleDevice {
	return self.LastOr(arg)
}
func (self *NullableSingleDeviceStream) ToList() []NullableSingleDevice {
	return self.Val()
}
func (self *NullableSingleDeviceStream) Unique() *NullableSingleDeviceStream {
	return self.Distinct()
}
func (self *NullableSingleDeviceStream) Val() []NullableSingleDevice {
	if self == nil {
		return []NullableSingleDevice{}
	}
	return *self.Copy()
}
func (self *NullableSingleDeviceStream) While(fn func(NullableSingleDevice, int) bool) *NullableSingleDeviceStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSingleDeviceStream) Where(fn func(NullableSingleDevice) bool) *NullableSingleDeviceStream {
	result := NullableSingleDeviceStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleDeviceStream) WhereSlim(fn func(NullableSingleDevice) bool) *NullableSingleDeviceStream {
	result := NullableSingleDeviceStreamOf()
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
