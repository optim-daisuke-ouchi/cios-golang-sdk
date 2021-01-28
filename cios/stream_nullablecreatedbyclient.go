/*
 * Collection utility of NullableCreatedByClient Struct
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

type NullableCreatedByClientStream []NullableCreatedByClient

func NullableCreatedByClientStreamOf(arg ...NullableCreatedByClient) NullableCreatedByClientStream {
	return arg
}
func NullableCreatedByClientStreamFrom(arg []NullableCreatedByClient) NullableCreatedByClientStream {
	return arg
}
func CreateNullableCreatedByClientStream(arg ...NullableCreatedByClient) *NullableCreatedByClientStream {
	tmp := NullableCreatedByClientStreamOf(arg...)
	return &tmp
}
func GenerateNullableCreatedByClientStream(arg []NullableCreatedByClient) *NullableCreatedByClientStream {
	tmp := NullableCreatedByClientStreamFrom(arg)
	return &tmp
}

func (self *NullableCreatedByClientStream) Add(arg NullableCreatedByClient) *NullableCreatedByClientStream {
	return self.AddAll(arg)
}
func (self *NullableCreatedByClientStream) AddAll(arg ...NullableCreatedByClient) *NullableCreatedByClientStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableCreatedByClientStream) AddSafe(arg *NullableCreatedByClient) *NullableCreatedByClientStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableCreatedByClientStream) Aggregate(fn func(NullableCreatedByClient, NullableCreatedByClient) NullableCreatedByClient) *NullableCreatedByClientStream {
	result := NullableCreatedByClientStreamOf()
	self.ForEach(func(v NullableCreatedByClient, i int) {
		if i == 0 {
			result.Add(fn(NullableCreatedByClient{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableCreatedByClientStream) AllMatch(fn func(NullableCreatedByClient, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableCreatedByClientStream) AnyMatch(fn func(NullableCreatedByClient, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableCreatedByClientStream) Clone() *NullableCreatedByClientStream {
	temp := make([]NullableCreatedByClient, self.Len())
	copy(temp, *self)
	return (*NullableCreatedByClientStream)(&temp)
}
func (self *NullableCreatedByClientStream) Copy() *NullableCreatedByClientStream {
	return self.Clone()
}
func (self *NullableCreatedByClientStream) Concat(arg []NullableCreatedByClient) *NullableCreatedByClientStream {
	return self.AddAll(arg...)
}
func (self *NullableCreatedByClientStream) Contains(arg NullableCreatedByClient) bool {
	return self.FindIndex(func(_arg NullableCreatedByClient, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableCreatedByClientStream) Clean() *NullableCreatedByClientStream {
	*self = NullableCreatedByClientStreamOf()
	return self
}
func (self *NullableCreatedByClientStream) Delete(index int) *NullableCreatedByClientStream {
	return self.DeleteRange(index, index)
}
func (self *NullableCreatedByClientStream) DeleteRange(startIndex, endIndex int) *NullableCreatedByClientStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableCreatedByClientStream) Distinct() *NullableCreatedByClientStream {
	caches := map[string]bool{}
	result := NullableCreatedByClientStreamOf()
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
func (self *NullableCreatedByClientStream) Each(fn func(NullableCreatedByClient)) *NullableCreatedByClientStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableCreatedByClientStream) EachRight(fn func(NullableCreatedByClient)) *NullableCreatedByClientStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableCreatedByClientStream) Equals(arr []NullableCreatedByClient) bool {
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
func (self *NullableCreatedByClientStream) Filter(fn func(NullableCreatedByClient, int) bool) *NullableCreatedByClientStream {
	result := NullableCreatedByClientStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableCreatedByClientStream) FilterSlim(fn func(NullableCreatedByClient, int) bool) *NullableCreatedByClientStream {
	result := NullableCreatedByClientStreamOf()
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
func (self *NullableCreatedByClientStream) Find(fn func(NullableCreatedByClient, int) bool) *NullableCreatedByClient {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableCreatedByClientStream) FindOr(fn func(NullableCreatedByClient, int) bool, or NullableCreatedByClient) NullableCreatedByClient {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableCreatedByClientStream) FindIndex(fn func(NullableCreatedByClient, int) bool) int {
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
func (self *NullableCreatedByClientStream) First() *NullableCreatedByClient {
	return self.Get(0)
}
func (self *NullableCreatedByClientStream) FirstOr(arg NullableCreatedByClient) NullableCreatedByClient {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableCreatedByClientStream) ForEach(fn func(NullableCreatedByClient, int)) *NullableCreatedByClientStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableCreatedByClientStream) ForEachRight(fn func(NullableCreatedByClient, int)) *NullableCreatedByClientStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableCreatedByClientStream) GroupBy(fn func(NullableCreatedByClient, int) string) map[string][]NullableCreatedByClient {
	m := map[string][]NullableCreatedByClient{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableCreatedByClientStream) GroupByValues(fn func(NullableCreatedByClient, int) string) [][]NullableCreatedByClient {
	var tmp [][]NullableCreatedByClient
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableCreatedByClientStream) IndexOf(arg NullableCreatedByClient) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableCreatedByClientStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableCreatedByClientStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableCreatedByClientStream) Last() *NullableCreatedByClient {
	return self.Get(self.Len() - 1)
}
func (self *NullableCreatedByClientStream) LastOr(arg NullableCreatedByClient) NullableCreatedByClient {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableCreatedByClientStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableCreatedByClientStream) Limit(limit int) *NullableCreatedByClientStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableCreatedByClientStream) Map(fn func(NullableCreatedByClient, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Map2Int(fn func(NullableCreatedByClient, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Map2Int32(fn func(NullableCreatedByClient, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Map2Int64(fn func(NullableCreatedByClient, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Map2Float32(fn func(NullableCreatedByClient, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Map2Float64(fn func(NullableCreatedByClient, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Map2Bool(fn func(NullableCreatedByClient, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Map2Bytes(fn func(NullableCreatedByClient, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Map2String(fn func(NullableCreatedByClient, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Max(fn func(NullableCreatedByClient, int) float64) *NullableCreatedByClient {
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
func (self *NullableCreatedByClientStream) Min(fn func(NullableCreatedByClient, int) float64) *NullableCreatedByClient {
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
func (self *NullableCreatedByClientStream) NoneMatch(fn func(NullableCreatedByClient, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableCreatedByClientStream) Get(index int) *NullableCreatedByClient {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableCreatedByClientStream) GetOr(index int, arg NullableCreatedByClient) NullableCreatedByClient {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableCreatedByClientStream) Peek(fn func(*NullableCreatedByClient, int)) *NullableCreatedByClientStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableCreatedByClientStream) Reduce(fn func(NullableCreatedByClient, NullableCreatedByClient, int) NullableCreatedByClient) *NullableCreatedByClientStream {
	return self.ReduceInit(fn, NullableCreatedByClient{})
}
func (self *NullableCreatedByClientStream) ReduceInit(fn func(NullableCreatedByClient, NullableCreatedByClient, int) NullableCreatedByClient, initialValue NullableCreatedByClient) *NullableCreatedByClientStream {
	result := NullableCreatedByClientStreamOf()
	self.ForEach(func(v NullableCreatedByClient, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableCreatedByClientStream) ReduceInterface(fn func(interface{}, NullableCreatedByClient, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableCreatedByClient{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableCreatedByClientStream) ReduceString(fn func(string, NullableCreatedByClient, int) string) []string {
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
func (self *NullableCreatedByClientStream) ReduceInt(fn func(int, NullableCreatedByClient, int) int) []int {
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
func (self *NullableCreatedByClientStream) ReduceInt32(fn func(int32, NullableCreatedByClient, int) int32) []int32 {
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
func (self *NullableCreatedByClientStream) ReduceInt64(fn func(int64, NullableCreatedByClient, int) int64) []int64 {
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
func (self *NullableCreatedByClientStream) ReduceFloat32(fn func(float32, NullableCreatedByClient, int) float32) []float32 {
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
func (self *NullableCreatedByClientStream) ReduceFloat64(fn func(float64, NullableCreatedByClient, int) float64) []float64 {
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
func (self *NullableCreatedByClientStream) ReduceBool(fn func(bool, NullableCreatedByClient, int) bool) []bool {
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
func (self *NullableCreatedByClientStream) Reverse() *NullableCreatedByClientStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableCreatedByClientStream) Replace(fn func(NullableCreatedByClient, int) NullableCreatedByClient) *NullableCreatedByClientStream {
	return self.ForEach(func(v NullableCreatedByClient, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableCreatedByClientStream) Select(fn func(NullableCreatedByClient) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableCreatedByClientStream) Set(index int, val NullableCreatedByClient) *NullableCreatedByClientStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableCreatedByClientStream) Skip(skip int) *NullableCreatedByClientStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableCreatedByClientStream) SkippingEach(fn func(NullableCreatedByClient, int) int) *NullableCreatedByClientStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableCreatedByClientStream) Slice(startIndex, n int) *NullableCreatedByClientStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableCreatedByClient{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableCreatedByClientStream) Sort(fn func(i, j int) bool) *NullableCreatedByClientStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableCreatedByClientStream) Tail() *NullableCreatedByClient {
	return self.Last()
}
func (self *NullableCreatedByClientStream) TailOr(arg NullableCreatedByClient) NullableCreatedByClient {
	return self.LastOr(arg)
}
func (self *NullableCreatedByClientStream) ToList() []NullableCreatedByClient {
	return self.Val()
}
func (self *NullableCreatedByClientStream) Unique() *NullableCreatedByClientStream {
	return self.Distinct()
}
func (self *NullableCreatedByClientStream) Val() []NullableCreatedByClient {
	if self == nil {
		return []NullableCreatedByClient{}
	}
	return *self.Copy()
}
func (self *NullableCreatedByClientStream) While(fn func(NullableCreatedByClient, int) bool) *NullableCreatedByClientStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableCreatedByClientStream) Where(fn func(NullableCreatedByClient) bool) *NullableCreatedByClientStream {
	result := NullableCreatedByClientStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableCreatedByClientStream) WhereSlim(fn func(NullableCreatedByClient) bool) *NullableCreatedByClientStream {
	result := NullableCreatedByClientStreamOf()
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
