/*
 * Collection utility of DeviceClientRsaKey Struct
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

type DeviceClientRsaKeyStream []DeviceClientRsaKey

func DeviceClientRsaKeyStreamOf(arg ...DeviceClientRsaKey) DeviceClientRsaKeyStream {
	return arg
}
func DeviceClientRsaKeyStreamFrom(arg []DeviceClientRsaKey) DeviceClientRsaKeyStream {
	return arg
}
func CreateDeviceClientRsaKeyStream(arg ...DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	tmp := DeviceClientRsaKeyStreamOf(arg...)
	return &tmp
}
func GenerateDeviceClientRsaKeyStream(arg []DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	tmp := DeviceClientRsaKeyStreamFrom(arg)
	return &tmp
}

func (self *DeviceClientRsaKeyStream) Add(arg DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	return self.AddAll(arg)
}
func (self *DeviceClientRsaKeyStream) AddAll(arg ...DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	*self = append(*self, arg...)
	return self
}
func (self *DeviceClientRsaKeyStream) AddSafe(arg *DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *DeviceClientRsaKeyStream) Aggregate(fn func(DeviceClientRsaKey, DeviceClientRsaKey) DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	result := DeviceClientRsaKeyStreamOf()
	self.ForEach(func(v DeviceClientRsaKey, i int) {
		if i == 0 {
			result.Add(fn(DeviceClientRsaKey{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *DeviceClientRsaKeyStream) AllMatch(fn func(DeviceClientRsaKey, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *DeviceClientRsaKeyStream) AnyMatch(fn func(DeviceClientRsaKey, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *DeviceClientRsaKeyStream) Clone() *DeviceClientRsaKeyStream {
	temp := make([]DeviceClientRsaKey, self.Len())
	copy(temp, *self)
	return (*DeviceClientRsaKeyStream)(&temp)
}
func (self *DeviceClientRsaKeyStream) Copy() *DeviceClientRsaKeyStream {
	return self.Clone()
}
func (self *DeviceClientRsaKeyStream) Concat(arg []DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	return self.AddAll(arg...)
}
func (self *DeviceClientRsaKeyStream) Contains(arg DeviceClientRsaKey) bool {
	return self.FindIndex(func(_arg DeviceClientRsaKey, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *DeviceClientRsaKeyStream) Clean() *DeviceClientRsaKeyStream {
	*self = DeviceClientRsaKeyStreamOf()
	return self
}
func (self *DeviceClientRsaKeyStream) Delete(index int) *DeviceClientRsaKeyStream {
	return self.DeleteRange(index, index)
}
func (self *DeviceClientRsaKeyStream) DeleteRange(startIndex, endIndex int) *DeviceClientRsaKeyStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *DeviceClientRsaKeyStream) Distinct() *DeviceClientRsaKeyStream {
	caches := map[string]bool{}
	result := DeviceClientRsaKeyStreamOf()
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
func (self *DeviceClientRsaKeyStream) Each(fn func(DeviceClientRsaKey)) *DeviceClientRsaKeyStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *DeviceClientRsaKeyStream) EachRight(fn func(DeviceClientRsaKey)) *DeviceClientRsaKeyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *DeviceClientRsaKeyStream) Equals(arr []DeviceClientRsaKey) bool {
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
func (self *DeviceClientRsaKeyStream) Filter(fn func(DeviceClientRsaKey, int) bool) *DeviceClientRsaKeyStream {
	result := DeviceClientRsaKeyStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DeviceClientRsaKeyStream) FilterSlim(fn func(DeviceClientRsaKey, int) bool) *DeviceClientRsaKeyStream {
	result := DeviceClientRsaKeyStreamOf()
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
func (self *DeviceClientRsaKeyStream) Find(fn func(DeviceClientRsaKey, int) bool) *DeviceClientRsaKey {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *DeviceClientRsaKeyStream) FindOr(fn func(DeviceClientRsaKey, int) bool, or DeviceClientRsaKey) DeviceClientRsaKey {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *DeviceClientRsaKeyStream) FindIndex(fn func(DeviceClientRsaKey, int) bool) int {
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
func (self *DeviceClientRsaKeyStream) First() *DeviceClientRsaKey {
	return self.Get(0)
}
func (self *DeviceClientRsaKeyStream) FirstOr(arg DeviceClientRsaKey) DeviceClientRsaKey {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceClientRsaKeyStream) ForEach(fn func(DeviceClientRsaKey, int)) *DeviceClientRsaKeyStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *DeviceClientRsaKeyStream) ForEachRight(fn func(DeviceClientRsaKey, int)) *DeviceClientRsaKeyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *DeviceClientRsaKeyStream) GroupBy(fn func(DeviceClientRsaKey, int) string) map[string][]DeviceClientRsaKey {
	m := map[string][]DeviceClientRsaKey{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *DeviceClientRsaKeyStream) GroupByValues(fn func(DeviceClientRsaKey, int) string) [][]DeviceClientRsaKey {
	var tmp [][]DeviceClientRsaKey
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *DeviceClientRsaKeyStream) IndexOf(arg DeviceClientRsaKey) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *DeviceClientRsaKeyStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *DeviceClientRsaKeyStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *DeviceClientRsaKeyStream) Last() *DeviceClientRsaKey {
	return self.Get(self.Len() - 1)
}
func (self *DeviceClientRsaKeyStream) LastOr(arg DeviceClientRsaKey) DeviceClientRsaKey {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceClientRsaKeyStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *DeviceClientRsaKeyStream) Limit(limit int) *DeviceClientRsaKeyStream {
	self.Slice(0, limit)
	return self
}

func (self *DeviceClientRsaKeyStream) Map(fn func(DeviceClientRsaKey, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Map2Int(fn func(DeviceClientRsaKey, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Map2Int32(fn func(DeviceClientRsaKey, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Map2Int64(fn func(DeviceClientRsaKey, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Map2Float32(fn func(DeviceClientRsaKey, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Map2Float64(fn func(DeviceClientRsaKey, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Map2Bool(fn func(DeviceClientRsaKey, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Map2Bytes(fn func(DeviceClientRsaKey, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Map2String(fn func(DeviceClientRsaKey, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Max(fn func(DeviceClientRsaKey, int) float64) *DeviceClientRsaKey {
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
func (self *DeviceClientRsaKeyStream) Min(fn func(DeviceClientRsaKey, int) float64) *DeviceClientRsaKey {
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
func (self *DeviceClientRsaKeyStream) NoneMatch(fn func(DeviceClientRsaKey, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *DeviceClientRsaKeyStream) Get(index int) *DeviceClientRsaKey {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *DeviceClientRsaKeyStream) GetOr(index int, arg DeviceClientRsaKey) DeviceClientRsaKey {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceClientRsaKeyStream) Peek(fn func(*DeviceClientRsaKey, int)) *DeviceClientRsaKeyStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *DeviceClientRsaKeyStream) Reduce(fn func(DeviceClientRsaKey, DeviceClientRsaKey, int) DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	return self.ReduceInit(fn, DeviceClientRsaKey{})
}
func (self *DeviceClientRsaKeyStream) ReduceInit(fn func(DeviceClientRsaKey, DeviceClientRsaKey, int) DeviceClientRsaKey, initialValue DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	result := DeviceClientRsaKeyStreamOf()
	self.ForEach(func(v DeviceClientRsaKey, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *DeviceClientRsaKeyStream) ReduceInterface(fn func(interface{}, DeviceClientRsaKey, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(DeviceClientRsaKey{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *DeviceClientRsaKeyStream) ReduceString(fn func(string, DeviceClientRsaKey, int) string) []string {
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
func (self *DeviceClientRsaKeyStream) ReduceInt(fn func(int, DeviceClientRsaKey, int) int) []int {
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
func (self *DeviceClientRsaKeyStream) ReduceInt32(fn func(int32, DeviceClientRsaKey, int) int32) []int32 {
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
func (self *DeviceClientRsaKeyStream) ReduceInt64(fn func(int64, DeviceClientRsaKey, int) int64) []int64 {
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
func (self *DeviceClientRsaKeyStream) ReduceFloat32(fn func(float32, DeviceClientRsaKey, int) float32) []float32 {
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
func (self *DeviceClientRsaKeyStream) ReduceFloat64(fn func(float64, DeviceClientRsaKey, int) float64) []float64 {
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
func (self *DeviceClientRsaKeyStream) ReduceBool(fn func(bool, DeviceClientRsaKey, int) bool) []bool {
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
func (self *DeviceClientRsaKeyStream) Reverse() *DeviceClientRsaKeyStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *DeviceClientRsaKeyStream) Replace(fn func(DeviceClientRsaKey, int) DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	return self.ForEach(func(v DeviceClientRsaKey, i int) { self.Set(i, fn(v, i)) })
}
func (self *DeviceClientRsaKeyStream) Select(fn func(DeviceClientRsaKey) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *DeviceClientRsaKeyStream) Set(index int, val DeviceClientRsaKey) *DeviceClientRsaKeyStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *DeviceClientRsaKeyStream) Skip(skip int) *DeviceClientRsaKeyStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *DeviceClientRsaKeyStream) SkippingEach(fn func(DeviceClientRsaKey, int) int) *DeviceClientRsaKeyStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *DeviceClientRsaKeyStream) Slice(startIndex, n int) *DeviceClientRsaKeyStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []DeviceClientRsaKey{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *DeviceClientRsaKeyStream) Sort(fn func(i, j int) bool) *DeviceClientRsaKeyStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *DeviceClientRsaKeyStream) Tail() *DeviceClientRsaKey {
	return self.Last()
}
func (self *DeviceClientRsaKeyStream) TailOr(arg DeviceClientRsaKey) DeviceClientRsaKey {
	return self.LastOr(arg)
}
func (self *DeviceClientRsaKeyStream) ToList() []DeviceClientRsaKey {
	return self.Val()
}
func (self *DeviceClientRsaKeyStream) Unique() *DeviceClientRsaKeyStream {
	return self.Distinct()
}
func (self *DeviceClientRsaKeyStream) Val() []DeviceClientRsaKey {
	if self == nil {
		return []DeviceClientRsaKey{}
	}
	return *self.Copy()
}
func (self *DeviceClientRsaKeyStream) While(fn func(DeviceClientRsaKey, int) bool) *DeviceClientRsaKeyStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *DeviceClientRsaKeyStream) Where(fn func(DeviceClientRsaKey) bool) *DeviceClientRsaKeyStream {
	result := DeviceClientRsaKeyStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DeviceClientRsaKeyStream) WhereSlim(fn func(DeviceClientRsaKey) bool) *DeviceClientRsaKeyStream {
	result := DeviceClientRsaKeyStreamOf()
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
