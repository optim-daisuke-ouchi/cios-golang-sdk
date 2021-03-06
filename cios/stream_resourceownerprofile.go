/*
 * Collection utility of ResourceOwnerProfile Struct
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

type ResourceOwnerProfileStream []ResourceOwnerProfile

func ResourceOwnerProfileStreamOf(arg ...ResourceOwnerProfile) ResourceOwnerProfileStream {
	return arg
}
func ResourceOwnerProfileStreamFrom(arg []ResourceOwnerProfile) ResourceOwnerProfileStream {
	return arg
}
func CreateResourceOwnerProfileStream(arg ...ResourceOwnerProfile) *ResourceOwnerProfileStream {
	tmp := ResourceOwnerProfileStreamOf(arg...)
	return &tmp
}
func GenerateResourceOwnerProfileStream(arg []ResourceOwnerProfile) *ResourceOwnerProfileStream {
	tmp := ResourceOwnerProfileStreamFrom(arg)
	return &tmp
}

func (self *ResourceOwnerProfileStream) Add(arg ResourceOwnerProfile) *ResourceOwnerProfileStream {
	return self.AddAll(arg)
}
func (self *ResourceOwnerProfileStream) AddAll(arg ...ResourceOwnerProfile) *ResourceOwnerProfileStream {
	*self = append(*self, arg...)
	return self
}
func (self *ResourceOwnerProfileStream) AddSafe(arg *ResourceOwnerProfile) *ResourceOwnerProfileStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ResourceOwnerProfileStream) Aggregate(fn func(ResourceOwnerProfile, ResourceOwnerProfile) ResourceOwnerProfile) *ResourceOwnerProfileStream {
	result := ResourceOwnerProfileStreamOf()
	self.ForEach(func(v ResourceOwnerProfile, i int) {
		if i == 0 {
			result.Add(fn(ResourceOwnerProfile{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ResourceOwnerProfileStream) AllMatch(fn func(ResourceOwnerProfile, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ResourceOwnerProfileStream) AnyMatch(fn func(ResourceOwnerProfile, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ResourceOwnerProfileStream) Clone() *ResourceOwnerProfileStream {
	temp := make([]ResourceOwnerProfile, self.Len())
	copy(temp, *self)
	return (*ResourceOwnerProfileStream)(&temp)
}
func (self *ResourceOwnerProfileStream) Copy() *ResourceOwnerProfileStream {
	return self.Clone()
}
func (self *ResourceOwnerProfileStream) Concat(arg []ResourceOwnerProfile) *ResourceOwnerProfileStream {
	return self.AddAll(arg...)
}
func (self *ResourceOwnerProfileStream) Contains(arg ResourceOwnerProfile) bool {
	return self.FindIndex(func(_arg ResourceOwnerProfile, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ResourceOwnerProfileStream) Clean() *ResourceOwnerProfileStream {
	*self = ResourceOwnerProfileStreamOf()
	return self
}
func (self *ResourceOwnerProfileStream) Delete(index int) *ResourceOwnerProfileStream {
	return self.DeleteRange(index, index)
}
func (self *ResourceOwnerProfileStream) DeleteRange(startIndex, endIndex int) *ResourceOwnerProfileStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ResourceOwnerProfileStream) Distinct() *ResourceOwnerProfileStream {
	caches := map[string]bool{}
	result := ResourceOwnerProfileStreamOf()
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
func (self *ResourceOwnerProfileStream) Each(fn func(ResourceOwnerProfile)) *ResourceOwnerProfileStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ResourceOwnerProfileStream) EachRight(fn func(ResourceOwnerProfile)) *ResourceOwnerProfileStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ResourceOwnerProfileStream) Equals(arr []ResourceOwnerProfile) bool {
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
func (self *ResourceOwnerProfileStream) Filter(fn func(ResourceOwnerProfile, int) bool) *ResourceOwnerProfileStream {
	result := ResourceOwnerProfileStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ResourceOwnerProfileStream) FilterSlim(fn func(ResourceOwnerProfile, int) bool) *ResourceOwnerProfileStream {
	result := ResourceOwnerProfileStreamOf()
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
func (self *ResourceOwnerProfileStream) Find(fn func(ResourceOwnerProfile, int) bool) *ResourceOwnerProfile {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ResourceOwnerProfileStream) FindOr(fn func(ResourceOwnerProfile, int) bool, or ResourceOwnerProfile) ResourceOwnerProfile {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ResourceOwnerProfileStream) FindIndex(fn func(ResourceOwnerProfile, int) bool) int {
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
func (self *ResourceOwnerProfileStream) First() *ResourceOwnerProfile {
	return self.Get(0)
}
func (self *ResourceOwnerProfileStream) FirstOr(arg ResourceOwnerProfile) ResourceOwnerProfile {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ResourceOwnerProfileStream) ForEach(fn func(ResourceOwnerProfile, int)) *ResourceOwnerProfileStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ResourceOwnerProfileStream) ForEachRight(fn func(ResourceOwnerProfile, int)) *ResourceOwnerProfileStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ResourceOwnerProfileStream) GroupBy(fn func(ResourceOwnerProfile, int) string) map[string][]ResourceOwnerProfile {
	m := map[string][]ResourceOwnerProfile{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ResourceOwnerProfileStream) GroupByValues(fn func(ResourceOwnerProfile, int) string) [][]ResourceOwnerProfile {
	var tmp [][]ResourceOwnerProfile
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ResourceOwnerProfileStream) IndexOf(arg ResourceOwnerProfile) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ResourceOwnerProfileStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ResourceOwnerProfileStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ResourceOwnerProfileStream) Last() *ResourceOwnerProfile {
	return self.Get(self.Len() - 1)
}
func (self *ResourceOwnerProfileStream) LastOr(arg ResourceOwnerProfile) ResourceOwnerProfile {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ResourceOwnerProfileStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ResourceOwnerProfileStream) Limit(limit int) *ResourceOwnerProfileStream {
	self.Slice(0, limit)
	return self
}

func (self *ResourceOwnerProfileStream) Map(fn func(ResourceOwnerProfile, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Map2Int(fn func(ResourceOwnerProfile, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Map2Int32(fn func(ResourceOwnerProfile, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Map2Int64(fn func(ResourceOwnerProfile, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Map2Float32(fn func(ResourceOwnerProfile, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Map2Float64(fn func(ResourceOwnerProfile, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Map2Bool(fn func(ResourceOwnerProfile, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Map2Bytes(fn func(ResourceOwnerProfile, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Map2String(fn func(ResourceOwnerProfile, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Max(fn func(ResourceOwnerProfile, int) float64) *ResourceOwnerProfile {
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
func (self *ResourceOwnerProfileStream) Min(fn func(ResourceOwnerProfile, int) float64) *ResourceOwnerProfile {
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
func (self *ResourceOwnerProfileStream) NoneMatch(fn func(ResourceOwnerProfile, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ResourceOwnerProfileStream) Get(index int) *ResourceOwnerProfile {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ResourceOwnerProfileStream) GetOr(index int, arg ResourceOwnerProfile) ResourceOwnerProfile {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ResourceOwnerProfileStream) Peek(fn func(*ResourceOwnerProfile, int)) *ResourceOwnerProfileStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ResourceOwnerProfileStream) Reduce(fn func(ResourceOwnerProfile, ResourceOwnerProfile, int) ResourceOwnerProfile) *ResourceOwnerProfileStream {
	return self.ReduceInit(fn, ResourceOwnerProfile{})
}
func (self *ResourceOwnerProfileStream) ReduceInit(fn func(ResourceOwnerProfile, ResourceOwnerProfile, int) ResourceOwnerProfile, initialValue ResourceOwnerProfile) *ResourceOwnerProfileStream {
	result := ResourceOwnerProfileStreamOf()
	self.ForEach(func(v ResourceOwnerProfile, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ResourceOwnerProfileStream) ReduceInterface(fn func(interface{}, ResourceOwnerProfile, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ResourceOwnerProfile{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ResourceOwnerProfileStream) ReduceString(fn func(string, ResourceOwnerProfile, int) string) []string {
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
func (self *ResourceOwnerProfileStream) ReduceInt(fn func(int, ResourceOwnerProfile, int) int) []int {
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
func (self *ResourceOwnerProfileStream) ReduceInt32(fn func(int32, ResourceOwnerProfile, int) int32) []int32 {
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
func (self *ResourceOwnerProfileStream) ReduceInt64(fn func(int64, ResourceOwnerProfile, int) int64) []int64 {
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
func (self *ResourceOwnerProfileStream) ReduceFloat32(fn func(float32, ResourceOwnerProfile, int) float32) []float32 {
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
func (self *ResourceOwnerProfileStream) ReduceFloat64(fn func(float64, ResourceOwnerProfile, int) float64) []float64 {
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
func (self *ResourceOwnerProfileStream) ReduceBool(fn func(bool, ResourceOwnerProfile, int) bool) []bool {
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
func (self *ResourceOwnerProfileStream) Reverse() *ResourceOwnerProfileStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ResourceOwnerProfileStream) Replace(fn func(ResourceOwnerProfile, int) ResourceOwnerProfile) *ResourceOwnerProfileStream {
	return self.ForEach(func(v ResourceOwnerProfile, i int) { self.Set(i, fn(v, i)) })
}
func (self *ResourceOwnerProfileStream) Select(fn func(ResourceOwnerProfile) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ResourceOwnerProfileStream) Set(index int, val ResourceOwnerProfile) *ResourceOwnerProfileStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ResourceOwnerProfileStream) Skip(skip int) *ResourceOwnerProfileStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ResourceOwnerProfileStream) SkippingEach(fn func(ResourceOwnerProfile, int) int) *ResourceOwnerProfileStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ResourceOwnerProfileStream) Slice(startIndex, n int) *ResourceOwnerProfileStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ResourceOwnerProfile{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ResourceOwnerProfileStream) Sort(fn func(i, j int) bool) *ResourceOwnerProfileStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ResourceOwnerProfileStream) Tail() *ResourceOwnerProfile {
	return self.Last()
}
func (self *ResourceOwnerProfileStream) TailOr(arg ResourceOwnerProfile) ResourceOwnerProfile {
	return self.LastOr(arg)
}
func (self *ResourceOwnerProfileStream) ToList() []ResourceOwnerProfile {
	return self.Val()
}
func (self *ResourceOwnerProfileStream) Unique() *ResourceOwnerProfileStream {
	return self.Distinct()
}
func (self *ResourceOwnerProfileStream) Val() []ResourceOwnerProfile {
	if self == nil {
		return []ResourceOwnerProfile{}
	}
	return *self.Copy()
}
func (self *ResourceOwnerProfileStream) While(fn func(ResourceOwnerProfile, int) bool) *ResourceOwnerProfileStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ResourceOwnerProfileStream) Where(fn func(ResourceOwnerProfile) bool) *ResourceOwnerProfileStream {
	result := ResourceOwnerProfileStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ResourceOwnerProfileStream) WhereSlim(fn func(ResourceOwnerProfile) bool) *ResourceOwnerProfileStream {
	result := ResourceOwnerProfileStreamOf()
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
