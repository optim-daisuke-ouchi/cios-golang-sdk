/*
 * Collection utility of ApiGetBucketsRequest Struct
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

type ApiGetBucketsRequestStream []ApiGetBucketsRequest

func ApiGetBucketsRequestStreamOf(arg ...ApiGetBucketsRequest) ApiGetBucketsRequestStream {
	return arg
}
func ApiGetBucketsRequestStreamFrom(arg []ApiGetBucketsRequest) ApiGetBucketsRequestStream {
	return arg
}
func CreateApiGetBucketsRequestStream(arg ...ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	tmp := ApiGetBucketsRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetBucketsRequestStream(arg []ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	tmp := ApiGetBucketsRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetBucketsRequestStream) Add(arg ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetBucketsRequestStream) AddAll(arg ...ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetBucketsRequestStream) AddSafe(arg *ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetBucketsRequestStream) Aggregate(fn func(ApiGetBucketsRequest, ApiGetBucketsRequest) ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	result := ApiGetBucketsRequestStreamOf()
	self.ForEach(func(v ApiGetBucketsRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetBucketsRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetBucketsRequestStream) AllMatch(fn func(ApiGetBucketsRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetBucketsRequestStream) AnyMatch(fn func(ApiGetBucketsRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetBucketsRequestStream) Clone() *ApiGetBucketsRequestStream {
	temp := make([]ApiGetBucketsRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetBucketsRequestStream)(&temp)
}
func (self *ApiGetBucketsRequestStream) Copy() *ApiGetBucketsRequestStream {
	return self.Clone()
}
func (self *ApiGetBucketsRequestStream) Concat(arg []ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetBucketsRequestStream) Contains(arg ApiGetBucketsRequest) bool {
	return self.FindIndex(func(_arg ApiGetBucketsRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetBucketsRequestStream) Clean() *ApiGetBucketsRequestStream {
	*self = ApiGetBucketsRequestStreamOf()
	return self
}
func (self *ApiGetBucketsRequestStream) Delete(index int) *ApiGetBucketsRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetBucketsRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetBucketsRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetBucketsRequestStream) Distinct() *ApiGetBucketsRequestStream {
	caches := map[string]bool{}
	result := ApiGetBucketsRequestStreamOf()
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
func (self *ApiGetBucketsRequestStream) Each(fn func(ApiGetBucketsRequest)) *ApiGetBucketsRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetBucketsRequestStream) EachRight(fn func(ApiGetBucketsRequest)) *ApiGetBucketsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetBucketsRequestStream) Equals(arr []ApiGetBucketsRequest) bool {
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
func (self *ApiGetBucketsRequestStream) Filter(fn func(ApiGetBucketsRequest, int) bool) *ApiGetBucketsRequestStream {
	result := ApiGetBucketsRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetBucketsRequestStream) FilterSlim(fn func(ApiGetBucketsRequest, int) bool) *ApiGetBucketsRequestStream {
	result := ApiGetBucketsRequestStreamOf()
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
func (self *ApiGetBucketsRequestStream) Find(fn func(ApiGetBucketsRequest, int) bool) *ApiGetBucketsRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetBucketsRequestStream) FindOr(fn func(ApiGetBucketsRequest, int) bool, or ApiGetBucketsRequest) ApiGetBucketsRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetBucketsRequestStream) FindIndex(fn func(ApiGetBucketsRequest, int) bool) int {
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
func (self *ApiGetBucketsRequestStream) First() *ApiGetBucketsRequest {
	return self.Get(0)
}
func (self *ApiGetBucketsRequestStream) FirstOr(arg ApiGetBucketsRequest) ApiGetBucketsRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetBucketsRequestStream) ForEach(fn func(ApiGetBucketsRequest, int)) *ApiGetBucketsRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetBucketsRequestStream) ForEachRight(fn func(ApiGetBucketsRequest, int)) *ApiGetBucketsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetBucketsRequestStream) GroupBy(fn func(ApiGetBucketsRequest, int) string) map[string][]ApiGetBucketsRequest {
	m := map[string][]ApiGetBucketsRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetBucketsRequestStream) GroupByValues(fn func(ApiGetBucketsRequest, int) string) [][]ApiGetBucketsRequest {
	var tmp [][]ApiGetBucketsRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetBucketsRequestStream) IndexOf(arg ApiGetBucketsRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetBucketsRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetBucketsRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetBucketsRequestStream) Last() *ApiGetBucketsRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetBucketsRequestStream) LastOr(arg ApiGetBucketsRequest) ApiGetBucketsRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetBucketsRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetBucketsRequestStream) Limit(limit int) *ApiGetBucketsRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetBucketsRequestStream) Map(fn func(ApiGetBucketsRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Map2Int(fn func(ApiGetBucketsRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Map2Int32(fn func(ApiGetBucketsRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Map2Int64(fn func(ApiGetBucketsRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Map2Float32(fn func(ApiGetBucketsRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Map2Float64(fn func(ApiGetBucketsRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Map2Bool(fn func(ApiGetBucketsRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Map2Bytes(fn func(ApiGetBucketsRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Map2String(fn func(ApiGetBucketsRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Max(fn func(ApiGetBucketsRequest, int) float64) *ApiGetBucketsRequest {
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
func (self *ApiGetBucketsRequestStream) Min(fn func(ApiGetBucketsRequest, int) float64) *ApiGetBucketsRequest {
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
func (self *ApiGetBucketsRequestStream) NoneMatch(fn func(ApiGetBucketsRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetBucketsRequestStream) Get(index int) *ApiGetBucketsRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetBucketsRequestStream) GetOr(index int, arg ApiGetBucketsRequest) ApiGetBucketsRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetBucketsRequestStream) Peek(fn func(*ApiGetBucketsRequest, int)) *ApiGetBucketsRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetBucketsRequestStream) Reduce(fn func(ApiGetBucketsRequest, ApiGetBucketsRequest, int) ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	return self.ReduceInit(fn, ApiGetBucketsRequest{})
}
func (self *ApiGetBucketsRequestStream) ReduceInit(fn func(ApiGetBucketsRequest, ApiGetBucketsRequest, int) ApiGetBucketsRequest, initialValue ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	result := ApiGetBucketsRequestStreamOf()
	self.ForEach(func(v ApiGetBucketsRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetBucketsRequestStream) ReduceInterface(fn func(interface{}, ApiGetBucketsRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetBucketsRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetBucketsRequestStream) ReduceString(fn func(string, ApiGetBucketsRequest, int) string) []string {
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
func (self *ApiGetBucketsRequestStream) ReduceInt(fn func(int, ApiGetBucketsRequest, int) int) []int {
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
func (self *ApiGetBucketsRequestStream) ReduceInt32(fn func(int32, ApiGetBucketsRequest, int) int32) []int32 {
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
func (self *ApiGetBucketsRequestStream) ReduceInt64(fn func(int64, ApiGetBucketsRequest, int) int64) []int64 {
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
func (self *ApiGetBucketsRequestStream) ReduceFloat32(fn func(float32, ApiGetBucketsRequest, int) float32) []float32 {
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
func (self *ApiGetBucketsRequestStream) ReduceFloat64(fn func(float64, ApiGetBucketsRequest, int) float64) []float64 {
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
func (self *ApiGetBucketsRequestStream) ReduceBool(fn func(bool, ApiGetBucketsRequest, int) bool) []bool {
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
func (self *ApiGetBucketsRequestStream) Reverse() *ApiGetBucketsRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetBucketsRequestStream) Replace(fn func(ApiGetBucketsRequest, int) ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	return self.ForEach(func(v ApiGetBucketsRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetBucketsRequestStream) Select(fn func(ApiGetBucketsRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetBucketsRequestStream) Set(index int, val ApiGetBucketsRequest) *ApiGetBucketsRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetBucketsRequestStream) Skip(skip int) *ApiGetBucketsRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetBucketsRequestStream) SkippingEach(fn func(ApiGetBucketsRequest, int) int) *ApiGetBucketsRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetBucketsRequestStream) Slice(startIndex, n int) *ApiGetBucketsRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetBucketsRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetBucketsRequestStream) Sort(fn func(i, j int) bool) *ApiGetBucketsRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetBucketsRequestStream) Tail() *ApiGetBucketsRequest {
	return self.Last()
}
func (self *ApiGetBucketsRequestStream) TailOr(arg ApiGetBucketsRequest) ApiGetBucketsRequest {
	return self.LastOr(arg)
}
func (self *ApiGetBucketsRequestStream) ToList() []ApiGetBucketsRequest {
	return self.Val()
}
func (self *ApiGetBucketsRequestStream) Unique() *ApiGetBucketsRequestStream {
	return self.Distinct()
}
func (self *ApiGetBucketsRequestStream) Val() []ApiGetBucketsRequest {
	if self == nil {
		return []ApiGetBucketsRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetBucketsRequestStream) While(fn func(ApiGetBucketsRequest, int) bool) *ApiGetBucketsRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetBucketsRequestStream) Where(fn func(ApiGetBucketsRequest) bool) *ApiGetBucketsRequestStream {
	result := ApiGetBucketsRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetBucketsRequestStream) WhereSlim(fn func(ApiGetBucketsRequest) bool) *ApiGetBucketsRequestStream {
	result := ApiGetBucketsRequestStreamOf()
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
