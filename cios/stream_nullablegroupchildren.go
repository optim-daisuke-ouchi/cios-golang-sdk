/*
 * Collection utility of NullableGroupChildren Struct
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

type NullableGroupChildrenStream []NullableGroupChildren

func NullableGroupChildrenStreamOf(arg ...NullableGroupChildren) NullableGroupChildrenStream {
	return arg
}
func NullableGroupChildrenStreamFrom(arg []NullableGroupChildren) NullableGroupChildrenStream {
	return arg
}
func CreateNullableGroupChildrenStream(arg ...NullableGroupChildren) *NullableGroupChildrenStream {
	tmp := NullableGroupChildrenStreamOf(arg...)
	return &tmp
}
func GenerateNullableGroupChildrenStream(arg []NullableGroupChildren) *NullableGroupChildrenStream {
	tmp := NullableGroupChildrenStreamFrom(arg)
	return &tmp
}

func (self *NullableGroupChildrenStream) Add(arg NullableGroupChildren) *NullableGroupChildrenStream {
	return self.AddAll(arg)
}
func (self *NullableGroupChildrenStream) AddAll(arg ...NullableGroupChildren) *NullableGroupChildrenStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableGroupChildrenStream) AddSafe(arg *NullableGroupChildren) *NullableGroupChildrenStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableGroupChildrenStream) Aggregate(fn func(NullableGroupChildren, NullableGroupChildren) NullableGroupChildren) *NullableGroupChildrenStream {
	result := NullableGroupChildrenStreamOf()
	self.ForEach(func(v NullableGroupChildren, i int) {
		if i == 0 {
			result.Add(fn(NullableGroupChildren{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableGroupChildrenStream) AllMatch(fn func(NullableGroupChildren, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableGroupChildrenStream) AnyMatch(fn func(NullableGroupChildren, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableGroupChildrenStream) Clone() *NullableGroupChildrenStream {
	temp := make([]NullableGroupChildren, self.Len())
	copy(temp, *self)
	return (*NullableGroupChildrenStream)(&temp)
}
func (self *NullableGroupChildrenStream) Copy() *NullableGroupChildrenStream {
	return self.Clone()
}
func (self *NullableGroupChildrenStream) Concat(arg []NullableGroupChildren) *NullableGroupChildrenStream {
	return self.AddAll(arg...)
}
func (self *NullableGroupChildrenStream) Contains(arg NullableGroupChildren) bool {
	return self.FindIndex(func(_arg NullableGroupChildren, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableGroupChildrenStream) Clean() *NullableGroupChildrenStream {
	*self = NullableGroupChildrenStreamOf()
	return self
}
func (self *NullableGroupChildrenStream) Delete(index int) *NullableGroupChildrenStream {
	return self.DeleteRange(index, index)
}
func (self *NullableGroupChildrenStream) DeleteRange(startIndex, endIndex int) *NullableGroupChildrenStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableGroupChildrenStream) Distinct() *NullableGroupChildrenStream {
	caches := map[string]bool{}
	result := NullableGroupChildrenStreamOf()
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
func (self *NullableGroupChildrenStream) Each(fn func(NullableGroupChildren)) *NullableGroupChildrenStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableGroupChildrenStream) EachRight(fn func(NullableGroupChildren)) *NullableGroupChildrenStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableGroupChildrenStream) Equals(arr []NullableGroupChildren) bool {
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
func (self *NullableGroupChildrenStream) Filter(fn func(NullableGroupChildren, int) bool) *NullableGroupChildrenStream {
	result := NullableGroupChildrenStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableGroupChildrenStream) FilterSlim(fn func(NullableGroupChildren, int) bool) *NullableGroupChildrenStream {
	result := NullableGroupChildrenStreamOf()
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
func (self *NullableGroupChildrenStream) Find(fn func(NullableGroupChildren, int) bool) *NullableGroupChildren {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableGroupChildrenStream) FindOr(fn func(NullableGroupChildren, int) bool, or NullableGroupChildren) NullableGroupChildren {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableGroupChildrenStream) FindIndex(fn func(NullableGroupChildren, int) bool) int {
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
func (self *NullableGroupChildrenStream) First() *NullableGroupChildren {
	return self.Get(0)
}
func (self *NullableGroupChildrenStream) FirstOr(arg NullableGroupChildren) NullableGroupChildren {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableGroupChildrenStream) ForEach(fn func(NullableGroupChildren, int)) *NullableGroupChildrenStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableGroupChildrenStream) ForEachRight(fn func(NullableGroupChildren, int)) *NullableGroupChildrenStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableGroupChildrenStream) GroupBy(fn func(NullableGroupChildren, int) string) map[string][]NullableGroupChildren {
	m := map[string][]NullableGroupChildren{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableGroupChildrenStream) GroupByValues(fn func(NullableGroupChildren, int) string) [][]NullableGroupChildren {
	var tmp [][]NullableGroupChildren
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableGroupChildrenStream) IndexOf(arg NullableGroupChildren) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableGroupChildrenStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableGroupChildrenStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableGroupChildrenStream) Last() *NullableGroupChildren {
	return self.Get(self.Len() - 1)
}
func (self *NullableGroupChildrenStream) LastOr(arg NullableGroupChildren) NullableGroupChildren {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableGroupChildrenStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableGroupChildrenStream) Limit(limit int) *NullableGroupChildrenStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableGroupChildrenStream) Map(fn func(NullableGroupChildren, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Map2Int(fn func(NullableGroupChildren, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Map2Int32(fn func(NullableGroupChildren, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Map2Int64(fn func(NullableGroupChildren, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Map2Float32(fn func(NullableGroupChildren, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Map2Float64(fn func(NullableGroupChildren, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Map2Bool(fn func(NullableGroupChildren, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Map2Bytes(fn func(NullableGroupChildren, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Map2String(fn func(NullableGroupChildren, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Max(fn func(NullableGroupChildren, int) float64) *NullableGroupChildren {
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
func (self *NullableGroupChildrenStream) Min(fn func(NullableGroupChildren, int) float64) *NullableGroupChildren {
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
func (self *NullableGroupChildrenStream) NoneMatch(fn func(NullableGroupChildren, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableGroupChildrenStream) Get(index int) *NullableGroupChildren {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableGroupChildrenStream) GetOr(index int, arg NullableGroupChildren) NullableGroupChildren {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableGroupChildrenStream) Peek(fn func(*NullableGroupChildren, int)) *NullableGroupChildrenStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableGroupChildrenStream) Reduce(fn func(NullableGroupChildren, NullableGroupChildren, int) NullableGroupChildren) *NullableGroupChildrenStream {
	return self.ReduceInit(fn, NullableGroupChildren{})
}
func (self *NullableGroupChildrenStream) ReduceInit(fn func(NullableGroupChildren, NullableGroupChildren, int) NullableGroupChildren, initialValue NullableGroupChildren) *NullableGroupChildrenStream {
	result := NullableGroupChildrenStreamOf()
	self.ForEach(func(v NullableGroupChildren, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableGroupChildrenStream) ReduceInterface(fn func(interface{}, NullableGroupChildren, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableGroupChildren{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableGroupChildrenStream) ReduceString(fn func(string, NullableGroupChildren, int) string) []string {
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
func (self *NullableGroupChildrenStream) ReduceInt(fn func(int, NullableGroupChildren, int) int) []int {
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
func (self *NullableGroupChildrenStream) ReduceInt32(fn func(int32, NullableGroupChildren, int) int32) []int32 {
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
func (self *NullableGroupChildrenStream) ReduceInt64(fn func(int64, NullableGroupChildren, int) int64) []int64 {
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
func (self *NullableGroupChildrenStream) ReduceFloat32(fn func(float32, NullableGroupChildren, int) float32) []float32 {
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
func (self *NullableGroupChildrenStream) ReduceFloat64(fn func(float64, NullableGroupChildren, int) float64) []float64 {
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
func (self *NullableGroupChildrenStream) ReduceBool(fn func(bool, NullableGroupChildren, int) bool) []bool {
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
func (self *NullableGroupChildrenStream) Reverse() *NullableGroupChildrenStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableGroupChildrenStream) Replace(fn func(NullableGroupChildren, int) NullableGroupChildren) *NullableGroupChildrenStream {
	return self.ForEach(func(v NullableGroupChildren, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableGroupChildrenStream) Select(fn func(NullableGroupChildren) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableGroupChildrenStream) Set(index int, val NullableGroupChildren) *NullableGroupChildrenStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableGroupChildrenStream) Skip(skip int) *NullableGroupChildrenStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableGroupChildrenStream) SkippingEach(fn func(NullableGroupChildren, int) int) *NullableGroupChildrenStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableGroupChildrenStream) Slice(startIndex, n int) *NullableGroupChildrenStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableGroupChildren{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableGroupChildrenStream) Sort(fn func(i, j int) bool) *NullableGroupChildrenStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableGroupChildrenStream) Tail() *NullableGroupChildren {
	return self.Last()
}
func (self *NullableGroupChildrenStream) TailOr(arg NullableGroupChildren) NullableGroupChildren {
	return self.LastOr(arg)
}
func (self *NullableGroupChildrenStream) ToList() []NullableGroupChildren {
	return self.Val()
}
func (self *NullableGroupChildrenStream) Unique() *NullableGroupChildrenStream {
	return self.Distinct()
}
func (self *NullableGroupChildrenStream) Val() []NullableGroupChildren {
	if self == nil {
		return []NullableGroupChildren{}
	}
	return *self.Copy()
}
func (self *NullableGroupChildrenStream) While(fn func(NullableGroupChildren, int) bool) *NullableGroupChildrenStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableGroupChildrenStream) Where(fn func(NullableGroupChildren) bool) *NullableGroupChildrenStream {
	result := NullableGroupChildrenStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableGroupChildrenStream) WhereSlim(fn func(NullableGroupChildren) bool) *NullableGroupChildrenStream {
	result := NullableGroupChildrenStreamOf()
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
