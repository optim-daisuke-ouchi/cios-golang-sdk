/*
 * Collection utility of NullableSingleRoute Struct
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

type NullableSingleRouteStream []NullableSingleRoute

func NullableSingleRouteStreamOf(arg ...NullableSingleRoute) NullableSingleRouteStream {
	return arg
}
func NullableSingleRouteStreamFrom(arg []NullableSingleRoute) NullableSingleRouteStream {
	return arg
}
func CreateNullableSingleRouteStream(arg ...NullableSingleRoute) *NullableSingleRouteStream {
	tmp := NullableSingleRouteStreamOf(arg...)
	return &tmp
}
func GenerateNullableSingleRouteStream(arg []NullableSingleRoute) *NullableSingleRouteStream {
	tmp := NullableSingleRouteStreamFrom(arg)
	return &tmp
}

func (self *NullableSingleRouteStream) Add(arg NullableSingleRoute) *NullableSingleRouteStream {
	return self.AddAll(arg)
}
func (self *NullableSingleRouteStream) AddAll(arg ...NullableSingleRoute) *NullableSingleRouteStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSingleRouteStream) AddSafe(arg *NullableSingleRoute) *NullableSingleRouteStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSingleRouteStream) Aggregate(fn func(NullableSingleRoute, NullableSingleRoute) NullableSingleRoute) *NullableSingleRouteStream {
	result := NullableSingleRouteStreamOf()
	self.ForEach(func(v NullableSingleRoute, i int) {
		if i == 0 {
			result.Add(fn(NullableSingleRoute{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleRouteStream) AllMatch(fn func(NullableSingleRoute, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSingleRouteStream) AnyMatch(fn func(NullableSingleRoute, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSingleRouteStream) Clone() *NullableSingleRouteStream {
	temp := make([]NullableSingleRoute, self.Len())
	copy(temp, *self)
	return (*NullableSingleRouteStream)(&temp)
}
func (self *NullableSingleRouteStream) Copy() *NullableSingleRouteStream {
	return self.Clone()
}
func (self *NullableSingleRouteStream) Concat(arg []NullableSingleRoute) *NullableSingleRouteStream {
	return self.AddAll(arg...)
}
func (self *NullableSingleRouteStream) Contains(arg NullableSingleRoute) bool {
	return self.FindIndex(func(_arg NullableSingleRoute, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSingleRouteStream) Clean() *NullableSingleRouteStream {
	*self = NullableSingleRouteStreamOf()
	return self
}
func (self *NullableSingleRouteStream) Delete(index int) *NullableSingleRouteStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSingleRouteStream) DeleteRange(startIndex, endIndex int) *NullableSingleRouteStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSingleRouteStream) Distinct() *NullableSingleRouteStream {
	caches := map[string]bool{}
	result := NullableSingleRouteStreamOf()
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
func (self *NullableSingleRouteStream) Each(fn func(NullableSingleRoute)) *NullableSingleRouteStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSingleRouteStream) EachRight(fn func(NullableSingleRoute)) *NullableSingleRouteStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSingleRouteStream) Equals(arr []NullableSingleRoute) bool {
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
func (self *NullableSingleRouteStream) Filter(fn func(NullableSingleRoute, int) bool) *NullableSingleRouteStream {
	result := NullableSingleRouteStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleRouteStream) FilterSlim(fn func(NullableSingleRoute, int) bool) *NullableSingleRouteStream {
	result := NullableSingleRouteStreamOf()
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
func (self *NullableSingleRouteStream) Find(fn func(NullableSingleRoute, int) bool) *NullableSingleRoute {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSingleRouteStream) FindOr(fn func(NullableSingleRoute, int) bool, or NullableSingleRoute) NullableSingleRoute {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSingleRouteStream) FindIndex(fn func(NullableSingleRoute, int) bool) int {
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
func (self *NullableSingleRouteStream) First() *NullableSingleRoute {
	return self.Get(0)
}
func (self *NullableSingleRouteStream) FirstOr(arg NullableSingleRoute) NullableSingleRoute {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleRouteStream) ForEach(fn func(NullableSingleRoute, int)) *NullableSingleRouteStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSingleRouteStream) ForEachRight(fn func(NullableSingleRoute, int)) *NullableSingleRouteStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSingleRouteStream) GroupBy(fn func(NullableSingleRoute, int) string) map[string][]NullableSingleRoute {
	m := map[string][]NullableSingleRoute{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSingleRouteStream) GroupByValues(fn func(NullableSingleRoute, int) string) [][]NullableSingleRoute {
	var tmp [][]NullableSingleRoute
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSingleRouteStream) IndexOf(arg NullableSingleRoute) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSingleRouteStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSingleRouteStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSingleRouteStream) Last() *NullableSingleRoute {
	return self.Get(self.Len() - 1)
}
func (self *NullableSingleRouteStream) LastOr(arg NullableSingleRoute) NullableSingleRoute {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleRouteStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSingleRouteStream) Limit(limit int) *NullableSingleRouteStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSingleRouteStream) Map(fn func(NullableSingleRoute, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRouteStream) Map2Int(fn func(NullableSingleRoute, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRouteStream) Map2Int32(fn func(NullableSingleRoute, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRouteStream) Map2Int64(fn func(NullableSingleRoute, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRouteStream) Map2Float32(fn func(NullableSingleRoute, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRouteStream) Map2Float64(fn func(NullableSingleRoute, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRouteStream) Map2Bool(fn func(NullableSingleRoute, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRouteStream) Map2Bytes(fn func(NullableSingleRoute, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRouteStream) Map2String(fn func(NullableSingleRoute, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRouteStream) Max(fn func(NullableSingleRoute, int) float64) *NullableSingleRoute {
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
func (self *NullableSingleRouteStream) Min(fn func(NullableSingleRoute, int) float64) *NullableSingleRoute {
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
func (self *NullableSingleRouteStream) NoneMatch(fn func(NullableSingleRoute, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSingleRouteStream) Get(index int) *NullableSingleRoute {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSingleRouteStream) GetOr(index int, arg NullableSingleRoute) NullableSingleRoute {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleRouteStream) Peek(fn func(*NullableSingleRoute, int)) *NullableSingleRouteStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSingleRouteStream) Reduce(fn func(NullableSingleRoute, NullableSingleRoute, int) NullableSingleRoute) *NullableSingleRouteStream {
	return self.ReduceInit(fn, NullableSingleRoute{})
}
func (self *NullableSingleRouteStream) ReduceInit(fn func(NullableSingleRoute, NullableSingleRoute, int) NullableSingleRoute, initialValue NullableSingleRoute) *NullableSingleRouteStream {
	result := NullableSingleRouteStreamOf()
	self.ForEach(func(v NullableSingleRoute, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleRouteStream) ReduceInterface(fn func(interface{}, NullableSingleRoute, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSingleRoute{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSingleRouteStream) ReduceString(fn func(string, NullableSingleRoute, int) string) []string {
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
func (self *NullableSingleRouteStream) ReduceInt(fn func(int, NullableSingleRoute, int) int) []int {
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
func (self *NullableSingleRouteStream) ReduceInt32(fn func(int32, NullableSingleRoute, int) int32) []int32 {
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
func (self *NullableSingleRouteStream) ReduceInt64(fn func(int64, NullableSingleRoute, int) int64) []int64 {
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
func (self *NullableSingleRouteStream) ReduceFloat32(fn func(float32, NullableSingleRoute, int) float32) []float32 {
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
func (self *NullableSingleRouteStream) ReduceFloat64(fn func(float64, NullableSingleRoute, int) float64) []float64 {
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
func (self *NullableSingleRouteStream) ReduceBool(fn func(bool, NullableSingleRoute, int) bool) []bool {
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
func (self *NullableSingleRouteStream) Reverse() *NullableSingleRouteStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSingleRouteStream) Replace(fn func(NullableSingleRoute, int) NullableSingleRoute) *NullableSingleRouteStream {
	return self.ForEach(func(v NullableSingleRoute, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSingleRouteStream) Select(fn func(NullableSingleRoute) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSingleRouteStream) Set(index int, val NullableSingleRoute) *NullableSingleRouteStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSingleRouteStream) Skip(skip int) *NullableSingleRouteStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSingleRouteStream) SkippingEach(fn func(NullableSingleRoute, int) int) *NullableSingleRouteStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSingleRouteStream) Slice(startIndex, n int) *NullableSingleRouteStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSingleRoute{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSingleRouteStream) Sort(fn func(i, j int) bool) *NullableSingleRouteStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSingleRouteStream) Tail() *NullableSingleRoute {
	return self.Last()
}
func (self *NullableSingleRouteStream) TailOr(arg NullableSingleRoute) NullableSingleRoute {
	return self.LastOr(arg)
}
func (self *NullableSingleRouteStream) ToList() []NullableSingleRoute {
	return self.Val()
}
func (self *NullableSingleRouteStream) Unique() *NullableSingleRouteStream {
	return self.Distinct()
}
func (self *NullableSingleRouteStream) Val() []NullableSingleRoute {
	if self == nil {
		return []NullableSingleRoute{}
	}
	return *self.Copy()
}
func (self *NullableSingleRouteStream) While(fn func(NullableSingleRoute, int) bool) *NullableSingleRouteStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSingleRouteStream) Where(fn func(NullableSingleRoute) bool) *NullableSingleRouteStream {
	result := NullableSingleRouteStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleRouteStream) WhereSlim(fn func(NullableSingleRoute) bool) *NullableSingleRouteStream {
	result := NullableSingleRouteStreamOf()
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
