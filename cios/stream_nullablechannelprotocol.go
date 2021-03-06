/*
 * Collection utility of NullableChannelProtocol Struct
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

type NullableChannelProtocolStream []NullableChannelProtocol

func NullableChannelProtocolStreamOf(arg ...NullableChannelProtocol) NullableChannelProtocolStream {
	return arg
}
func NullableChannelProtocolStreamFrom(arg []NullableChannelProtocol) NullableChannelProtocolStream {
	return arg
}
func CreateNullableChannelProtocolStream(arg ...NullableChannelProtocol) *NullableChannelProtocolStream {
	tmp := NullableChannelProtocolStreamOf(arg...)
	return &tmp
}
func GenerateNullableChannelProtocolStream(arg []NullableChannelProtocol) *NullableChannelProtocolStream {
	tmp := NullableChannelProtocolStreamFrom(arg)
	return &tmp
}

func (self *NullableChannelProtocolStream) Add(arg NullableChannelProtocol) *NullableChannelProtocolStream {
	return self.AddAll(arg)
}
func (self *NullableChannelProtocolStream) AddAll(arg ...NullableChannelProtocol) *NullableChannelProtocolStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableChannelProtocolStream) AddSafe(arg *NullableChannelProtocol) *NullableChannelProtocolStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableChannelProtocolStream) Aggregate(fn func(NullableChannelProtocol, NullableChannelProtocol) NullableChannelProtocol) *NullableChannelProtocolStream {
	result := NullableChannelProtocolStreamOf()
	self.ForEach(func(v NullableChannelProtocol, i int) {
		if i == 0 {
			result.Add(fn(NullableChannelProtocol{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableChannelProtocolStream) AllMatch(fn func(NullableChannelProtocol, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableChannelProtocolStream) AnyMatch(fn func(NullableChannelProtocol, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableChannelProtocolStream) Clone() *NullableChannelProtocolStream {
	temp := make([]NullableChannelProtocol, self.Len())
	copy(temp, *self)
	return (*NullableChannelProtocolStream)(&temp)
}
func (self *NullableChannelProtocolStream) Copy() *NullableChannelProtocolStream {
	return self.Clone()
}
func (self *NullableChannelProtocolStream) Concat(arg []NullableChannelProtocol) *NullableChannelProtocolStream {
	return self.AddAll(arg...)
}
func (self *NullableChannelProtocolStream) Contains(arg NullableChannelProtocol) bool {
	return self.FindIndex(func(_arg NullableChannelProtocol, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableChannelProtocolStream) Clean() *NullableChannelProtocolStream {
	*self = NullableChannelProtocolStreamOf()
	return self
}
func (self *NullableChannelProtocolStream) Delete(index int) *NullableChannelProtocolStream {
	return self.DeleteRange(index, index)
}
func (self *NullableChannelProtocolStream) DeleteRange(startIndex, endIndex int) *NullableChannelProtocolStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableChannelProtocolStream) Distinct() *NullableChannelProtocolStream {
	caches := map[string]bool{}
	result := NullableChannelProtocolStreamOf()
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
func (self *NullableChannelProtocolStream) Each(fn func(NullableChannelProtocol)) *NullableChannelProtocolStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableChannelProtocolStream) EachRight(fn func(NullableChannelProtocol)) *NullableChannelProtocolStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableChannelProtocolStream) Equals(arr []NullableChannelProtocol) bool {
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
func (self *NullableChannelProtocolStream) Filter(fn func(NullableChannelProtocol, int) bool) *NullableChannelProtocolStream {
	result := NullableChannelProtocolStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableChannelProtocolStream) FilterSlim(fn func(NullableChannelProtocol, int) bool) *NullableChannelProtocolStream {
	result := NullableChannelProtocolStreamOf()
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
func (self *NullableChannelProtocolStream) Find(fn func(NullableChannelProtocol, int) bool) *NullableChannelProtocol {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableChannelProtocolStream) FindOr(fn func(NullableChannelProtocol, int) bool, or NullableChannelProtocol) NullableChannelProtocol {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableChannelProtocolStream) FindIndex(fn func(NullableChannelProtocol, int) bool) int {
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
func (self *NullableChannelProtocolStream) First() *NullableChannelProtocol {
	return self.Get(0)
}
func (self *NullableChannelProtocolStream) FirstOr(arg NullableChannelProtocol) NullableChannelProtocol {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableChannelProtocolStream) ForEach(fn func(NullableChannelProtocol, int)) *NullableChannelProtocolStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableChannelProtocolStream) ForEachRight(fn func(NullableChannelProtocol, int)) *NullableChannelProtocolStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableChannelProtocolStream) GroupBy(fn func(NullableChannelProtocol, int) string) map[string][]NullableChannelProtocol {
	m := map[string][]NullableChannelProtocol{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableChannelProtocolStream) GroupByValues(fn func(NullableChannelProtocol, int) string) [][]NullableChannelProtocol {
	var tmp [][]NullableChannelProtocol
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableChannelProtocolStream) IndexOf(arg NullableChannelProtocol) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableChannelProtocolStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableChannelProtocolStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableChannelProtocolStream) Last() *NullableChannelProtocol {
	return self.Get(self.Len() - 1)
}
func (self *NullableChannelProtocolStream) LastOr(arg NullableChannelProtocol) NullableChannelProtocol {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableChannelProtocolStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableChannelProtocolStream) Limit(limit int) *NullableChannelProtocolStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableChannelProtocolStream) Map(fn func(NullableChannelProtocol, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Map2Int(fn func(NullableChannelProtocol, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Map2Int32(fn func(NullableChannelProtocol, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Map2Int64(fn func(NullableChannelProtocol, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Map2Float32(fn func(NullableChannelProtocol, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Map2Float64(fn func(NullableChannelProtocol, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Map2Bool(fn func(NullableChannelProtocol, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Map2Bytes(fn func(NullableChannelProtocol, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Map2String(fn func(NullableChannelProtocol, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Max(fn func(NullableChannelProtocol, int) float64) *NullableChannelProtocol {
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
func (self *NullableChannelProtocolStream) Min(fn func(NullableChannelProtocol, int) float64) *NullableChannelProtocol {
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
func (self *NullableChannelProtocolStream) NoneMatch(fn func(NullableChannelProtocol, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableChannelProtocolStream) Get(index int) *NullableChannelProtocol {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableChannelProtocolStream) GetOr(index int, arg NullableChannelProtocol) NullableChannelProtocol {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableChannelProtocolStream) Peek(fn func(*NullableChannelProtocol, int)) *NullableChannelProtocolStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableChannelProtocolStream) Reduce(fn func(NullableChannelProtocol, NullableChannelProtocol, int) NullableChannelProtocol) *NullableChannelProtocolStream {
	return self.ReduceInit(fn, NullableChannelProtocol{})
}
func (self *NullableChannelProtocolStream) ReduceInit(fn func(NullableChannelProtocol, NullableChannelProtocol, int) NullableChannelProtocol, initialValue NullableChannelProtocol) *NullableChannelProtocolStream {
	result := NullableChannelProtocolStreamOf()
	self.ForEach(func(v NullableChannelProtocol, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableChannelProtocolStream) ReduceInterface(fn func(interface{}, NullableChannelProtocol, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableChannelProtocol{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableChannelProtocolStream) ReduceString(fn func(string, NullableChannelProtocol, int) string) []string {
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
func (self *NullableChannelProtocolStream) ReduceInt(fn func(int, NullableChannelProtocol, int) int) []int {
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
func (self *NullableChannelProtocolStream) ReduceInt32(fn func(int32, NullableChannelProtocol, int) int32) []int32 {
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
func (self *NullableChannelProtocolStream) ReduceInt64(fn func(int64, NullableChannelProtocol, int) int64) []int64 {
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
func (self *NullableChannelProtocolStream) ReduceFloat32(fn func(float32, NullableChannelProtocol, int) float32) []float32 {
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
func (self *NullableChannelProtocolStream) ReduceFloat64(fn func(float64, NullableChannelProtocol, int) float64) []float64 {
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
func (self *NullableChannelProtocolStream) ReduceBool(fn func(bool, NullableChannelProtocol, int) bool) []bool {
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
func (self *NullableChannelProtocolStream) Reverse() *NullableChannelProtocolStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableChannelProtocolStream) Replace(fn func(NullableChannelProtocol, int) NullableChannelProtocol) *NullableChannelProtocolStream {
	return self.ForEach(func(v NullableChannelProtocol, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableChannelProtocolStream) Select(fn func(NullableChannelProtocol) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableChannelProtocolStream) Set(index int, val NullableChannelProtocol) *NullableChannelProtocolStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableChannelProtocolStream) Skip(skip int) *NullableChannelProtocolStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableChannelProtocolStream) SkippingEach(fn func(NullableChannelProtocol, int) int) *NullableChannelProtocolStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableChannelProtocolStream) Slice(startIndex, n int) *NullableChannelProtocolStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableChannelProtocol{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableChannelProtocolStream) Sort(fn func(i, j int) bool) *NullableChannelProtocolStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableChannelProtocolStream) Tail() *NullableChannelProtocol {
	return self.Last()
}
func (self *NullableChannelProtocolStream) TailOr(arg NullableChannelProtocol) NullableChannelProtocol {
	return self.LastOr(arg)
}
func (self *NullableChannelProtocolStream) ToList() []NullableChannelProtocol {
	return self.Val()
}
func (self *NullableChannelProtocolStream) Unique() *NullableChannelProtocolStream {
	return self.Distinct()
}
func (self *NullableChannelProtocolStream) Val() []NullableChannelProtocol {
	if self == nil {
		return []NullableChannelProtocol{}
	}
	return *self.Copy()
}
func (self *NullableChannelProtocolStream) While(fn func(NullableChannelProtocol, int) bool) *NullableChannelProtocolStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableChannelProtocolStream) Where(fn func(NullableChannelProtocol) bool) *NullableChannelProtocolStream {
	result := NullableChannelProtocolStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableChannelProtocolStream) WhereSlim(fn func(NullableChannelProtocol) bool) *NullableChannelProtocolStream {
	result := NullableChannelProtocolStreamOf()
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
