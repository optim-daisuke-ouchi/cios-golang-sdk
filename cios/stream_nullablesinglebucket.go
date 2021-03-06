/*
 * Collection utility of NullableSingleBucket Struct
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

type NullableSingleBucketStream []NullableSingleBucket

func NullableSingleBucketStreamOf(arg ...NullableSingleBucket) NullableSingleBucketStream {
	return arg
}
func NullableSingleBucketStreamFrom(arg []NullableSingleBucket) NullableSingleBucketStream {
	return arg
}
func CreateNullableSingleBucketStream(arg ...NullableSingleBucket) *NullableSingleBucketStream {
	tmp := NullableSingleBucketStreamOf(arg...)
	return &tmp
}
func GenerateNullableSingleBucketStream(arg []NullableSingleBucket) *NullableSingleBucketStream {
	tmp := NullableSingleBucketStreamFrom(arg)
	return &tmp
}

func (self *NullableSingleBucketStream) Add(arg NullableSingleBucket) *NullableSingleBucketStream {
	return self.AddAll(arg)
}
func (self *NullableSingleBucketStream) AddAll(arg ...NullableSingleBucket) *NullableSingleBucketStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSingleBucketStream) AddSafe(arg *NullableSingleBucket) *NullableSingleBucketStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSingleBucketStream) Aggregate(fn func(NullableSingleBucket, NullableSingleBucket) NullableSingleBucket) *NullableSingleBucketStream {
	result := NullableSingleBucketStreamOf()
	self.ForEach(func(v NullableSingleBucket, i int) {
		if i == 0 {
			result.Add(fn(NullableSingleBucket{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleBucketStream) AllMatch(fn func(NullableSingleBucket, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSingleBucketStream) AnyMatch(fn func(NullableSingleBucket, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSingleBucketStream) Clone() *NullableSingleBucketStream {
	temp := make([]NullableSingleBucket, self.Len())
	copy(temp, *self)
	return (*NullableSingleBucketStream)(&temp)
}
func (self *NullableSingleBucketStream) Copy() *NullableSingleBucketStream {
	return self.Clone()
}
func (self *NullableSingleBucketStream) Concat(arg []NullableSingleBucket) *NullableSingleBucketStream {
	return self.AddAll(arg...)
}
func (self *NullableSingleBucketStream) Contains(arg NullableSingleBucket) bool {
	return self.FindIndex(func(_arg NullableSingleBucket, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSingleBucketStream) Clean() *NullableSingleBucketStream {
	*self = NullableSingleBucketStreamOf()
	return self
}
func (self *NullableSingleBucketStream) Delete(index int) *NullableSingleBucketStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSingleBucketStream) DeleteRange(startIndex, endIndex int) *NullableSingleBucketStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSingleBucketStream) Distinct() *NullableSingleBucketStream {
	caches := map[string]bool{}
	result := NullableSingleBucketStreamOf()
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
func (self *NullableSingleBucketStream) Each(fn func(NullableSingleBucket)) *NullableSingleBucketStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSingleBucketStream) EachRight(fn func(NullableSingleBucket)) *NullableSingleBucketStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSingleBucketStream) Equals(arr []NullableSingleBucket) bool {
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
func (self *NullableSingleBucketStream) Filter(fn func(NullableSingleBucket, int) bool) *NullableSingleBucketStream {
	result := NullableSingleBucketStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleBucketStream) FilterSlim(fn func(NullableSingleBucket, int) bool) *NullableSingleBucketStream {
	result := NullableSingleBucketStreamOf()
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
func (self *NullableSingleBucketStream) Find(fn func(NullableSingleBucket, int) bool) *NullableSingleBucket {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSingleBucketStream) FindOr(fn func(NullableSingleBucket, int) bool, or NullableSingleBucket) NullableSingleBucket {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSingleBucketStream) FindIndex(fn func(NullableSingleBucket, int) bool) int {
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
func (self *NullableSingleBucketStream) First() *NullableSingleBucket {
	return self.Get(0)
}
func (self *NullableSingleBucketStream) FirstOr(arg NullableSingleBucket) NullableSingleBucket {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleBucketStream) ForEach(fn func(NullableSingleBucket, int)) *NullableSingleBucketStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSingleBucketStream) ForEachRight(fn func(NullableSingleBucket, int)) *NullableSingleBucketStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSingleBucketStream) GroupBy(fn func(NullableSingleBucket, int) string) map[string][]NullableSingleBucket {
	m := map[string][]NullableSingleBucket{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSingleBucketStream) GroupByValues(fn func(NullableSingleBucket, int) string) [][]NullableSingleBucket {
	var tmp [][]NullableSingleBucket
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSingleBucketStream) IndexOf(arg NullableSingleBucket) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSingleBucketStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSingleBucketStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSingleBucketStream) Last() *NullableSingleBucket {
	return self.Get(self.Len() - 1)
}
func (self *NullableSingleBucketStream) LastOr(arg NullableSingleBucket) NullableSingleBucket {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleBucketStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSingleBucketStream) Limit(limit int) *NullableSingleBucketStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSingleBucketStream) Map(fn func(NullableSingleBucket, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleBucketStream) Map2Int(fn func(NullableSingleBucket, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleBucketStream) Map2Int32(fn func(NullableSingleBucket, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleBucketStream) Map2Int64(fn func(NullableSingleBucket, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleBucketStream) Map2Float32(fn func(NullableSingleBucket, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleBucketStream) Map2Float64(fn func(NullableSingleBucket, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleBucketStream) Map2Bool(fn func(NullableSingleBucket, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleBucketStream) Map2Bytes(fn func(NullableSingleBucket, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleBucketStream) Map2String(fn func(NullableSingleBucket, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleBucketStream) Max(fn func(NullableSingleBucket, int) float64) *NullableSingleBucket {
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
func (self *NullableSingleBucketStream) Min(fn func(NullableSingleBucket, int) float64) *NullableSingleBucket {
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
func (self *NullableSingleBucketStream) NoneMatch(fn func(NullableSingleBucket, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSingleBucketStream) Get(index int) *NullableSingleBucket {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSingleBucketStream) GetOr(index int, arg NullableSingleBucket) NullableSingleBucket {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleBucketStream) Peek(fn func(*NullableSingleBucket, int)) *NullableSingleBucketStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSingleBucketStream) Reduce(fn func(NullableSingleBucket, NullableSingleBucket, int) NullableSingleBucket) *NullableSingleBucketStream {
	return self.ReduceInit(fn, NullableSingleBucket{})
}
func (self *NullableSingleBucketStream) ReduceInit(fn func(NullableSingleBucket, NullableSingleBucket, int) NullableSingleBucket, initialValue NullableSingleBucket) *NullableSingleBucketStream {
	result := NullableSingleBucketStreamOf()
	self.ForEach(func(v NullableSingleBucket, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleBucketStream) ReduceInterface(fn func(interface{}, NullableSingleBucket, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSingleBucket{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSingleBucketStream) ReduceString(fn func(string, NullableSingleBucket, int) string) []string {
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
func (self *NullableSingleBucketStream) ReduceInt(fn func(int, NullableSingleBucket, int) int) []int {
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
func (self *NullableSingleBucketStream) ReduceInt32(fn func(int32, NullableSingleBucket, int) int32) []int32 {
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
func (self *NullableSingleBucketStream) ReduceInt64(fn func(int64, NullableSingleBucket, int) int64) []int64 {
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
func (self *NullableSingleBucketStream) ReduceFloat32(fn func(float32, NullableSingleBucket, int) float32) []float32 {
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
func (self *NullableSingleBucketStream) ReduceFloat64(fn func(float64, NullableSingleBucket, int) float64) []float64 {
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
func (self *NullableSingleBucketStream) ReduceBool(fn func(bool, NullableSingleBucket, int) bool) []bool {
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
func (self *NullableSingleBucketStream) Reverse() *NullableSingleBucketStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSingleBucketStream) Replace(fn func(NullableSingleBucket, int) NullableSingleBucket) *NullableSingleBucketStream {
	return self.ForEach(func(v NullableSingleBucket, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSingleBucketStream) Select(fn func(NullableSingleBucket) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSingleBucketStream) Set(index int, val NullableSingleBucket) *NullableSingleBucketStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSingleBucketStream) Skip(skip int) *NullableSingleBucketStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSingleBucketStream) SkippingEach(fn func(NullableSingleBucket, int) int) *NullableSingleBucketStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSingleBucketStream) Slice(startIndex, n int) *NullableSingleBucketStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSingleBucket{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSingleBucketStream) Sort(fn func(i, j int) bool) *NullableSingleBucketStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSingleBucketStream) Tail() *NullableSingleBucket {
	return self.Last()
}
func (self *NullableSingleBucketStream) TailOr(arg NullableSingleBucket) NullableSingleBucket {
	return self.LastOr(arg)
}
func (self *NullableSingleBucketStream) ToList() []NullableSingleBucket {
	return self.Val()
}
func (self *NullableSingleBucketStream) Unique() *NullableSingleBucketStream {
	return self.Distinct()
}
func (self *NullableSingleBucketStream) Val() []NullableSingleBucket {
	if self == nil {
		return []NullableSingleBucket{}
	}
	return *self.Copy()
}
func (self *NullableSingleBucketStream) While(fn func(NullableSingleBucket, int) bool) *NullableSingleBucketStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSingleBucketStream) Where(fn func(NullableSingleBucket) bool) *NullableSingleBucketStream {
	result := NullableSingleBucketStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleBucketStream) WhereSlim(fn func(NullableSingleBucket) bool) *NullableSingleBucketStream {
	result := NullableSingleBucketStreamOf()
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
