/*
 * Collection utility of ApiGetTimeSeriesDataRequest Struct
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

type ApiGetTimeSeriesDataRequestStream []ApiGetTimeSeriesDataRequest

func ApiGetTimeSeriesDataRequestStreamOf(arg ...ApiGetTimeSeriesDataRequest) ApiGetTimeSeriesDataRequestStream {
	return arg
}
func ApiGetTimeSeriesDataRequestStreamFrom(arg []ApiGetTimeSeriesDataRequest) ApiGetTimeSeriesDataRequestStream {
	return arg
}
func CreateApiGetTimeSeriesDataRequestStream(arg ...ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	tmp := ApiGetTimeSeriesDataRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetTimeSeriesDataRequestStream(arg []ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	tmp := ApiGetTimeSeriesDataRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetTimeSeriesDataRequestStream) Add(arg ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetTimeSeriesDataRequestStream) AddAll(arg ...ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) AddSafe(arg *ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) Aggregate(fn func(ApiGetTimeSeriesDataRequest, ApiGetTimeSeriesDataRequest) ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	result := ApiGetTimeSeriesDataRequestStreamOf()
	self.ForEach(func(v ApiGetTimeSeriesDataRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetTimeSeriesDataRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) AllMatch(fn func(ApiGetTimeSeriesDataRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetTimeSeriesDataRequestStream) AnyMatch(fn func(ApiGetTimeSeriesDataRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetTimeSeriesDataRequestStream) Clone() *ApiGetTimeSeriesDataRequestStream {
	temp := make([]ApiGetTimeSeriesDataRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetTimeSeriesDataRequestStream)(&temp)
}
func (self *ApiGetTimeSeriesDataRequestStream) Copy() *ApiGetTimeSeriesDataRequestStream {
	return self.Clone()
}
func (self *ApiGetTimeSeriesDataRequestStream) Concat(arg []ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetTimeSeriesDataRequestStream) Contains(arg ApiGetTimeSeriesDataRequest) bool {
	return self.FindIndex(func(_arg ApiGetTimeSeriesDataRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetTimeSeriesDataRequestStream) Clean() *ApiGetTimeSeriesDataRequestStream {
	*self = ApiGetTimeSeriesDataRequestStreamOf()
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) Delete(index int) *ApiGetTimeSeriesDataRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetTimeSeriesDataRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetTimeSeriesDataRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) Distinct() *ApiGetTimeSeriesDataRequestStream {
	caches := map[string]bool{}
	result := ApiGetTimeSeriesDataRequestStreamOf()
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
func (self *ApiGetTimeSeriesDataRequestStream) Each(fn func(ApiGetTimeSeriesDataRequest)) *ApiGetTimeSeriesDataRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) EachRight(fn func(ApiGetTimeSeriesDataRequest)) *ApiGetTimeSeriesDataRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) Equals(arr []ApiGetTimeSeriesDataRequest) bool {
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
func (self *ApiGetTimeSeriesDataRequestStream) Filter(fn func(ApiGetTimeSeriesDataRequest, int) bool) *ApiGetTimeSeriesDataRequestStream {
	result := ApiGetTimeSeriesDataRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) FilterSlim(fn func(ApiGetTimeSeriesDataRequest, int) bool) *ApiGetTimeSeriesDataRequestStream {
	result := ApiGetTimeSeriesDataRequestStreamOf()
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
func (self *ApiGetTimeSeriesDataRequestStream) Find(fn func(ApiGetTimeSeriesDataRequest, int) bool) *ApiGetTimeSeriesDataRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetTimeSeriesDataRequestStream) FindOr(fn func(ApiGetTimeSeriesDataRequest, int) bool, or ApiGetTimeSeriesDataRequest) ApiGetTimeSeriesDataRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetTimeSeriesDataRequestStream) FindIndex(fn func(ApiGetTimeSeriesDataRequest, int) bool) int {
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
func (self *ApiGetTimeSeriesDataRequestStream) First() *ApiGetTimeSeriesDataRequest {
	return self.Get(0)
}
func (self *ApiGetTimeSeriesDataRequestStream) FirstOr(arg ApiGetTimeSeriesDataRequest) ApiGetTimeSeriesDataRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetTimeSeriesDataRequestStream) ForEach(fn func(ApiGetTimeSeriesDataRequest, int)) *ApiGetTimeSeriesDataRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) ForEachRight(fn func(ApiGetTimeSeriesDataRequest, int)) *ApiGetTimeSeriesDataRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) GroupBy(fn func(ApiGetTimeSeriesDataRequest, int) string) map[string][]ApiGetTimeSeriesDataRequest {
	m := map[string][]ApiGetTimeSeriesDataRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetTimeSeriesDataRequestStream) GroupByValues(fn func(ApiGetTimeSeriesDataRequest, int) string) [][]ApiGetTimeSeriesDataRequest {
	var tmp [][]ApiGetTimeSeriesDataRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetTimeSeriesDataRequestStream) IndexOf(arg ApiGetTimeSeriesDataRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetTimeSeriesDataRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetTimeSeriesDataRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetTimeSeriesDataRequestStream) Last() *ApiGetTimeSeriesDataRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetTimeSeriesDataRequestStream) LastOr(arg ApiGetTimeSeriesDataRequest) ApiGetTimeSeriesDataRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetTimeSeriesDataRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetTimeSeriesDataRequestStream) Limit(limit int) *ApiGetTimeSeriesDataRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetTimeSeriesDataRequestStream) Map(fn func(ApiGetTimeSeriesDataRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Map2Int(fn func(ApiGetTimeSeriesDataRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Map2Int32(fn func(ApiGetTimeSeriesDataRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Map2Int64(fn func(ApiGetTimeSeriesDataRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Map2Float32(fn func(ApiGetTimeSeriesDataRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Map2Float64(fn func(ApiGetTimeSeriesDataRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Map2Bool(fn func(ApiGetTimeSeriesDataRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Map2Bytes(fn func(ApiGetTimeSeriesDataRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Map2String(fn func(ApiGetTimeSeriesDataRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Max(fn func(ApiGetTimeSeriesDataRequest, int) float64) *ApiGetTimeSeriesDataRequest {
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
func (self *ApiGetTimeSeriesDataRequestStream) Min(fn func(ApiGetTimeSeriesDataRequest, int) float64) *ApiGetTimeSeriesDataRequest {
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
func (self *ApiGetTimeSeriesDataRequestStream) NoneMatch(fn func(ApiGetTimeSeriesDataRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetTimeSeriesDataRequestStream) Get(index int) *ApiGetTimeSeriesDataRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetTimeSeriesDataRequestStream) GetOr(index int, arg ApiGetTimeSeriesDataRequest) ApiGetTimeSeriesDataRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetTimeSeriesDataRequestStream) Peek(fn func(*ApiGetTimeSeriesDataRequest, int)) *ApiGetTimeSeriesDataRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetTimeSeriesDataRequestStream) Reduce(fn func(ApiGetTimeSeriesDataRequest, ApiGetTimeSeriesDataRequest, int) ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	return self.ReduceInit(fn, ApiGetTimeSeriesDataRequest{})
}
func (self *ApiGetTimeSeriesDataRequestStream) ReduceInit(fn func(ApiGetTimeSeriesDataRequest, ApiGetTimeSeriesDataRequest, int) ApiGetTimeSeriesDataRequest, initialValue ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	result := ApiGetTimeSeriesDataRequestStreamOf()
	self.ForEach(func(v ApiGetTimeSeriesDataRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) ReduceInterface(fn func(interface{}, ApiGetTimeSeriesDataRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetTimeSeriesDataRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetTimeSeriesDataRequestStream) ReduceString(fn func(string, ApiGetTimeSeriesDataRequest, int) string) []string {
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
func (self *ApiGetTimeSeriesDataRequestStream) ReduceInt(fn func(int, ApiGetTimeSeriesDataRequest, int) int) []int {
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
func (self *ApiGetTimeSeriesDataRequestStream) ReduceInt32(fn func(int32, ApiGetTimeSeriesDataRequest, int) int32) []int32 {
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
func (self *ApiGetTimeSeriesDataRequestStream) ReduceInt64(fn func(int64, ApiGetTimeSeriesDataRequest, int) int64) []int64 {
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
func (self *ApiGetTimeSeriesDataRequestStream) ReduceFloat32(fn func(float32, ApiGetTimeSeriesDataRequest, int) float32) []float32 {
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
func (self *ApiGetTimeSeriesDataRequestStream) ReduceFloat64(fn func(float64, ApiGetTimeSeriesDataRequest, int) float64) []float64 {
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
func (self *ApiGetTimeSeriesDataRequestStream) ReduceBool(fn func(bool, ApiGetTimeSeriesDataRequest, int) bool) []bool {
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
func (self *ApiGetTimeSeriesDataRequestStream) Reverse() *ApiGetTimeSeriesDataRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) Replace(fn func(ApiGetTimeSeriesDataRequest, int) ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	return self.ForEach(func(v ApiGetTimeSeriesDataRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetTimeSeriesDataRequestStream) Select(fn func(ApiGetTimeSeriesDataRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetTimeSeriesDataRequestStream) Set(index int, val ApiGetTimeSeriesDataRequest) *ApiGetTimeSeriesDataRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) Skip(skip int) *ApiGetTimeSeriesDataRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetTimeSeriesDataRequestStream) SkippingEach(fn func(ApiGetTimeSeriesDataRequest, int) int) *ApiGetTimeSeriesDataRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) Slice(startIndex, n int) *ApiGetTimeSeriesDataRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetTimeSeriesDataRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) Sort(fn func(i, j int) bool) *ApiGetTimeSeriesDataRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetTimeSeriesDataRequestStream) Tail() *ApiGetTimeSeriesDataRequest {
	return self.Last()
}
func (self *ApiGetTimeSeriesDataRequestStream) TailOr(arg ApiGetTimeSeriesDataRequest) ApiGetTimeSeriesDataRequest {
	return self.LastOr(arg)
}
func (self *ApiGetTimeSeriesDataRequestStream) ToList() []ApiGetTimeSeriesDataRequest {
	return self.Val()
}
func (self *ApiGetTimeSeriesDataRequestStream) Unique() *ApiGetTimeSeriesDataRequestStream {
	return self.Distinct()
}
func (self *ApiGetTimeSeriesDataRequestStream) Val() []ApiGetTimeSeriesDataRequest {
	if self == nil {
		return []ApiGetTimeSeriesDataRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetTimeSeriesDataRequestStream) While(fn func(ApiGetTimeSeriesDataRequest, int) bool) *ApiGetTimeSeriesDataRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) Where(fn func(ApiGetTimeSeriesDataRequest) bool) *ApiGetTimeSeriesDataRequestStream {
	result := ApiGetTimeSeriesDataRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetTimeSeriesDataRequestStream) WhereSlim(fn func(ApiGetTimeSeriesDataRequest) bool) *ApiGetTimeSeriesDataRequestStream {
	result := ApiGetTimeSeriesDataRequestStreamOf()
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
