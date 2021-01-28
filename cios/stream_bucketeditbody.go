/*
 * Collection utility of BucketEditBody Struct
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

type BucketEditBodyStream []BucketEditBody

func BucketEditBodyStreamOf(arg ...BucketEditBody) BucketEditBodyStream {
	return arg
}
func BucketEditBodyStreamFrom(arg []BucketEditBody) BucketEditBodyStream {
	return arg
}
func CreateBucketEditBodyStream(arg ...BucketEditBody) *BucketEditBodyStream {
	tmp := BucketEditBodyStreamOf(arg...)
	return &tmp
}
func GenerateBucketEditBodyStream(arg []BucketEditBody) *BucketEditBodyStream {
	tmp := BucketEditBodyStreamFrom(arg)
	return &tmp
}

func (self *BucketEditBodyStream) Add(arg BucketEditBody) *BucketEditBodyStream {
	return self.AddAll(arg)
}
func (self *BucketEditBodyStream) AddAll(arg ...BucketEditBody) *BucketEditBodyStream {
	*self = append(*self, arg...)
	return self
}
func (self *BucketEditBodyStream) AddSafe(arg *BucketEditBody) *BucketEditBodyStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *BucketEditBodyStream) Aggregate(fn func(BucketEditBody, BucketEditBody) BucketEditBody) *BucketEditBodyStream {
	result := BucketEditBodyStreamOf()
	self.ForEach(func(v BucketEditBody, i int) {
		if i == 0 {
			result.Add(fn(BucketEditBody{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *BucketEditBodyStream) AllMatch(fn func(BucketEditBody, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *BucketEditBodyStream) AnyMatch(fn func(BucketEditBody, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *BucketEditBodyStream) Clone() *BucketEditBodyStream {
	temp := make([]BucketEditBody, self.Len())
	copy(temp, *self)
	return (*BucketEditBodyStream)(&temp)
}
func (self *BucketEditBodyStream) Copy() *BucketEditBodyStream {
	return self.Clone()
}
func (self *BucketEditBodyStream) Concat(arg []BucketEditBody) *BucketEditBodyStream {
	return self.AddAll(arg...)
}
func (self *BucketEditBodyStream) Contains(arg BucketEditBody) bool {
	return self.FindIndex(func(_arg BucketEditBody, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *BucketEditBodyStream) Clean() *BucketEditBodyStream {
	*self = BucketEditBodyStreamOf()
	return self
}
func (self *BucketEditBodyStream) Delete(index int) *BucketEditBodyStream {
	return self.DeleteRange(index, index)
}
func (self *BucketEditBodyStream) DeleteRange(startIndex, endIndex int) *BucketEditBodyStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *BucketEditBodyStream) Distinct() *BucketEditBodyStream {
	caches := map[string]bool{}
	result := BucketEditBodyStreamOf()
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
func (self *BucketEditBodyStream) Each(fn func(BucketEditBody)) *BucketEditBodyStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *BucketEditBodyStream) EachRight(fn func(BucketEditBody)) *BucketEditBodyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *BucketEditBodyStream) Equals(arr []BucketEditBody) bool {
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
func (self *BucketEditBodyStream) Filter(fn func(BucketEditBody, int) bool) *BucketEditBodyStream {
	result := BucketEditBodyStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *BucketEditBodyStream) FilterSlim(fn func(BucketEditBody, int) bool) *BucketEditBodyStream {
	result := BucketEditBodyStreamOf()
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
func (self *BucketEditBodyStream) Find(fn func(BucketEditBody, int) bool) *BucketEditBody {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *BucketEditBodyStream) FindOr(fn func(BucketEditBody, int) bool, or BucketEditBody) BucketEditBody {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *BucketEditBodyStream) FindIndex(fn func(BucketEditBody, int) bool) int {
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
func (self *BucketEditBodyStream) First() *BucketEditBody {
	return self.Get(0)
}
func (self *BucketEditBodyStream) FirstOr(arg BucketEditBody) BucketEditBody {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *BucketEditBodyStream) ForEach(fn func(BucketEditBody, int)) *BucketEditBodyStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *BucketEditBodyStream) ForEachRight(fn func(BucketEditBody, int)) *BucketEditBodyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *BucketEditBodyStream) GroupBy(fn func(BucketEditBody, int) string) map[string][]BucketEditBody {
	m := map[string][]BucketEditBody{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *BucketEditBodyStream) GroupByValues(fn func(BucketEditBody, int) string) [][]BucketEditBody {
	var tmp [][]BucketEditBody
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *BucketEditBodyStream) IndexOf(arg BucketEditBody) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *BucketEditBodyStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *BucketEditBodyStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *BucketEditBodyStream) Last() *BucketEditBody {
	return self.Get(self.Len() - 1)
}
func (self *BucketEditBodyStream) LastOr(arg BucketEditBody) BucketEditBody {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *BucketEditBodyStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *BucketEditBodyStream) Limit(limit int) *BucketEditBodyStream {
	self.Slice(0, limit)
	return self
}

func (self *BucketEditBodyStream) Map(fn func(BucketEditBody, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketEditBodyStream) Map2Int(fn func(BucketEditBody, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketEditBodyStream) Map2Int32(fn func(BucketEditBody, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketEditBodyStream) Map2Int64(fn func(BucketEditBody, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketEditBodyStream) Map2Float32(fn func(BucketEditBody, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketEditBodyStream) Map2Float64(fn func(BucketEditBody, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketEditBodyStream) Map2Bool(fn func(BucketEditBody, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketEditBodyStream) Map2Bytes(fn func(BucketEditBody, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketEditBodyStream) Map2String(fn func(BucketEditBody, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketEditBodyStream) Max(fn func(BucketEditBody, int) float64) *BucketEditBody {
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
func (self *BucketEditBodyStream) Min(fn func(BucketEditBody, int) float64) *BucketEditBody {
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
func (self *BucketEditBodyStream) NoneMatch(fn func(BucketEditBody, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *BucketEditBodyStream) Get(index int) *BucketEditBody {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *BucketEditBodyStream) GetOr(index int, arg BucketEditBody) BucketEditBody {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *BucketEditBodyStream) Peek(fn func(*BucketEditBody, int)) *BucketEditBodyStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *BucketEditBodyStream) Reduce(fn func(BucketEditBody, BucketEditBody, int) BucketEditBody) *BucketEditBodyStream {
	return self.ReduceInit(fn, BucketEditBody{})
}
func (self *BucketEditBodyStream) ReduceInit(fn func(BucketEditBody, BucketEditBody, int) BucketEditBody, initialValue BucketEditBody) *BucketEditBodyStream {
	result := BucketEditBodyStreamOf()
	self.ForEach(func(v BucketEditBody, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *BucketEditBodyStream) ReduceInterface(fn func(interface{}, BucketEditBody, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(BucketEditBody{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *BucketEditBodyStream) ReduceString(fn func(string, BucketEditBody, int) string) []string {
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
func (self *BucketEditBodyStream) ReduceInt(fn func(int, BucketEditBody, int) int) []int {
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
func (self *BucketEditBodyStream) ReduceInt32(fn func(int32, BucketEditBody, int) int32) []int32 {
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
func (self *BucketEditBodyStream) ReduceInt64(fn func(int64, BucketEditBody, int) int64) []int64 {
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
func (self *BucketEditBodyStream) ReduceFloat32(fn func(float32, BucketEditBody, int) float32) []float32 {
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
func (self *BucketEditBodyStream) ReduceFloat64(fn func(float64, BucketEditBody, int) float64) []float64 {
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
func (self *BucketEditBodyStream) ReduceBool(fn func(bool, BucketEditBody, int) bool) []bool {
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
func (self *BucketEditBodyStream) Reverse() *BucketEditBodyStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *BucketEditBodyStream) Replace(fn func(BucketEditBody, int) BucketEditBody) *BucketEditBodyStream {
	return self.ForEach(func(v BucketEditBody, i int) { self.Set(i, fn(v, i)) })
}
func (self *BucketEditBodyStream) Select(fn func(BucketEditBody) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *BucketEditBodyStream) Set(index int, val BucketEditBody) *BucketEditBodyStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *BucketEditBodyStream) Skip(skip int) *BucketEditBodyStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *BucketEditBodyStream) SkippingEach(fn func(BucketEditBody, int) int) *BucketEditBodyStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *BucketEditBodyStream) Slice(startIndex, n int) *BucketEditBodyStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []BucketEditBody{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *BucketEditBodyStream) Sort(fn func(i, j int) bool) *BucketEditBodyStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *BucketEditBodyStream) Tail() *BucketEditBody {
	return self.Last()
}
func (self *BucketEditBodyStream) TailOr(arg BucketEditBody) BucketEditBody {
	return self.LastOr(arg)
}
func (self *BucketEditBodyStream) ToList() []BucketEditBody {
	return self.Val()
}
func (self *BucketEditBodyStream) Unique() *BucketEditBodyStream {
	return self.Distinct()
}
func (self *BucketEditBodyStream) Val() []BucketEditBody {
	if self == nil {
		return []BucketEditBody{}
	}
	return *self.Copy()
}
func (self *BucketEditBodyStream) While(fn func(BucketEditBody, int) bool) *BucketEditBodyStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *BucketEditBodyStream) Where(fn func(BucketEditBody) bool) *BucketEditBodyStream {
	result := BucketEditBodyStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *BucketEditBodyStream) WhereSlim(fn func(BucketEditBody) bool) *BucketEditBodyStream {
	result := BucketEditBodyStreamOf()
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
