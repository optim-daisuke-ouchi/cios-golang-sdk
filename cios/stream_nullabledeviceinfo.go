/*
 * Collection utility of NullableDeviceInfo Struct
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

type NullableDeviceInfoStream []NullableDeviceInfo

func NullableDeviceInfoStreamOf(arg ...NullableDeviceInfo) NullableDeviceInfoStream {
	return arg
}
func NullableDeviceInfoStreamFrom(arg []NullableDeviceInfo) NullableDeviceInfoStream {
	return arg
}
func CreateNullableDeviceInfoStream(arg ...NullableDeviceInfo) *NullableDeviceInfoStream {
	tmp := NullableDeviceInfoStreamOf(arg...)
	return &tmp
}
func GenerateNullableDeviceInfoStream(arg []NullableDeviceInfo) *NullableDeviceInfoStream {
	tmp := NullableDeviceInfoStreamFrom(arg)
	return &tmp
}

func (self *NullableDeviceInfoStream) Add(arg NullableDeviceInfo) *NullableDeviceInfoStream {
	return self.AddAll(arg)
}
func (self *NullableDeviceInfoStream) AddAll(arg ...NullableDeviceInfo) *NullableDeviceInfoStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableDeviceInfoStream) AddSafe(arg *NullableDeviceInfo) *NullableDeviceInfoStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableDeviceInfoStream) Aggregate(fn func(NullableDeviceInfo, NullableDeviceInfo) NullableDeviceInfo) *NullableDeviceInfoStream {
	result := NullableDeviceInfoStreamOf()
	self.ForEach(func(v NullableDeviceInfo, i int) {
		if i == 0 {
			result.Add(fn(NullableDeviceInfo{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableDeviceInfoStream) AllMatch(fn func(NullableDeviceInfo, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableDeviceInfoStream) AnyMatch(fn func(NullableDeviceInfo, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableDeviceInfoStream) Clone() *NullableDeviceInfoStream {
	temp := make([]NullableDeviceInfo, self.Len())
	copy(temp, *self)
	return (*NullableDeviceInfoStream)(&temp)
}
func (self *NullableDeviceInfoStream) Copy() *NullableDeviceInfoStream {
	return self.Clone()
}
func (self *NullableDeviceInfoStream) Concat(arg []NullableDeviceInfo) *NullableDeviceInfoStream {
	return self.AddAll(arg...)
}
func (self *NullableDeviceInfoStream) Contains(arg NullableDeviceInfo) bool {
	return self.FindIndex(func(_arg NullableDeviceInfo, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableDeviceInfoStream) Clean() *NullableDeviceInfoStream {
	*self = NullableDeviceInfoStreamOf()
	return self
}
func (self *NullableDeviceInfoStream) Delete(index int) *NullableDeviceInfoStream {
	return self.DeleteRange(index, index)
}
func (self *NullableDeviceInfoStream) DeleteRange(startIndex, endIndex int) *NullableDeviceInfoStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableDeviceInfoStream) Distinct() *NullableDeviceInfoStream {
	caches := map[string]bool{}
	result := NullableDeviceInfoStreamOf()
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
func (self *NullableDeviceInfoStream) Each(fn func(NullableDeviceInfo)) *NullableDeviceInfoStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableDeviceInfoStream) EachRight(fn func(NullableDeviceInfo)) *NullableDeviceInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableDeviceInfoStream) Equals(arr []NullableDeviceInfo) bool {
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
func (self *NullableDeviceInfoStream) Filter(fn func(NullableDeviceInfo, int) bool) *NullableDeviceInfoStream {
	result := NullableDeviceInfoStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDeviceInfoStream) FilterSlim(fn func(NullableDeviceInfo, int) bool) *NullableDeviceInfoStream {
	result := NullableDeviceInfoStreamOf()
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
func (self *NullableDeviceInfoStream) Find(fn func(NullableDeviceInfo, int) bool) *NullableDeviceInfo {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableDeviceInfoStream) FindOr(fn func(NullableDeviceInfo, int) bool, or NullableDeviceInfo) NullableDeviceInfo {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableDeviceInfoStream) FindIndex(fn func(NullableDeviceInfo, int) bool) int {
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
func (self *NullableDeviceInfoStream) First() *NullableDeviceInfo {
	return self.Get(0)
}
func (self *NullableDeviceInfoStream) FirstOr(arg NullableDeviceInfo) NullableDeviceInfo {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceInfoStream) ForEach(fn func(NullableDeviceInfo, int)) *NullableDeviceInfoStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableDeviceInfoStream) ForEachRight(fn func(NullableDeviceInfo, int)) *NullableDeviceInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableDeviceInfoStream) GroupBy(fn func(NullableDeviceInfo, int) string) map[string][]NullableDeviceInfo {
	m := map[string][]NullableDeviceInfo{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableDeviceInfoStream) GroupByValues(fn func(NullableDeviceInfo, int) string) [][]NullableDeviceInfo {
	var tmp [][]NullableDeviceInfo
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableDeviceInfoStream) IndexOf(arg NullableDeviceInfo) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableDeviceInfoStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableDeviceInfoStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableDeviceInfoStream) Last() *NullableDeviceInfo {
	return self.Get(self.Len() - 1)
}
func (self *NullableDeviceInfoStream) LastOr(arg NullableDeviceInfo) NullableDeviceInfo {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceInfoStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableDeviceInfoStream) Limit(limit int) *NullableDeviceInfoStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableDeviceInfoStream) Map(fn func(NullableDeviceInfo, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Map2Int(fn func(NullableDeviceInfo, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Map2Int32(fn func(NullableDeviceInfo, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Map2Int64(fn func(NullableDeviceInfo, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Map2Float32(fn func(NullableDeviceInfo, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Map2Float64(fn func(NullableDeviceInfo, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Map2Bool(fn func(NullableDeviceInfo, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Map2Bytes(fn func(NullableDeviceInfo, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Map2String(fn func(NullableDeviceInfo, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Max(fn func(NullableDeviceInfo, int) float64) *NullableDeviceInfo {
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
func (self *NullableDeviceInfoStream) Min(fn func(NullableDeviceInfo, int) float64) *NullableDeviceInfo {
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
func (self *NullableDeviceInfoStream) NoneMatch(fn func(NullableDeviceInfo, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableDeviceInfoStream) Get(index int) *NullableDeviceInfo {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableDeviceInfoStream) GetOr(index int, arg NullableDeviceInfo) NullableDeviceInfo {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceInfoStream) Peek(fn func(*NullableDeviceInfo, int)) *NullableDeviceInfoStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableDeviceInfoStream) Reduce(fn func(NullableDeviceInfo, NullableDeviceInfo, int) NullableDeviceInfo) *NullableDeviceInfoStream {
	return self.ReduceInit(fn, NullableDeviceInfo{})
}
func (self *NullableDeviceInfoStream) ReduceInit(fn func(NullableDeviceInfo, NullableDeviceInfo, int) NullableDeviceInfo, initialValue NullableDeviceInfo) *NullableDeviceInfoStream {
	result := NullableDeviceInfoStreamOf()
	self.ForEach(func(v NullableDeviceInfo, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableDeviceInfoStream) ReduceInterface(fn func(interface{}, NullableDeviceInfo, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableDeviceInfo{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableDeviceInfoStream) ReduceString(fn func(string, NullableDeviceInfo, int) string) []string {
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
func (self *NullableDeviceInfoStream) ReduceInt(fn func(int, NullableDeviceInfo, int) int) []int {
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
func (self *NullableDeviceInfoStream) ReduceInt32(fn func(int32, NullableDeviceInfo, int) int32) []int32 {
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
func (self *NullableDeviceInfoStream) ReduceInt64(fn func(int64, NullableDeviceInfo, int) int64) []int64 {
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
func (self *NullableDeviceInfoStream) ReduceFloat32(fn func(float32, NullableDeviceInfo, int) float32) []float32 {
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
func (self *NullableDeviceInfoStream) ReduceFloat64(fn func(float64, NullableDeviceInfo, int) float64) []float64 {
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
func (self *NullableDeviceInfoStream) ReduceBool(fn func(bool, NullableDeviceInfo, int) bool) []bool {
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
func (self *NullableDeviceInfoStream) Reverse() *NullableDeviceInfoStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableDeviceInfoStream) Replace(fn func(NullableDeviceInfo, int) NullableDeviceInfo) *NullableDeviceInfoStream {
	return self.ForEach(func(v NullableDeviceInfo, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableDeviceInfoStream) Select(fn func(NullableDeviceInfo) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableDeviceInfoStream) Set(index int, val NullableDeviceInfo) *NullableDeviceInfoStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableDeviceInfoStream) Skip(skip int) *NullableDeviceInfoStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableDeviceInfoStream) SkippingEach(fn func(NullableDeviceInfo, int) int) *NullableDeviceInfoStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableDeviceInfoStream) Slice(startIndex, n int) *NullableDeviceInfoStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableDeviceInfo{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableDeviceInfoStream) Sort(fn func(i, j int) bool) *NullableDeviceInfoStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableDeviceInfoStream) Tail() *NullableDeviceInfo {
	return self.Last()
}
func (self *NullableDeviceInfoStream) TailOr(arg NullableDeviceInfo) NullableDeviceInfo {
	return self.LastOr(arg)
}
func (self *NullableDeviceInfoStream) ToList() []NullableDeviceInfo {
	return self.Val()
}
func (self *NullableDeviceInfoStream) Unique() *NullableDeviceInfoStream {
	return self.Distinct()
}
func (self *NullableDeviceInfoStream) Val() []NullableDeviceInfo {
	if self == nil {
		return []NullableDeviceInfo{}
	}
	return *self.Copy()
}
func (self *NullableDeviceInfoStream) While(fn func(NullableDeviceInfo, int) bool) *NullableDeviceInfoStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableDeviceInfoStream) Where(fn func(NullableDeviceInfo) bool) *NullableDeviceInfoStream {
	result := NullableDeviceInfoStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDeviceInfoStream) WhereSlim(fn func(NullableDeviceInfo) bool) *NullableDeviceInfoStream {
	result := NullableDeviceInfoStreamOf()
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
